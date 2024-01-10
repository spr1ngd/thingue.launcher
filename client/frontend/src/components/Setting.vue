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
  ExternalEditorPath: '',
  LogLevel: 0
})

const enableRestartTask = ref(false)
const versionInfo = ref({})

onMounted(async () => {
  const appConfig = await GetAppConfig();
  versionInfo.value = await GetVersionInfo();
  systemSettings.RestartTaskCron = appConfig.systemSettings.restartTaskCron
  systemSettings.LogLevel = appConfig.systemSettings.logLevel
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
  <div class="text-h6">设置</div>
  <q-list>
    <q-item-label header>定时重启</q-item-label>
    <q-item dense>
      <q-item-section avatar>
        <q-toggle v-model="enableRestartTask" @update:model-value="updateEnableRestartTask"/>
      </q-item-section>
      <q-item-section>
        <q-item-label>开启定时重启任务</q-item-label>
        <q-item-label caption>定时重启本机实例提高UE长时间运行的稳定性</q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-input :disable="enableRestartTask" dense v-model="systemSettings.RestartTaskCron" label="Cron (5位)"/>
      </q-item-section>
    </q-item>

    <q-item-label header>其他</q-item-label>
    <q-item dense>
      <q-item-section>
        <q-item-label>日志输出级别</q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-select dense options-dense v-model="systemSettings.LogLevel"
                  :options="['debug','info','warn','error']"></q-select>
      </q-item-section>
    </q-item>
    <q-item dense>
      <q-item-section>
        <q-item-label>UE日志查看器路径(默认使用vscode)</q-item-label>
      </q-item-section>
      <q-item-section>
        <q-input dense v-model="systemSettings.ExternalEditorPath">
          <template v-slot:append>
            <q-btn padding="none" icon="sym_o_file_open" flat dense @click="select"/>
          </template>
        </q-input>
      </q-item-section>
    </q-item>
    <q-item-label header>版本信息</q-item-label>
    <q-item dense>
      <q-item-section>
        <q-item-label>Version:</q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-item-label>{{ versionInfo.Version }}</q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense>
      <q-item-section>
        <q-item-label>BuildDate:</q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-item-label>{{ versionInfo.BuildDate }}</q-item-label>
      </q-item-section>
    </q-item>
    <q-item dense>
      <q-item-section>
        <q-item-label>GitCommit:</q-item-label>
      </q-item-section>
      <q-item-section side>
        <q-item-label>{{ versionInfo.GitCommit }}</q-item-label>
      </q-item-section>
    </q-item>
  </q-list>
</template>
<style>
</style>