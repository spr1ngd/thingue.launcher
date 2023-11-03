<script setup>
import {onMounted, onUnmounted, reactive, ref, watch} from 'vue'
import {
  GetLocalServerStatus,
  LocalServerShutdown,
  LocalServerStart,
  OpenLocalServerUrl,
  UpdateLocalServerConfig,
  UpdatePeerConnectionOptions
} from "@wails/go/api/serverApi.js";
import {GetAppConfig} from "@wails/go/api/systemApi";
import {Notify} from "quasar";
import * as monaco from "monaco-editor";


const tab = ref("local")

let editor = null
const PeerConnectionOptionsRef = ref()

const localServerConfig = reactive({
  bindAddr: "",
  contentPath: "",
  autoStart: false,
  useExternalStatic: false,
  staticDir: ""
})

let localServerStatus = ref(false)

function serverStart() {
  localServerStatus.value = true
  LocalServerStart();
}

async function serverShutdown() {
  await LocalServerShutdown()
  localServerStatus.value = await GetLocalServerStatus()
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
  localServerConfig.bindAddr = appConfig.localServer.bindAddr
  localServerConfig.contentPath = appConfig.localServer.contentPath
  localServerConfig.autoStart = appConfig.localServer.autoStart
  localServerConfig.useExternalStatic = appConfig.localServer.useExternalStatic
  localServerConfig.staticDir = appConfig.localServer.staticDir
  // 获取本地server状态
  localServerStatus.value = await GetLocalServerStatus()
  //注册事件监听
  window.runtime.EventsOn("local_server_close", (err) => {
    localServerStatus.value = false
    Notify.create(`服务关闭退出信息 ${err}`)
  })
  watch(localServerConfig, async (value, oldValue, onCleanup) => {
    UpdateLocalServerConfig({
      contentPath: localServerConfig.contentPath,
      bindAddr: localServerConfig.bindAddr,
      autoStart: localServerConfig.autoStart,
      useExternalStatic: localServerConfig.useExternalStatic,
      staticDir: localServerConfig.staticDir,
    })
  })
})

onUnmounted(() => {
  window.runtime.EventsOff("local_server_close")
})
</script>

<template>
  <div class="row">
    <div class="col-6">
      <q-card style="width: 300px">
        <q-card-section class="q-pa-sm">
          <div class="row no-wrap items-center q-pa-sm">
            <div class="text-h6">内置信令服务配置</div>
          </div>
        </q-card-section>
        <q-card-section class="q-pa-none q-pt-sm">
          <q-list dense>
            <q-item>
              <q-item-section>
                <q-item-label>监听地址</q-item-label>
                <q-input :readonly="localServerStatus" dense outlined square type="text"
                         v-model="localServerConfig.bindAddr"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section>
                <q-item-label>服务路径</q-item-label>
                <q-input :readonly="localServerStatus" dense outlined square type="text"
                         v-model="localServerConfig.contentPath"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section avatar>
                <q-toggle left-label :disable="localServerStatus"
                          v-model="localServerConfig.useExternalStatic" label="使用外部静态资源"/>
              </q-item-section>
            </q-item>
            <q-item v-if="localServerConfig.useExternalStatic">
              <q-item-section>
                <q-item-label>静态资源路径</q-item-label>
                <q-input :readonly="localServerStatus" dense outlined square type="text"
                         v-model="localServerConfig.staticDir"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section avatar>
                <q-toggle left-label v-model="localServerConfig.autoStart" label="随应用启动"/>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section>
                <q-btn dense label="启动" color="positive" @click="serverStart" :disable="localServerStatus"></q-btn>
              </q-item-section>
              <q-item-section>
                <q-btn dense label="关闭" color="negative" @click="serverShutdown"
                       :disable="!localServerStatus"></q-btn>
              </q-item-section>
              <q-item-section avatar>
                <q-btn flat round icon="open_in_new" @click="OpenLocalServerUrl"/>
              </q-item-section>
            </q-item>
          </q-list>
        </q-card-section>
      </q-card>
    </div>
    <div class="col-6">
      <q-card style="">
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