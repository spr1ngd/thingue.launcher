<script setup>
import {reactive, ref} from "vue";
import {SelectExePath, UnrealStart} from "@wails/go/agent/App.js";

const data = reactive({
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
})

function select() {
  SelectExePath("").then(result => {
    data.path = result
  })
}

function unrealStart() {
  UnrealStart(data.path, [
    "-AudioMixer",
    "-RenderOffScreen",
    "-ForceRes",
    "-ResX=1920",
    "-ResX=1080",
    "-PixelStreamingURL=ws://127.0.0.1:8080/ws/streamer/abcd",
  ])
}

const rows = reactive([{
  name: "test1",
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
}, {
  name: "test2",
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
}, {
  name: "test3",
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
}])

const settingsDialog = ref(false)

const columns = [
  {name: 'name', field: 'name', label: '实例ID:'},
  {name: 'path', field: 'path', label: '路径:'},
]

const options = [
  '本地', '127.0.0.1:8080', '127.0.0.1:8081', '127.0.0.1:8082', '127.0.0.1:8083'
]

const selected = ref("")
</script>

<template>
  <div class="q-pa-sm">
    <q-table grid title="ThingUE列表" :rows="rows" :columns="columns" hide-pagination v-model="selected"
             selection="multiple">
      <template v-slot:top-right>
        <q-select size="sm" dense :options="options" options-dense v-model="selected"/>
        <q-space/>
        <q-btn dense size="sm" color="primary" round icon="add"/>
      </template>
      <template v-slot:item="props">
        <div
            class="q-pa-sm col-xs-12 col-sm-6 col-md-4 col-lg-3 grid-style-transition"
            :style="props.selected ? 'transform: scale(0.95);' : ''"
        >
          <q-card class="q-pa-sm">
            <!--            <q-card-section>-->
            <!--              <q-checkbox dense v-model="props.selected" :label="props.row.name"/>-->
            <!--            </q-card-section>-->
            <!--            <q-separator/>-->
            <q-list dense>
              <q-item v-for="col in props.cols" :key="col.name">
                <q-item-section>
                  <q-item-label>{{ col.label }}</q-item-label>
                  <q-item-label caption>{{ col.value }}</q-item-label>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section thumbnail>
                  <div>
                    <q-btn color="green" flat dense icon="sym_o_play_circle"/>
                    <q-btn color="red" flat dense icon="sym_o_stop_circle"/>
                    <q-btn color="blue" flat dense icon="sym_o_settings" @click="settingsDialog=true"/>
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
  <q-dialog v-model="settingsDialog">
    <q-card>
      <q-list separator>
        <q-item>
          <q-item-section side>
            <q-input dense outlined v-model="data.path" style="width: 500px">
              <template v-slot:append>
                <q-icon name="sym_o_file_open" @click="select" class="cursor-pointer"/>
              </template>
            </q-input>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card>
  </q-dialog>
</template>

<style scoped>

</style>