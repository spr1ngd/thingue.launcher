<script setup>
import {onMounted, onUnmounted, ref} from "vue";
import {
  DeleteInstance,
  GetDefaultLaunchArguments,
  ListInstance,
  OpenInstanceLog,
  StartDownload,
  StartInstance,
  StartUpload,
  StopInstance,
} from "@wails/go/api/instanceApi";
import {OpenExplorer} from "@wails/go/api/systemApi.js";
import {
  ConnectServer,
  DisconnectServer,
  GetConnectServerOptions,
  GetServerConnInfo,
  OpenInstancePreviewUrl
} from "@wails/go/api/serverApi";

import {Notify} from "quasar";
import {GoTimeFormat, RunnerStateCodeToString} from "@/utils";

const emit = defineEmits(["openSettingsPanel", "gotoServer"])

const rows = ref([])
const selected = ref(null)
const columns = [
  {name: 'Name', field: 'name', label: '实例标识'},
  {name: 'ExecPath', field: 'execPath', label: '启动位置'},
]

const options = ref([])

const serverAddr = ref(null)
const isConnected = ref(null)

onMounted(async () => {
  await list()
  let strings = await GetConnectServerOptions();
  if (strings) {
    options.value = strings
  }

  //注册事件监听
  window.runtime.EventsOn("remote_server_conn_update", async () => {
    const connInfo = await GetServerConnInfo()
    serverAddr.value = connInfo.serverAddr
    isConnected.value = connInfo.isConnected
  })
  window.runtime.EventsOn("runner_unexpected_exit", () => {
    list()
  })
  window.runtime.EventsOn("runner_status_update", () => {
    list()
  })
  const connInfo = await GetServerConnInfo()
  serverAddr.value = connInfo.serverAddr
  isConnected.value = connInfo.isConnected
})

onUnmounted(() => {
  window.runtime.EventsOff("remote_server_conn_close", "runner_unexpected_exit", "runner_status_update")
})

async function list() {
  rows.value = await ListInstance()
}

function handleNewSettings() {
  GetDefaultLaunchArguments().then((args) => {
    emit("openSettingsPanel", {
      type: 'new',
      settings: {
        launchArguments: args.split('\n'),
        cloudRes: "",
        faultRecover: false,
        enableMultiuserControl: false,
        autoResizeRes: false,
        autoControl: false,
        stopDelay: 5,
        enableRelay: true,
        enableRenderControl: false,
        playerConfig: {
          matchViewportRes: true,
          hideUI: false,
          idleDisconnect: false,
          idleTimeout: 5
        }
      }
    })
  })

}

function handleEditSettings(row) {
  emit("openSettingsPanel", {
    type: 'edit',
    settings: row
  })
}

function handleDelete(cid) {
  DeleteInstance(cid).then(() => {
    list()
  }).catch(err => {
    Notify.create(err)
  })
}

async function handleOpenDir(path) {
  await OpenExplorer(path)
}

async function handleCloudUpload(path) {
  try {
    Notify.create(await StartUpload(path))
  } catch (e) {
    Notify.create(e)
  }
}

async function handleCloudDownload(path) {
  try {
    Notify.create(await StartDownload(path))
  } catch (e) {
    Notify.create(e)
  }
}

async function handleSelectChange() {
  if (serverAddr.value) {
    await ConnectServer(serverAddr.value);
  }
}

async function handleOpenInstanceLog(cid) {
  try {
    await OpenInstanceLog(cid)
  } catch (err) {
    Notify.create(err)
  }
}

function handleStartInstance(cid) {
  StartInstance(cid).then(() => {
    Notify.create("操作成功")
  }).catch(err => {
    Notify.create(err)
  })
}

function handleStopInstance(row) {
  row.loading = true
  StopInstance(row.cid).then(() => {
    Notify.create("进程退出成功")
  }).catch(err => {
    Notify.create(err)
  }).finally(() => {
    row.loading = false
  })
}

function handleGotoServer(tab) {
  emit("gotoServer", tab)
}

</script>

