<script setup>
import {onMounted, ref, watch} from "vue";
import {Notify, useQuasar} from "quasar";
import {OpenFileDialog} from "@wails/go/api/systemApi";
import {CreateInstance, UpdateConfig, StopInstance} from "@wails/go/api/instanceApi";
import {createLaunchArgumentsEditor, createMetadataEditor, createPaksConfigEditor} from "@/components/unreal/settingsEditor";

const $q = useQuasar();
const emit = defineEmits(['openListPanel'])
const props = defineProps(['data']);
const launchArgumentsEditorRef = ref()
const metadataEditorRef = ref()
const paksConfigEditorRef = ref()
const enableH265 = ref(false)
const editor = {
  launchArgumentsEditor: null,
  metadataEditor: null,
  paksConfigEditor: null,
}
onMounted(async () => {
  enableH265.value = props.data.settings.instanceConfig.launchArguments.includes('-PSForceH265')
  editor.launchArgumentsEditor = createLaunchArgumentsEditor(launchArgumentsEditorRef.value, props.data.settings.instanceConfig.launchArguments)
  editor.metadataEditor = createMetadataEditor(metadataEditorRef.value, props.data.settings.instanceConfig.metadata)
  editor.paksConfigEditor = createPaksConfigEditor(paksConfigEditorRef.value, props.data.settings.instanceConfig.paksConfig)
})

watch(enableH265, (newValue, oldValue) => {
  const lines = editor.launchArgumentsEditor.getValue().split('\n');
  if (newValue) {
    if (!lines.includes('-PSForceH265')) {
      lines.push('-PSForceH265')
    }
  } else {
    lines.forEach((line, index) => {
      if (line === '-PSForceH265') {
        lines.splice(index, 1)
      }
    })
  }
  editor.launchArgumentsEditor.setValue(lines.join('\n'))
})

function select() {
  OpenFileDialog("选择文件", "ThingUE (*.exe, *.sh)", "*.exe;*.sh").then(result => {
    if (result) {
      props.data.settings.instanceConfig.execPath = result;
    } else {
      Notify.create({
        message: '文件选择取消'
      })
    }
  }).catch(err => {
    Notify.create({
      message: '文件选择出错,' + err
    })
  })
}

async function save() {
  const settings = JSON.parse(JSON.stringify(props.data.settings))
  settings.instanceConfig.stopDelay = Number(settings.instanceConfig.stopDelay)
  settings.playerConfig.idleTimeout = Number(settings.playerConfig.idleTimeout)
  settings.instanceConfig.launchArguments = editor.launchArgumentsEditor.getValue().split('\n')
  settings.instanceConfig.metadata = editor.metadataEditor.getValue()
  settings.instanceConfig.paksConfig = editor.paksConfigEditor.getValue()
  if (props.data.type === 'new') {
    try {
      await CreateInstance(settings)
      emit('openListPanel')
    } catch (err) {
      Notify.create({
        message: err
      })
    }
  } else if (props.data.type === 'edit') {
    try {
      await UpdateConfig(settings)
      emit('openListPanel')
    } catch (err) {
      if (err === "实例运行中无法修改配置") {
        $q.dialog({
          title: '确认',
          message: err + '，你想要立即停止当前实例吗?',
          cancel: true,
          persistent: true
        }).onOk(() => {
          StopInstance(settings.cid).then(() => {
            Notify.create("进程退出成功，请重新保存")
          }).catch(err => {
            Notify.create(err)
          })
        })
      } else {
        Notify.create({
          message: err
        })
      }
    }
  }
}
</script>

