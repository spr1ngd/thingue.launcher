<script setup>
import {onMounted, onUnmounted, reactive, ref, watch} from 'vue'
import {
  GetHttpServerStatus,
  GetGrpcServerStatus,
  LocalServerShutdown,
  LocalServerStart,
  OpenLocalServerUrl,
  UpdateLocalServerConfig,
  UpdatePeerConnectionOptions
} from "@wails/go/api/serverApi.js";
import {GetAppConfig} from "@wails/go/api/systemApi";
import {Notify} from "quasar";
import * as monaco from "monaco-editor";
import {GetMqttServerStatus} from "@wails/go/api/serverApi";


const tab = ref("local")

let editor = null
const PeerConnectionOptionsRef = ref()

const localServer = reactive({
  config: {
    bindAddr: "",
    contentPath: "",
    autoStart: false,
    enableMQTT: false,
    useExternalStatic: false,
    staticDir: ""
  },
  status: {
    httpServerRunning: false,
    grpcServerRunning: false,
    mqttServerRunning: false
  }
})

async function serverStart() {
  localServer.status.httpServerRunning = true
  localServer.status.grpcServerRunning = true
  try {
    await LocalServerStart();
  } catch (error) {
    localServer.status.httpServerRunning = false
    localServer.status.grpcServerRunning = false
    Notify.create(`启动失败，${error}`)
  }
}

async function serverShutdown() {
  try {
    await LocalServerShutdown()
  } catch (error) {
    Notify.create(`关闭失败，${error}`)
  }
}

async function savePeerConnectionOptions() {
  try {
    await UpdatePeerConnectionOptions(editor.getValue())
    Notify.create(`保存成功，已启动实例重启后生效`)
  } catch (e) {
    Notify.create(`保存失败，${e}`)
  }
}

function fillSampleCode() {
  editor.setValue("#配置示例\r\niceServers:\r\n  - urls:\r\n    - stun:10.100.40.6:19303\r\n    - turn:10.100.40.6:19303\r\n    username: username\r\n    credential: password")
}

async function rollbackChange() {
  let appConfig = await GetAppConfig();
  editor.setValue(appConfig.peerConnectionOptions)
}

async function getServerStatus() {
  localServer.status.httpServerRunning = await GetHttpServerStatus()
  localServer.status.grpcServerRunning = await GetGrpcServerStatus()
  localServer.status.mqttServerRunning = await GetMqttServerStatus()
}

onMounted(async () => {
  editor = monaco.editor.create(PeerConnectionOptionsRef.value, {
    value: "",
    language: 'yaml',
    lineNumbers: 'off',
    theme: 'vs-dark',
    minimap: {
      enabled: false // 是否启用预览图
    },
    automaticLayout: true,
    scrollBeyondLastLine: false,
  });
  editor.getModel().updateOptions({tabSize: 2})
  // 获取本地server配置
  let appConfig = await GetAppConfig();
  editor.setValue(appConfig.peerConnectionOptions)
  localServer.config = appConfig.localServer
  // 获取server状态
  await getServerStatus()
  // 注册事件监听
  window.runtime.EventsOn("local_server_close", async (errMsg) => {
    await getServerStatus()
    if (errMsg) {
      Notify.create(`服务关闭退出 ${errMsg}`)
    }
  })
  watch(localServer, async (value, oldValue, onCleanup) => {
    UpdateLocalServerConfig(localServer.config)
  })
})

onUnmounted(() => {
  window.runtime.EventsOff("local_server_close")
})
</script>

