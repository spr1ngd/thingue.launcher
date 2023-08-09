<script setup>

import {onMounted, reactive, ref, watch} from "vue";
import {ControlRestartTask, GetAppConfig, UpdateSystemSettings} from "@wails/go/app/App";
import {Notify} from "quasar";

const systemSettings = reactive({
  RestartTaskCron: ''
})

const enableRestartTask = ref(false)

onMounted(async () => {
  const appConfig = await GetAppConfig();
  systemSettings.RestartTaskCron = appConfig.SystemSettings.RestartTaskCron
  enableRestartTask.value = appConfig.EnableRestartTask
  watch(systemSettings, async (value, oldValue, onCleanup) => {
    UpdateSystemSettings(systemSettings)
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

function change() {
  console.log("sss")
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
              <q-input dense v-model="systemSettings.RestartTaskCron"/>
            </q-item-section>
          </q-item>
        </q-expansion-item>
        <q-expansion-item label="关于">
          <q-list dense class="q-pl-lg q-pa-sm">
            <q-item>
              <q-item-section side>
                <q-item-label>版本号</q-item-label>
              </q-item-section>
              <q-item-section side>
                <q-item-label>v0.0.1</q-item-label>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section side>
                <q-item-label>编译日期</q-item-label>
              </q-item-section>
              <q-item-section side>
                <q-item-label>2023-08-01</q-item-label>
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