<script setup>
import {defineEmits, onMounted, ref} from "vue";
import {DeleteInstance, ListInstance} from "@wails/go/unreal/Unreal.js";
import {OpenExplorer} from "@wails/go/app/App.js";

const emit = defineEmits(["openSettingsPanel"])

const rows = ref([])

const columns = [
  {name: 'Name', field: 'Name', label: '实例标识'},
  {name: 'ExecPath', field: 'ExecPath', label: '启动位置'},
]

const options = [
  '本地1', '127.0.0.1:8080', '127.0.0.1:8081', '127.0.0.1:8082', '127.0.0.1:8083'
]

const stations = [
  '宜宾换流站', '延庆换流站', '中都换流站'
]

const selected = ref(null)

onMounted(async () => {
  await list()
})

async function list() {
  let instances = await ListInstance();
  rows.value = instances
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
        "-PixelStreamingURL=ws://127.0.0.1:8080/ws/streamer/abcd",
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

async function handleDelete(id) {
  await DeleteInstance(id)
  await list()
}

async function handleOpenDir(path) {
  await OpenExplorer(path)
}


</script>

<template>
  <div class="q-pa-sm">
    <q-table grid title="实例列表" :rows="rows" :columns="columns" v-model="selected"
             selection="multiple" hide-pagination :pagination="{rowsPerPage:0}" hide-no-data>
      <template v-slot:top-right>
        <q-select size="sm" dense :options="options" options-dense v-model="selected"/>
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
                  <q-item-section avatar style="width: 100px">
                    <q-item-label caption class="ellipsis">标识</q-item-label>
                    <q-item-label class="ellipsis">{{ props.row.Name }}</q-item-label>
                  </q-item-section>
                  <q-item-section avatar style="width: 70px">
                    <q-item-label caption class="ellipsis">状态</q-item-label>
                    <q-item-label class="ellipsis">已启动</q-item-label>
                  </q-item-section>
                  <q-item-section>
<!--                    <q-item-label caption class="ellipsis">Pak资源加载</q-item-label>-->
<!--                    <q-select v-model="selected" :options="stations" label="Standard" />-->
                    <q-select dense :options="stations" options-dense clearable label="资源加载" v-model="selected"/>
                  </q-item-section>
                </q-item>
                <q-item>
                  <q-item-section>
                    <q-item-label caption class="ellipsis cursor-pointer" @click="handleOpenDir(props.row.ExecPath)">启动位置</q-item-label>
                    <q-item-label class="ellipsis cursor-pointer" @click="handleOpenDir(props.row.ExecPath)">{{ props.row.ExecPath }}</q-item-label>
                  </q-item-section>
                </q-item>
              </q-list>
            </q-card-section>
            <q-card-actions class="q-pt-none">
              <div class="q-gutter-md">
                <q-btn color="green" flat dense icon="sym_o_play_circle"/>
                <q-btn color="red" flat dense icon="sym_o_stop_circle"/>
                <q-btn color="blue" flat dense icon="sym_o_settings" @click="handleEditSettings(props.row)"/>
                <q-btn color="grey" flat dense icon="sym_o_delete" @click="handleDelete(props.row.ID)"/>
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