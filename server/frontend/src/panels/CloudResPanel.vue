<script setup>
import {usePanelStore} from "@/stores";
import {onMounted, ref, watch} from "vue";
import * as monaco from "monaco-editor";
import {createCloudRes, updateCloudRes} from "@/api";
import router from "@/router";

const panelStore = usePanelStore();
const props = defineProps(['data']);
const editor = ref();
const isNew = ref(true);
let monacoEditor;

onMounted(() => {
  if (props.data.name) {
    isNew.value = false
  }
  monacoEditor = monaco.editor.create(editor.value, {
    value: props.data.configs ? props.data.configs.join('\n') : '',
    language: 'ini',
    lineNumbers: 'off',
    theme: 'vs-dark',
    minimap: {
      enabled: false // 禁用预览图
    },
    automaticLayout: true,
    scrollBeyondLastLine: false
  });
})

watch(
    () => props.data.configs,
    async (newConfigs, oldConfig) => {
      monacoEditor.setValue(newConfigs ? newConfigs.join('\n') : '',)
    }
);

function save() {
  if (isNew.value) {
    createCloudRes({
      name: props.data.name,
      configs: monacoEditor.getValue().split("\n").map(s => s.trim())
    }).then((result) => {
      if (result.code === 200) {
        router.go(0)
      }
    })
  } else {
    updateCloudRes({
      name: props.data.name,
      configs: monacoEditor.getValue().split("\n").map(s => s.trim())
    }).then((result) => {
      if (result.code === 200) {
        router.go(0)
      }
    })
  }
}
</script>

<template>
  <div class="q-pa-md q-gutter-md">
    <div class="text-h5">资源配置</div>
    <div class="text-subtitle2">资源名称</div>
    <q-input dense filled v-model="props.data.name" :readonly="!isNew"/>
    <div class="text-subtitle2">文件路径</div>
    <div class="editor" ref="editor"></div>
    <q-btn color="white" text-color="primary" label="保存" @click="save"/>
    <q-btn color="white" text-color="primary" label="关闭" @click="panelStore.closePanel()"/>
  </div>
</template>

<style scoped>
.editor {
  height: 240px;
}
</style>