<script setup>
import {onMounted, ref} from "vue";
import * as monaco from 'monaco-editor'
import {Notify, useQuasar} from "quasar";
import {OpenFileDialog} from "@wails/go/api/systemApi";
import {CreateInstance, SaveInstance, StopInstance} from "@wails/go/api/instanceApi";

const $q = useQuasar();
const emit = defineEmits(['openListPanel'])
const props = defineProps(['data']);
const launchArgumentsEditorRef = ref()
const metadataEditorRef = ref()
const paksConfigEditorRef = ref()
const editor = {
  launchArgumentsEditor: null,
  metadataEditor: null,
  paksConfigEditor: null,
}
onMounted(async () => {
  editor.launchArgumentsEditor = monaco.editor.create(launchArgumentsEditorRef.value, {
    value: props.data.settings.launchArguments.join('\n'),
    language: 'ini',
    lineNumbers: 'off',
    theme: 'vs-dark',
    minimap: {
      enabled: false // 是否启用预览图
    },
    automaticLayout: true,
    scrollBeyondLastLine: false,
  })
  editor.launchArgumentsEditor.addAction({
    id: "id1",
    label: "填充示例",
    precondition: null,
    keybindingContext: null,
    contextMenuGroupId: "navigation",
    contextMenuOrder: 1.5,
    run: function (ed) {
      editor.launchArgumentsEditor.setValue("-AudioMixer\n-RenderOffScreen\n-ForceRes\n-ResX=1920\n-ResY=1080")
    },
  })
  editor.launchArgumentsEditor.addAction({
    id: "id2",
    label: "还原更改",
    precondition: null,
    keybindingContext: null,
    contextMenuGroupId: "navigation",
    contextMenuOrder: 2,
    run: function (ed) {
      editor.launchArgumentsEditor.setValue(props.data.settings.launchArguments.join('\n'))
    },
  })

  editor.metadataEditor = monaco.editor.create(metadataEditorRef.value, {
    value: props.data.settings.metadata,
    language: 'yaml',
    lineNumbers: 'off',
    theme: 'vs-dark',
    minimap: {
      enabled: false // 是否启用预览图
    },
    automaticLayout: true,
    scrollBeyondLastLine: false
  })
  editor.metadataEditor.getModel().updateOptions({tabSize: 2})
  editor.metadataEditor.addAction({
    id: "id1",
    label: "填充示例",
    precondition: null,
    keybindingContext: null,
    contextMenuGroupId: "navigation",
    contextMenuOrder: 1.5,
    run: function (ed) {
      editor.metadataEditor.setValue("labels: #以下是key: value格式\n  key1: value1\n  key2: value2")
    },
  })
  editor.metadataEditor.addAction({
    id: "id2",
    label: "还原更改",
    precondition: null,
    keybindingContext: null,
    contextMenuGroupId: "navigation",
    contextMenuOrder: 2,
    run: function (ed) {
      editor.metadataEditor.setValue(props.data.settings.metadata)
    },
  })

  editor.paksConfigEditor = monaco.editor.create(paksConfigEditorRef.value, {
    value: props.data.settings.paksConfig,
    language: 'yaml',
    lineNumbers: 'off',
    theme: 'vs-dark',
    minimap: {
      enabled: false // 是否启用预览图
    },
    automaticLayout: true,
    scrollBeyondLastLine: false
  })
  editor.paksConfigEditor.getModel().updateOptions({tabSize: 2})
  editor.paksConfigEditor.addAction({
    id: "id1",
    label: "填充示例",
    precondition: null,
    keybindingContext: null,
    contextMenuGroupId: "navigation",
    contextMenuOrder: 1.5,
    run: function (ed) {
      editor.paksConfigEditor.setValue("paks:\n  - name: 宜宾换流站    #列表里显示名称\n    value: yibin       #pak目录名称 \n" +
          "  - name: 雁门关换流站\n    value: yanmenguan\n  - name: 中都换流站\n    value: zhongdu")
    },
  })
  editor.paksConfigEditor.addAction({
    id: "id2",
    label: "还原更改",
    precondition: null,
    keybindingContext: null,
    contextMenuGroupId: "navigation",
    contextMenuOrder: 2,
    run: function (ed) {
      editor.paksConfigEditor.setValue(props.data.settings.paksConfig)
    },
  })

})

