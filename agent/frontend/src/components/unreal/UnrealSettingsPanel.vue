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
  });
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
  });
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
  });

  editor.launchArgumentsEditor.onDidChangeModelContent((event) => {
    props.data.settings.launchArguments = editor.launchArgumentsEditor.getValue().split('\n')
  })
  editor.metadataEditor.onDidChangeModelContent((event) => {
    props.data.settings.metadata = editor.metadataEditor.getValue()
  })
  editor.paksConfigEditor.onDidChangeModelContent((event) => {
    props.data.settings.paksConfig = editor.paksConfigEditor.getValue()
  })
})

function select() {
  OpenFileDialog("选择文件", "ThingUE (*.exe)", "*.exe").then(result => {
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
  props.data.settings.stopDelay = Number(props.data.settings.stopDelay)
  if (props.data.type === 'new') {
    await CreateInstance(props.data.settings)
    emit('openListPanel')
  } else if (props.data.type === 'edit') {
    try {
      await SaveInstance(props.data.settings)
      emit('openListPanel')
    } catch (err) {
      if (err === "实例运行中无法修改配置") {
        $q.dialog({
          title: '确认',
          message: err + '，你想要立即停止当前实例吗?',
          cancel: true,
          persistent: true
        }).onOk(() => {
          StopInstance(props.data.settings.cid).then(() => {
            Notify.create("进程退出成功")
            console.log("进程退出成功")
          }).catch(err => {
            console.log(err)
            Notify.create(err)
          })
          console.log(props.data.settings.cid)
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
            <q-item-label>实例名称</q-item-label>
            <q-input dense outlined square v-model="props.data.settings.name"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>&nbsp</q-item-label>
            <div class="row">
              <div class="col-6">
                <q-tooltip anchor="top middle" self="center middle" :delay="1000">
                  启用h265编码以支持8K，需要配合ThingBrowser使用
                </q-tooltip>
                <q-toggle label="H265编码" v-model="props.data.settings.enableH265"/>
              </div>
<!--              <div class="col-6">-->
<!--                <q-tooltip anchor="top middle" self="center middle" :delay="1000">-->
<!--                  启用分辨率自适应前端会自动根据浏览器窗口大小调整分辨率-->
<!--                </q-tooltip>-->
<!--                <q-toggle label="分辨率自适应" v-model="props.data.settings.autoResizeRes"/>-->
<!--              </div>-->
            </div>
          </q-item-section>
          <q-item-section side>
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
            <q-input  dense v-model="props.data.settings.stopDelay" label="关闭延迟时间（秒）" type="number"/>
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
