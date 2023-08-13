<script setup>
import {onMounted, onUnmounted, reactive, ref, watch} from 'vue'
import {GetLocalServerStatus, LocalServerShutdown, LocalServerStart, UpdateLocalServerConfig} from "@wails/go/api/serverApi.js";
import {GetAppConfig} from "@wails/go/api/systemApi";


const tab = ref("local")

const localServerConfig = reactive({
  bindAddr: "",
  basePath: "",
  autoStart: "",
  enable: ""
})

let localServerStatus = ref(false)

async function serverStart() {
  await LocalServerStart();
  localServerStatus.value = await GetLocalServerStatus()
}

async function serverShutdown() {
  await LocalServerShutdown()
  localServerStatus.value = await GetLocalServerStatus()
}

function handleOpenExplorer() {
  const port = localServerConfig.bindAddr.split(":")[1]
  window.runtime.BrowserOpenURL(`http://localhost:${port}${localServerConfig.basePath}/static/`)
}

onMounted(async () => {
  // 获取本地server配置
  let appConfig = await GetAppConfig();
  localServerConfig.bindAddr = appConfig.LocalServer.BindAddr
  localServerConfig.basePath = appConfig.LocalServer.BasePath
  localServerConfig.autoStart = appConfig.LocalServer.AutoStart
  localServerConfig.enable = appConfig.LocalServer.Enable
  // 获取本地server状态
  localServerStatus.value = await GetLocalServerStatus()
  //注册事件监听
  window.runtime.EventsOn("local_server_status_update", (status) => {
    localServerStatus.value = status
  })
  watch(localServerConfig, async (value, oldValue, onCleanup) => {
    UpdateLocalServerConfig({
      BasePath: localServerConfig.basePath,
      BindAddr: localServerConfig.bindAddr,
      Enable: localServerConfig.enable,
      AutoStart: localServerConfig.autoStart,
    })
  })
})
</script>

<template>
  <q-card style="width: 300px">
    <q-card-section class="q-pa-sm">
      <div class="row no-wrap items-center q-pa-sm">
        <div class="text-h6">本地信令服务配置</div>
      </div>
    </q-card-section>
    <q-card-section class="q-pa-none q-pt-sm">
      <q-list dense>
        <q-item>
          <q-item-section>
            <q-item-label>绑定地址</q-item-label>
            <q-input :readonly="localServerStatus" dense outlined square type="text"
                     v-model="localServerConfig.bindAddr"/>
          </q-item-section>
        </q-item>
        <q-item>
          <q-item-section>
            <q-item-label>服务路径</q-item-label>
            <q-input :readonly="localServerStatus" dense outlined square type="text"
                     v-model="localServerConfig.basePath"/>
          </q-item-section>
        </q-item>
        <q-item>
          <q-item-section avatar>
            <q-checkbox left-label v-model="localServerConfig.autoStart" label="随应用启动"
                        :disable="localServerStatus"/>
          </q-item-section>
        </q-item>
        <q-item>
          <q-item-section>
            <q-btn dense label="启动" color="positive" @click="serverStart" :disable="localServerStatus"></q-btn>
          </q-item-section>
          <q-item-section>
            <q-btn dense label="关闭" color="negative" @click="serverShutdown" :disable="!localServerStatus"></q-btn>
          </q-item-section>
          <q-item-section avatar>
            <q-btn flat round icon="open_in_new" @click="handleOpenExplorer"/>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card-section>
  </q-card>
</template>

<style scoped>

</style>