<template>
  <div class="row reverse-wrap">
    <div class="col-auto q-pa-sm" style="width: 320px">
      <q-card>
        <q-card-section class="q-pa-sm">
          <div class="row no-wrap items-center q-pa-sm">
            <div class="text-h6">服务控制面板</div>
            <q-space/>
            <div class="q-gutter-xs">
              <q-btn dense flat round icon="lens" size="8.5px"
                     :color="localServer.status.httpServerRunning?'green':'grey'">
                <q-tooltip anchor="top middle" self="center middle">
                  {{ localServer.status.httpServerRunning ? "HTTP服务运行中" : "HTTP服务未运行" }}
                </q-tooltip>
              </q-btn>
              <q-btn dense flat round icon="lens" size="8.5px"
                     :color="localServer.status.grpcServerRunning?'green':'grey'">
                <q-tooltip anchor="top middle" self="center middle">
                  {{ localServer.status.grpcServerRunning ? "gRPC服务运行中" : "gRPC服务未运行" }}
                </q-tooltip>
              </q-btn>
              <q-btn dense flat round icon="lens" size="8.5px"
                     :color="localServer.status.mqttServerRunning?'green':'grey'">
                <q-tooltip anchor="top middle" self="center middle">
                  {{ localServer.status.mqttServerRunning ? "MQTT服务运行中" : "MQTT服务未运行" }}
                </q-tooltip>
              </q-btn>
            </div>
          </div>
        </q-card-section>
        <q-card-section class="q-pa-none">
          <q-list dense>
            <q-item>
              <q-item-section>
                <q-item-label>监听地址</q-item-label>
                <q-input :readonly="localServer.status.httpServerRunning" dense outlined square type="text"
                         v-model="localServer.config.bindAddr"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section>
                <q-item-label>服务路径</q-item-label>
                <q-input :readonly="localServer.status.httpServerRunning" dense outlined square type="text"
                         v-model="localServer.config.contentPath"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section avatar>
                <q-toggle left-label :disable="localServer.status.httpServerRunning"
                          v-model="localServer.config.useExternalStatic" label="使用外部静态资源"/>
              </q-item-section>
            </q-item>
            <q-item v-if="localServer.config.useExternalStatic">
              <q-item-section>
                <q-item-label>静态资源路径</q-item-label>
                <q-input :readonly="localServer.status.httpServerRunning" dense outlined square type="text"
                         v-model="localServer.config.staticDir"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section avatar>
                <q-toggle left-label :disable="localServer.status.httpServerRunning"
                          v-model="localServer.config.enableMQTT" label="启用MQTT"/>
              </q-item-section>
              <q-item-section>
                <q-toggle left-label v-model="localServer.config.autoStart" label="自动启动">
                  <q-tooltip anchor="top middle" self="center middle" :delay="600">是否随应用程序自动启动</q-tooltip>
                </q-toggle>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section>
                <q-btn dense label="启动" color="positive" @click="serverStart"></q-btn>
              </q-item-section>
              <q-item-section>
                <q-btn dense label="关闭" color="negative" @click="serverShutdown"></q-btn>
              </q-item-section>
              <q-item-section avatar>
                <q-btn flat round icon="open_in_new" @click="OpenLocalServerUrl">
                  <q-tooltip :delay="600">打开控制台</q-tooltip>
                </q-btn>
              </q-item-section>
            </q-item>
          </q-list>
        </q-card-section>
      </q-card>
    </div>
    <div class="col q-pa-sm" style="min-width: 320px">
      <q-card>
        <q-card-section class="q-pa-sm">
          <div class="row no-wrap items-center q-pa-sm">
            <div class="text-h6">WebRTC中继配置</div>
            <q-icon name="sym_o_help" color="grey" class="q-pl-sm" size="xs">
              <q-tooltip max-width="200px">
                如果实例和浏览器之间的WebRTC连接被防火墙阻止或者是复杂网络环境下的跨网段相互访问，需要在这里配置中继服务
              </q-tooltip>
            </q-icon>
            <q-space/>
            <q-btn-dropdown split class="glossy" color="teal" label="保存" @click="savePeerConnectionOptions">
              <q-list>
                <q-item clickable v-close-popup @click="fillSampleCode" dense>
                  <q-item-section>
                    <q-item-label>填充示例</q-item-label>
                  </q-item-section>
                </q-item>
                <q-item clickable v-close-popup @click="rollbackChange" dense>
                  <q-item-section>
                    <q-item-label>还原更改</q-item-label>
                  </q-item-section>
                </q-item>
              </q-list>
            </q-btn-dropdown>
          </div>
        </q-card-section>
        <q-card-section class="q-pa-sm q-pt-none">
          <div class="editor" ref="PeerConnectionOptionsRef"></div>
        </q-card-section>
      </q-card>
    </div>
  </div>
</template>

<style scoped>
.editor {
  height: 240px;
}
</style>