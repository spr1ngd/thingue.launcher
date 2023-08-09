<script setup>
import { ref, watch, onMounted, reactive } from 'vue';
import * as monaco from 'monaco-editor';
import { createInstanceConfig, saveInstanceConfig } from '@/api/agent';
import { Notify } from 'quasar';
const props = defineProps(['row', 'sessionId']);
const config = reactive({
  name: props.row.name,
  execPath: props.row.execPath,
  enableMultiuserControl: props.row.enableMultiuserControl
});
const dom = ref();
let editor;
onMounted(() => {
  editor = monaco.editor.create(dom.value, {
    value: props.row.runParameters.join('\n'),
    language: 'ini',
    lineNumbers: 'off',
    theme: 'vs-dark',
    minimap: {
      enabled: false // 是否启用预览图
    },
    automaticLayout: true,
    scrollBeyondLastLine: false
  });
});

watch(
  () => props.row.name,
  async (newValue, oldValue) => {
    config.name = newValue;
  }
);

watch(
  () => props.row.execPath,
  async (newValue, oldValue) => {
    config.execPath = newValue;
  }
);

watch(
  () => props.row.enableMultiuserControl,
  async (newValue, oldValue) => {
    config.enableMultiuserControl = newValue;
  }
);

watch(
  () => props.row.runParameters,
  async (newValue, oldValue) => {
    editor.setValue(newValue.join('\n'));
  }
);

function doConfigSave() {
  saveInstanceConfig({
    sessionId: props.sessionId,
    oldName: props.row.name,
    name: config.name,
    execPath: config.execPath,
    enableMultiuserControl: config.enableMultiuserControl,
    runParameters: editor.getValue().split('\n')
  }).then((r) => {
    if(r.code===500) reset()
  });
}

function reset() {
  config.name = props.row.name;
  config.execPath = props.row.execPath;
  config.enableMultiuserControl = props.row.enableMultiuserControl
  editor.setValue(props.row.runParameters.join('\n'));
}

</script>
<template>
  <div class="q-pa-md q-gutter-md">
    <div class="text-h5">配置</div>
    <div class="text-subtitle2">实例名称（不可重复）</div>
    <q-input dense filled v-model="config.name" />
    <div class="text-subtitle2">可执行文件路径</div>
    <q-input dense filled v-model="config.execPath" />
    <q-toggle v-model="config.enableMultiuserControl" left-label label="协同使用" />
    <div class="text-subtitle2">启动参数</div>
    <div class="editor" ref="dom"></div>
    <q-btn color="primary" label="保存" @click="doConfigSave" />
    <q-btn color="white" text-color="primary" label="关闭" @click="$emit('close')" />
  </div>
</template>
<style scoped>
.editor {
  height: 240px;
}
</style>
