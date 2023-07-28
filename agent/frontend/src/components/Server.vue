<script setup>
import {onMounted, onUnmounted, reactive, ref, watch} from 'vue'
import {
  GetAppConfig,
  GetServerStatus,
  ServerShutdown,
  ServerStart,
  UpdateLocalServerConfig
} from "@wails/go/agent/App.js";

const localServerConfig = reactive({
  bindAddr: "",
  basePath: "",
  autoStart: "",
  enable: ""
})

let localServerStatus = ref(false)

const remoteServerConfig = reactive({
  addr: "127.0.0.1:8888",
})

function serverStart() {
  ServerStart();
}

function serverShutdown() {
  ServerShutdown()
}

onMounted(async () => {
  // 获取本地server配置
  let appConfig = await GetAppConfig();
  localServerConfig.bindAddr = appConfig.LocalServer.BindAddr
  localServerConfig.basePath = appConfig.LocalServer.BasePath
  localServerConfig.autoStart = appConfig.LocalServer.AutoStart
  localServerConfig.enable = appConfig.LocalServer.Enable
  // 获取本地server状态
  let newVar = await GetServerStatus();
  console.log(newVar)
  localServerStatus.value = newVar
  //注册事件监听
  window.runtime.EventsOn("local_server_status_update", (status) => {
    localServerStatus.value = status
  })
  watch(localServerConfig, async (value, oldValue, onCleanup) => {
    console.log(oldValue)
    UpdateLocalServerConfig({
      BasePath: localServerConfig.basePath,
      BindAddr: localServerConfig.bindAddr,
      Enable: localServerConfig.enable,
      AutoStart: localServerConfig.autoStart,
    })
  })
})

onUnmounted(async () => {

});

</script>

<template>
  <div class="q-pa-sm row items-start q-gutter-md">
    <q-card v-if="true">
      <q-card-section>
        <div class="row no-wrap items-center">
          <div class="text-subtitle1">本地信令服务</div>
          <!--          <q-icon color="red" name="lens" class="q-ma-sm">-->
          <!--            <q-tooltip>已启动</q-tooltip>-->
          <!--          </q-icon>-->
        </div>
      </q-card-section>
      <q-separator/>
      <q-card-actions vertical>
        <q-list dense>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                绑定地址：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input :readonly="localServerStatus" dense square outlined type="text" style="width: 120px"
                       v-model="localServerConfig.bindAddr"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                服务路径：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input :readonly="localServerStatus" dense square outlined style="width: 120px;" type="text"
                       v-model="localServerConfig.basePath"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section avatar>
              <q-toggle left-label v-model="localServerConfig.enable" label="是否启用" :disable="localServerStatus"/>
            </q-item-section>
            <q-item-section>
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
              <q-btn flat round icon="open_in_new"/>
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-actions>
    </q-card>
    <q-card v-if="true">
      <q-card-section>
        <div class="row no-wrap items-center">
          <div class="text-subtitle1">远程信令服务</div>
<!--          <q-icon color="green" name="lens" class="q-ma-sm">-->
<!--            <q-tooltip>已连接</q-tooltip>-->
<!--          </q-icon>-->
        </div>
      </q-card-section>
      <q-separator/>
      <q-card-actions vertical>
        <q-list dense>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                服务地址：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input dense square outlined type="text" style="width: 120px" v-model="remoteServerConfig.addr"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section avatar>
              <q-toggle left-label v-model="localServerConfig.enable" label="启用" :disable="localServerStatus"/>
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-actions>
    </q-card>
  </div>
</template>

<style scoped>

</style>