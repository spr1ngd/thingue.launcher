<script setup>
import {defineEmits, onMounted, reactive, ref} from "vue";
import {ListInstance} from "@wails/go/unreal/Unreal.js";

const emit = defineEmits(["openSettingsPanel"])

const rows = ref([])

const columns = [
  {name: 'name', field: 'name', label: '实例ID:'},
  {name: 'path', field: 'path', label: '路径:'},
]

const options = [
  '本地1', '127.0.0.1:8080', '127.0.0.1:8081', '127.0.0.1:8082', '127.0.0.1:8083'
]

const selected = ref("")

onMounted(async () => {
  let instances = await ListInstance();
  // rows.value = instances
})

function handleNewSettings() {
  emit("openSettingsPanel", {
    type: 'new',
    settings: {
      params: [
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


</script>

<template>
  <div class="q-pa-sm">
    <q-table grid title="ThingUE列表" :rows="rows" :columns="columns" v-model="selected"
             selection="multiple" hide-pagination :pagination="{rowsPerPage:0}">
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
          <q-card class="q-pt-md q-pb-sm">
            <q-list dense>
              <q-item v-for="col in props.cols" :key="col.name">
                <q-item-section>
                  <q-item-label class="ellipsis">{{ col.label }}</q-item-label>
                  <!--                  <q-item-label caption class="text-no-wrap overflow-auto hide-scrollbar">{{ col.value }}</q-item-label>-->
                  <q-item-label caption class="ellipsis">{{ col.value }}</q-item-label>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <div>
                    <q-btn color="green" flat dense icon="sym_o_play_circle"/>
                    <q-btn color="red" flat dense icon="sym_o_stop_circle"/>
                    <q-btn color="blue" flat dense icon="sym_o_settings" @click="handleEditSettings(props.row)"/>
                    <q-btn color="grey" flat dense icon="sym_o_delete"/>
                  </div>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card>
        </div>
      </template>
    </q-table>
  </div>
</template>

<style scoped>

</style>