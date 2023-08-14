<script setup>
import {defineEmits, onMounted, ref} from "vue";
import {DeleteInstance, ListInstance, StartInstance, StopInstance} from "@wails/go/api/instanceApi";
import {GetAppConfig, OpenExplorer} from "@wails/go/api/systemApi.js";
import {ConnectServer, DisconnectServer, GetConnectServerOptions} from "@wails/go/api/serverApi";

import {Notify} from "quasar";
import {RunnerStateCodeToString} from "@/utils";

const emit = defineEmits(["openSettingsPanel"])

const rows = ref([])

const columns = [
  {name: 'Name', field: 'Name', label: '实例标识'},
  {name: 'ExecPath', field: 'ExecPath', label: '启动位置'},
]

const options = ref([])

const currentServer = ref(null)
const selected = ref(null)

onMounted(async () => {
  await list()
  options.value = await GetConnectServerOptions()

  //注册事件监听
  window.runtime.EventsOn("remote_server_conn_close", () => {
    currentServer.value = null
  })
  window.runtime.EventsOn("runner_unexpected_exit", () => {
    list()
  })
  let appConfig = await GetAppConfig();
  currentServer.value = appConfig.ServerUrl;
})

onMounted(() => {
  window.runtime.EventsOff("remote_server_conn_close", "runner_unexpected_exit")
})

async function list() {
  rows.value = await ListInstance()
}

function handleNewSettings() {
  emit("openSettingsPanel", {
    type: 'new',
    settings: {
      LaunchArguments: [
        "-AudioMixer",
        "-RenderOffScreen",
        "-ForceRes",
        "-ResX=1920",
        "-ResX=1080",
      ]
    }
  })
}

function handleEditSettings(row) {
  emit("openSettingsPanel", {
    type: 'edit',
    settings: row
  })
}

function handleDelete(id) {
  DeleteInstance(id).then(() => {
    list()
  }).catch(err => {
    Notify.create(err)
  })
}

async function handleOpenDir(path) {
  await OpenExplorer(path)
}

async function handleOpenPreview(name) {
  const http = currentServer.value;
  const url = `${http.endsWith("/") ? http : http + "/"}static/player.html?name=${name}`
  window.runtime.BrowserOpenURL(url)
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

function handleStartInstance(id) {
  StartInstance(id).then(() => {
    Notify.create("操作成功")
    list()
  }).catch(err => {
    Notify.create(err)
  })
}

function handleStopInstance(id) {
  StopInstance(id).then(() => {
    Notify.create("进程退出成功")
    list()
  }).catch(err => {
    Notify.create(err)
  })
}


</script>

<template>
  <div class="q-pa-sm">
    <q-table grid title="实例列表" :rows="rows" :columns="columns" v-model="selected"
             selection="multiple" hide-pagination :pagination="{rowsPerPage:0}" hide-no-data>
      <template v-slot:top-right>
        <q-select size="sm" dense clearable :options="options" options-dense v-model="currentServer"
                  @clear="DisconnectServer" @update:model-value="handleSelectChange"/>
        <q-space/>
        <q-btn dense size="sm" color="primary" round icon="add" @click="handleNewSettings"/>
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
                                  @click="handleOpenPreview(props.row.Name)">
                    <q-item-label caption class="ellipsis">标识
                    </q-item-label>
                    <q-item-label class="ellipsis">{{
                        props.row.Name
                      }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="width: 100px">
                    <q-item-label caption class="ellipsis">状态</q-item-label>
                    <q-item-label class="ellipsis">{{ RunnerStateCodeToString(props.row.StateCode) }}</q-item-label>
                  </q-item-section>
                </q-item>
                <q-item>
                  <q-item-section @click="handleOpenDir(props.row.ExecPath)">
                    <q-item-label caption class="ellipsis cursor-pointer">
                      启动位置
                    </q-item-label>
                    <q-item-label class="ellipsis cursor-pointer">
                      {{ props.row.ExecPath }}
                    </q-item-label>
                  </q-item-section>
                </q-item>
              </q-list>
            </q-card-section>
            <q-card-actions class="q-pt-xs">
              <div class="q-gutter-md">
                <q-btn color="green" flat dense icon="sym_o_play_circle" @click="handleStartInstance(props.row.ID)"/>
                <q-btn color="red" flat dense icon="sym_o_stop_circle" @click="handleStopInstance(props.row.ID)"/>
                <q-btn color="blue" flat dense icon="sym_o_settings" @click="handleEditSettings(props.row)"/>
                <q-btn color="grey" flat dense icon="sym_o_delete" push>
                  <q-menu>
                    <div class="q-pa-sm">
                      确定要删除？
                      <q-btn dense size="sm" label="确认" color="blue" v-close-popup
                             @click="handleDelete(props.row.ID)"/>
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