function select() {
  OpenFileDialog("选择文件", "ThingUE (*.exe, *.sh)", "*.exe;*.sh").then(result => {
    if (result) {
      props.data.settings.execPath = result;
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
  settings.stopDelay = Number(settings.stopDelay)
  settings.launchArguments = editor.launchArgumentsEditor.getValue().split('\n')
  settings.metadata = editor.metadataEditor.getValue()
  settings.paksConfig = editor.paksConfigEditor.getValue()
  if (props.data.type === 'new') {
    await CreateInstance(settings)
    emit('openListPanel')
  } else if (props.data.type === 'edit') {
    try {
      await SaveInstance(settings)
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
  <q-card>
    <q-card-section class="q-pa-sm">
      <div class="row no-wrap items-center q-pa-sm">
        <div class="text-h6">实例配置</div>
        <q-space/>
        <div class="q-gutter-md">
          <q-btn color="primary" @click="save">保存</q-btn>
          <q-btn @click="emit('openListPanel')">关闭</q-btn>
        </div>
      </div>
    </q-card-section>
    <q-card-section class="q-pa-none q-pt-sm">
      <q-list>
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
            <q-input dense outlined square v-model="props.data.settings.name"/>
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
            <q-input dense outlined square v-model="props.data.settings.cloudRes"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>&nbsp</q-item-label>
            <div class="col-6">
              <q-tooltip anchor="top middle" self="center middle" :delay="1000">
                启用h265编码以支持8K，需要配合ThingBrowser使用
              </q-tooltip>
              <q-toggle label="H265编码" v-model="props.data.settings.enableH265"/>
            </div>
          </q-item-section>
        </q-item>
        <q-item>
          <q-item-section>
            <q-item-label>启动位置</q-item-label>
            <q-input dense outlined square v-model="props.data.settings.execPath">
              <template v-slot:append>
                <q-btn padding="none" icon="sym_o_file_open" flat dense @click="select"/>
              </template>
            </q-input>
          </q-item-section>
        </q-item>
        <q-expansion-item label="启动参数配置" caption="编辑虚幻引擎的启动参数">
          <q-card class="q-pa-md">
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
        <q-separator/>
        <q-item tag="label" v-ripple>
          <q-item-section side top>
            <q-checkbox v-model="props.data.settings.enableRelay"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>使用WebRTC中继</q-item-label>
            <q-item-label caption>
              当中继服务可用时开启此配置就可以通过中继的方式访问UE实例
            </q-item-label>
          </q-item-section>
        </q-item>
        <q-item tag="label" v-ripple>
          <q-item-section side top>
            <q-checkbox v-model="props.data.settings.enableRenderControl"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>渲染控制</q-item-label>
            <q-item-label caption>
              是否根据连接数控制开启关闭渲染
            </q-item-label>
          </q-item-section>
        </q-item>
        <q-item tag="label" v-ripple>
          <q-item-section side top>
            <q-checkbox v-model="props.data.settings.faultRecover"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>故障恢复</q-item-label>
            <q-item-label caption>
              实例非正常退出或心跳异常时尝试通过重新启动从异常状态中恢复
            </q-item-label>
          </q-item-section>
        </q-item>
        <q-item tag="label" v-ripple>
          <q-item-section side top>
            <q-checkbox v-model="props.data.settings.autoControl"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>自动启停</q-item-label>
            <q-item-label caption>
              实例有访问时自动开启，无访问时自动关闭
            </q-item-label>
          </q-item-section>
          <q-item-section side>
            <q-input dense v-model="props.data.settings.stopDelay" label="关闭延迟时间（秒）" type="number"/>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card-section>
  </q-card>
</template>
<style scoped>
.editor {
  height: 240px;
}
</style>
