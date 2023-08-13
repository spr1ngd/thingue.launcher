<script setup>
import {defineEmits, onMounted, ref} from "vue";
import * as monaco from 'monaco-editor'
import {Notify} from "quasar";
import {OpenFileDialog} from "@wails/go/api/systemApi";
import {CreateInstance, SaveInstance} from "@wails/go/api/instanceApi";

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
    value: props.data.settings.LaunchArguments.join('\n'),
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
    value: props.data.settings.Metadata,
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
    value: props.data.settings.PaksConfig,
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
    props.data.settings.LaunchArguments = editor.launchArgumentsEditor.getValue().split('\n')
  })
  editor.metadataEditor.onDidChangeModelContent((event) => {
    props.data.settings.Metadata = editor.metadataEditor.getValue()
  })
  editor.paksConfigEditor.onDidChangeModelContent((event) => {
    props.data.settings.PaksConfig = editor.paksConfigEditor.getValue()
  })
})

function select() {
  OpenFileDialog("选择文件", "ThingUE (*.exe)", "*.exe").then(result => {
    if (result) {
      props.data.settings.ExecPath = result;
    } else {
      Notify.create({
        message: '选择取消'
      })
    }
  }).catch(err => {
    Notify.create({
      message: '无法选择文件'
    })
  })
}

async function save() {
  if (props.data.type === 'new') {
    await CreateInstance(props.data.settings)
    emit('openListPanel')
  } else if (props.data.type === 'edit') {
    try {
      await SaveInstance(props.data.settings)
      emit('openListPanel')
    } catch (err) {
      Notify.create({
        message: err
      })
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
          <q-item-section avatar>
            <q-item-label>实例标识</q-item-label>
            <q-input dense outlined square v-model="props.data.settings.Name"/>
          </q-item-section>
        </q-item>
        <q-item>
          <q-item-section>
            <q-item-label>启动位置</q-item-label>
            <q-input dense outlined square v-model="props.data.settings.ExecPath">
              <template v-slot:append>
                <q-icon name="sym_o_file_open" @click="select" class="cursor-pointer"/>
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
            <q-checkbox v-model="props.data.settings.FaultRecover"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>故障恢复</q-item-label>
            <q-item-label caption>
              实例非正常退出或心跳异常时尝试通过重新启动从异常状态中恢复
            </q-item-label>
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
