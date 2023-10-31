<script setup>

import {onMounted, reactive, ref, watch} from "vue";
import {
  ControlRestartTask,
  GetAppConfig,
  GetVersionInfo,
  OpenFileDialog,
  UpdateSystemSettings
} from "@wails/go/api/systemApi";
import {Notify} from "quasar";

const systemSettings = reactive({
  RestartTaskCron: '',
  ExternalEditorPath: ''
})

const enableRestartTask = ref(false)
const versionInfo = ref({})

onMounted(async () => {
  const appConfig = await GetAppConfig();
  versionInfo.value = await GetVersionInfo();
  systemSettings.RestartTaskCron = appConfig.systemSettings.restartTaskCron
  systemSettings.ExternalEditorPath = appConfig.systemSettings.externalEditorPath
  enableRestartTask.value = appConfig.systemSettings.enableRestartTask
  watch(systemSettings, async (value, oldValue, onCleanup) => {
    await UpdateSystemSettings(systemSettings)
  })
})

async function updateEnableRestartTask(value, ev) {
  try {
    await ControlRestartTask(value)
  } catch (err) {
    enableRestartTask.value = !value
    Notify.create("任务开启失败，" + err)
  }
}

function select() {
  OpenFileDialog("选择文件", "*.exe").then(result => {
    if (result) {
      systemSettings.ExternalEditorPath = result;
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

</script>
<template>
  <q-card>
    <q-card-section class="q-pa-sm">
      <div class="row no-wrap items-center q-pa-sm">
        <div class="text-h6">通用设置</div>
      </div>
    </q-card-section>
    <q-card-section class="q-pa-none q-pt-sm">
      <q-list>
        <q-expansion-item>
          <template v-slot:header>
            <q-item-section avatar>
              <q-toggle v-model="enableRestartTask" @update:model-value="updateEnableRestartTask"/>
            </q-item-section>
            <q-item-section>
              <q-item-label>开启定时重启任务</q-item-label>
              <q-item-label caption>定时重启本机实例提高UE长时间运行的稳定性</q-item-label>
            </q-item-section>
          </template>
          <q-item>
            <q-item-section side>
              <div class="text-subtitle2">定时重启任务CRON表达式(5位)</div>
            </q-item-section>
            <q-item-section side>
              <q-input :disable="enableRestartTask" dense v-model="systemSettings.RestartTaskCron"/>
            </q-item-section>
          </q-item>
        </q-expansion-item>
        <q-item>
          <q-item-section avatar>
            <div class="text-subtitle2">外部日志查看器路径(默认使用vscode)</div>
          </q-item-section>
          <q-item-section>
            <q-input dense v-model="systemSettings.ExternalEditorPath">
              <template v-slot:append>
                <q-btn padding="none" icon="sym_o_file_open" flat dense @click="select"/>
              </template>
            </q-input>
          </q-item-section>
        </q-item>
        <q-expansion-item label="关于">
          <q-list dense class="q-pl-lg q-pa-sm">
            <q-item>
              <q-item-section side>
                <q-item-label>Version:</q-item-label>
              </q-item-section>
              <q-item-section side>
                <q-item-label>{{ versionInfo.Version }}</q-item-label>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section side>
                <q-item-label>BuildDate:</q-item-label>
              </q-item-section>
              <q-item-section side>
                <q-item-label>{{ versionInfo.BuildDate }}</q-item-label>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section side>
                <q-item-label>GitCommit:</q-item-label>
              </q-item-section>
              <q-item-section side>
                <q-item-label>{{ versionInfo.GitCommit }}</q-item-label>
              </q-item-section>
            </q-item>
          </q-list>
        </q-expansion-item>
      </q-list>
    </q-card-section>
  </q-card>
</template>
<style>
</style>