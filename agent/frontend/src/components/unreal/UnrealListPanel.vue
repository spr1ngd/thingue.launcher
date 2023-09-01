<script setup>
import {onMounted, onUnmounted, ref} from "vue";
import {DeleteInstance, ListInstance, StartInstance, StopInstance} from "@wails/go/api/instanceApi";
import {GetAppConfig, OpenExplorer} from "@wails/go/api/systemApi.js";
import {ConnectServer, DisconnectServer, GetConnectServerOptions, OpenInstancePreviewUrl} from "@wails/go/api/serverApi";

import {Notify} from "quasar";
import {GoTimeFormat, RunnerStateCodeToString} from "@/utils";

const emit = defineEmits(["openSettingsPanel", "gotoServer"])

const rows = ref([])

const columns = [
  {name: 'Name', field: 'name', label: '实例标识'},
  {name: 'ExecPath', field: 'execPath', label: '启动位置'},
]

const options = ref([])

const currentServer = ref(null)
const selected = ref(null)

onMounted(async () => {
  await list()
  let strings = await GetConnectServerOptions();
  if (strings) {
    options.value = strings
  }

  //注册事件监听
  window.runtime.EventsOn("remote_server_conn_close", () => {
    currentServer.value = null
  })
  window.runtime.EventsOn("runner_unexpected_exit", () => {
    list()
  })
  window.runtime.EventsOn("runner_status_update", () => {
    list()
  })
  let appConfig = await GetAppConfig();
  currentServer.value = appConfig.ServerUrl;
})

onUnmounted(() => {
  window.runtime.EventsOff("remote_server_conn_close", "runner_unexpected_exit", "runner_status_update")
})

async function list() {
  rows.value = await ListInstance()
}

function handleNewSettings() {
  emit("openSettingsPanel", {
    type: 'new',
    settings: {
      launchArguments: [
        "-AudioMixer",
        "-RenderOffScreen",
        "-ForceRes",
        "-ResX=1920",
        "-ResX=1080",
      ],
      faultRecover: false
    }
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

async function handleSelectChange() {
  if (currentServer.value) {
    try {
      await ConnectServer(currentServer.value);
      Notify.create("服务连接成功")
    } catch (e) {
      currentServer.value = ""
      Notify.create(`服务连接失败信息 ${e}`);
    }
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
    <q-table grid title="实例列表" :rows="rows" :columns="columns" v-model="selected"
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
          <q-select size="sm" dense clearable :options="options" options-dense v-model="currentServer"
                    @clear="DisconnectServer" @update:model-value="handleSelectChange">
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
                  <q-item-section avatar style="width: 100px" class="clickable  cursor-pointer"
                                  @click="OpenInstancePreviewUrl(props.row.sid)">
                    <q-item-label caption class="ellipsis">名称</q-item-label>
                    <q-item-label class="ellipsis">{{ props.row.name }}</q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="width: 100px">
                    <q-tooltip anchor="top middle" self="center middle">
                      最后启动时间：{{ GoTimeFormat(props.row.lastStartAt) }}<br>
                      最后停止时间：{{ GoTimeFormat(props.row.lastStopAt) }}
                    </q-tooltip>
                    <q-item-label caption class="ellipsis">状态</q-item-label>
                    <q-item-label class="ellipsis">
                      {{ RunnerStateCodeToString(props.row.stateCode) }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="width: 100px">
                    <q-item-label caption class="ellipsis">Streamer</q-item-label>
                    <q-item-label class="ellipsis">{{
                        props.row.streamerConnected ? "已连接" : "未连接"
                      }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="width: 100px">
                    <q-item-label caption class="ellipsis">进程号</q-item-label>
                    <q-item-label class="ellipsis">{{ props.row.pid }}</q-item-label>
                  </q-item-section>
                </q-item>
                <q-item>
                  <q-item-section @click="handleOpenDir(props.row.execPath)">
                    <q-item-label caption class="ellipsis cursor-pointer">
                      启动位置
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
                  <q-tooltip :delay="1000">启动</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="red" :loading="props.row.loading" flat dense icon="sym_o_stop_circle"
                       @click="handleStopInstance(props.row)">
                  <q-tooltip :delay="1000">停止</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="blue" flat dense icon="sym_o_settings"
                       @click="handleEditSettings(props.row)">
                  <q-tooltip :delay="1000">编辑或查看设置</q-tooltip>
                </q-btn>
                <q-btn padding="none" color="grey" flat dense icon="sym_o_delete" push>
                  <q-tooltip :delay="1000">删除</q-tooltip>
                  <q-menu>
                    <div class="q-pa-sm">
                      确定要删除？
                      <q-btn dense size="sm" label="确认" color="blue" v-close-popup
                             @click="handleDelete(props.row.cid)"/>
                    </div>
                  </q-menu>
                </q-btn>
              </div>
            </q-card-actions>
          </q-card>
        </div>
      </template>
    </q-table>
  </div>
</template>

<style scoped>

</style>