<template>
  <div class="q-pa-sm">
    <q-table grid :rows="rows" :columns="columns" v-model="selected"
             selection="multiple" hide-pagination :pagination="{rowsPerPage:0}">
      <template v-slot:top-left>
        <div class="full-width row flex-center q-gutter-sm">
          <span class="text-h6">实例列表</span>
          <q-space/>
          <q-btn dense size="sm" color="primary" round icon="add" @click="handleNewSettings"/>
        </div>
      </template>
      <template v-slot:top-right>
        <div style="min-width: 100px">
          <q-select size="sm" dense clearable :options="options" options-dense v-model="serverAddr"
                    @clear="DisconnectServer" @update:model-value="handleSelectChange">
            <template v-slot:prepend v-if="serverAddr">
              <q-btn dense flat round icon="lens" size="8.5px" :color="isConnected?'green':'red'">
                <q-tooltip>{{ isConnected ? "已连接" : "未连接，自动重连中" }}</q-tooltip>
              </q-btn>
            </template>
            <template v-slot:no-option>
              <q-item>
                <q-item-section class="text-italic text-grey">
                  <span>没有可用选项，<q-btn padding="none" color="primary" dense flat
                                            @click="handleGotoServer('local')">启动内置服务</q-btn>
                  </span>
                  <span>或 <q-btn padding="none" color="primary" dense flat
                                  @click="handleGotoServer('remote')">配置远程服务地址</q-btn>
                  </span>
                </q-item-section>
              </q-item>
            </template>
          </q-select>
        </div>
      </template>
      <template v-slot:no-data>
        <img src="@/assets/create.svg" style="padding-left: 108px"/>
        <q-space/>
        <img src="@/assets/connect.svg" style="padding-right: 40px"/>
      </template>
      <template v-slot:item="props">
        <div
            class="q-pa-sm col-xs-12 col-sm-6 col-md-4 col-lg-3 grid-style-transition"
            :style="props.selected ? 'transform: scale(0.95);' : ''"
        >
          <q-card>
            <q-card-section class="q-pt-md q-pa-none">
              <q-list dense>
                <q-item>
                  <q-item-section avatar class="clickable  cursor-pointer"
                                  @click="OpenInstancePreviewUrl(props.row.sid)">
                    <q-item-label caption class="ellipsis">名称</q-item-label>
                    <q-item-label class="ellipsis">{{ props.row.name }}</q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="min-width: 74px">
                    <q-tooltip anchor="top middle" self="center middle">
                      最后启动时间：{{ GoTimeFormat(props.row.lastStartAt) }}<br>
                      最后停止时间：{{ GoTimeFormat(props.row.lastStopAt) }}
                    </q-tooltip>
                    <q-item-label caption class="ellipsis">状态</q-item-label>
                    <q-item-label class="ellipsis">
                      {{ RunnerStateCodeToString(props.row.stateCode) }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="min-width: 70px">
                    <q-item-label caption class="ellipsis">Streamer</q-item-label>
                    <q-item-label class="ellipsis">{{
                        props.row.streamerConnected ? "已连接" : "未连接"
                      }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section style="">
                    <q-item-label caption class="ellipsis">进程号</q-item-label>
                    <q-item-label class="ellipsis">{{ props.row.pid }}</q-item-label>
                  </q-item-section>
                </q-item>
                <q-item>
                  <q-item-section @click="handleOpenDir(props.row.execPath)">
                    <q-item-label caption class="ellipsis cursor-pointer">
                      启动位置
                      <q-badge outline color="blue" v-if="props.row.isInternal">自动配置</q-badge>
                    </q-item-label>
                    <q-item-label class="ellipsis cursor-pointer">
                      {{ props.row.execPath }}
                    </q-item-label>
                  </q-item-section>
                </q-item>
              </q-list>
            </q-card-section>
            <q-card-actions class="q-pt-xs q-pl-md">
              <div class="q-gutter-md">
                <q-btn padding="none" color="green" :loading="false" flat dense icon="sym_o_play_circle"
                       @click="handleStartInstance(props.row.cid)">
                  <q-tooltip :delay="600">启动</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="red" :loading="props.row.loading" flat dense icon="sym_o_stop_circle"
                       @click="handleStopInstance(props.row)">
                  <q-tooltip :delay="600">停止</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="blue" flat dense icon="sym_o_settings"
                       @click="handleEditSettings(props.row)">
                  <q-tooltip :delay="600">编辑或查看设置</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="grey" flat dense icon="sym_o_delete" push>
                  <q-tooltip :delay="600">删除</q-tooltip>
                  <q-menu>
                    <div class="q-pa-sm">
                      确定要删除？
                      <q-btn dense size="sm" label="确认" color="blue" v-close-popup
                             @click="handleDelete(props.row.cid)"/>
                    </div>
                  </q-menu>
                </q-btn>
                <q-btn padding="none" color="info" flat dense icon="sym_o_description"
                       @click="handleOpenInstanceLog(props.row.cid)">
                  <q-tooltip :delay="600">打开实例日志</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="primary" flat dense icon="sym_o_cloud_upload"
                       @click="handleCloudUpload(props.row.cid)">
                  <q-tooltip :delay="600">上传</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="blue" flat dense icon="sym_o_cloud_download"
                       @click="handleCloudDownload(props.row.cid)">
                  <q-tooltip :delay="600">下载</q-tooltip>
                </q-btn>
                <!--                <q-icon class="flashing" color="primary" name="sym_o_cloud_upload" size="sm"></q-icon>-->
                <!--                <q-icon class="flashing" color="blue" name="sym_o_cloud_download" size="sm"></q-icon>-->
                <!--                <q-icon size="sm" color="green" name="sym_o_cloud_done"></q-icon>-->
                <!--                <q-icon size="sm" color="grey" name="sym_o_cloud_off"></q-icon>-->
                <!--                <q-icon size="sm" color="secondary" name="sym_o_cloud_sync"></q-icon>-->
              </div>
            </q-card-actions>
          </q-card>
        </div>
      </template>
    </q-table>
  </div>
</template>

<style scoped>
.flashing {
  animation: flashAnimation 1s ease infinite;
}

@keyframes flashAnimation {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0.3;
  }
}
</style>