<template>
  <div class="row">
    <span class="text-h6">实例配置</span>
    <q-space></q-space>
    <div class="q-gutter-md">
      <q-btn color="primary" @click="save">保存</q-btn>
      <q-btn @click="emit('openListPanel')">关闭</q-btn>
    </div>
  </div>


  <q-list>
    <q-item-label header>基本设置</q-item-label>
    <q-item>
      <q-item-section>
        <q-item-label style="display: flex;align-items: center">
          <span>实例名称</span>
          <q-icon name="sym_o_help" color="grey" class="q-pl-xs" size="xs">
            <q-tooltip anchor="top middle" self="center middle">
              设置有意义的实例名称作为标识
            </q-tooltip>
          </q-icon>
        </q-item-label>
        <q-input dense outlined square v-model="props.data.settings.instanceConfig.name"/>
      </q-item-section>
      <q-item-section>
        <q-item-label style="display: flex;align-items: center">
          <span>云资源</span>
          <q-icon name="sym_o_help" color="grey" class="q-pl-xs" size="xs">
            <q-tooltip anchor="top middle" self="center middle">
              具有相同云资源标识的实例之间可以同步云文件
            </q-tooltip>
          </q-icon>
        </q-item-label>
        <q-input dense outlined square v-model="props.data.settings.instanceConfig.cloudRes"/>
      </q-item-section>
    </q-item>
    <q-item>
      <q-item-section>
        <q-item-label>启动位置</q-item-label>
        <q-input dense outlined square v-model="props.data.settings.instanceConfig.execPath">
          <template v-slot:append>
            <q-btn padding="none" icon="sym_o_file_open" flat dense @click="select"/>
          </template>
        </q-input>
      </q-item-section>
    </q-item>
    <q-expansion-item label="启动参数配置" caption="编辑虚幻引擎的启动参数">
      <q-card class="q-pl-md q-pr-md">
        <q-toggle label="H265编码" v-model="enableH265"/>
        <div class="editor" ref="launchArgumentsEditorRef"></div>
      </q-card>
    </q-expansion-item>
    <q-expansion-item label="元数据配置" caption="设置元数据作为实例的自定义信息">
      <q-card class="q-pa-md">
        <div class="editor" ref="metadataEditorRef"></div>
      </q-card>
    </q-expansion-item>
    <q-expansion-item label="Pak资源配置" caption="设置壳加载模式下Pak资源选择切换列表">
      <q-card class="q-pa-md">
        <div class="editor" ref="paksConfigEditorRef"></div>
      </q-card>
    </q-expansion-item>
    <q-separator spaced/>

    <q-item-label header>高级设置</q-item-label>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.instanceConfig.enableRenderControl"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>渲染控制</q-item-label>
        <q-item-label caption>
          是否根据连接数控制开启关闭渲染
        </q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.instanceConfig.enableMultiuserControl"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>多用户同时操作</q-item-label>
        <q-item-label caption>
          是否允许多用户同时操作
        </q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.instanceConfig.faultRecover"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>故障恢复</q-item-label>
        <q-item-label caption>
          实例非正常退出或心跳异常时尝试通过重新启动从异常状态中恢复
        </q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.instanceConfig.autoControl"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>自动启停</q-item-label>
        <q-item-label caption>
          实例有访问时自动开启，无访问时自动关闭
        </q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-input dense v-model="props.data.settings.instanceConfig.stopDelay" label="关闭延迟时间（秒）" type="number"/>
      </q-item-section>
    </q-item>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.instanceConfig.enableRelay"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>使用WebRTC中继</q-item-label>
        <q-item-label caption>
          当中继服务可用时开启此配置就可以通过中继的方式访问UE实例
        </q-item-label>
      </q-item-section>
    </q-item>
    <q-separator spaced/>

    <q-item-label header>Player设置</q-item-label>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.playerConfig.matchViewportRes"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>根据视口大小调整分辨率</q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.playerConfig.hideUI"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>隐藏控制UI</q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense tag="label" v-ripple>
      <q-item-section side top>
        <q-checkbox v-model="props.data.settings.playerConfig.idleDisconnect"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>无操作关闭连接</q-item-label>
        <q-item-label caption>
          一段时间内没有收到键盘鼠标输入事件后关闭连接
        </q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-input dense v-model="props.data.settings.playerConfig.idleTimeout" label="无操作等待时间（分钟）"
                 type="number"/>
      </q-item-section>
    </q-item>
    <q-separator spaced/>

    <q-item>
      <div class="q-gutter-md">
        <q-btn color="primary" @click="save">保存</q-btn>
        <q-btn @click="emit('openListPanel')">关闭</q-btn>
      </div>
    </q-item>

  </q-list>
</template>
<style scoped>
.editor {
  height: 240px;
}
</style>
