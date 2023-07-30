<script setup>
import {reactive, ref, computed} from "vue";
import {SelectExePath, UnrealStart} from "@wails/go/agent/App.js";

const data = reactive({
  name: "test11",
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

const rows = reactive([
  {
    name: "test1",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test2",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test3",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test4",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test5",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test6",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test7",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test8",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test9",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test10",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }, {
    name: "test11",
    path: "D:/Test/ue4-game/game/Binaries/Win64/127.0.0.1:8080/ws/streamer/abcd/game.exe",
  }])

const settingsDialog = ref(false)

const columns = [
  {name: 'name', field: 'name', label: '实例ID:'},
  {name: 'path', field: 'path', label: '路径:'},
]

const options = [
  '本地1', '127.0.0.1:8080', '127.0.0.1:8081', '127.0.0.1:8082', '127.0.0.1:8083'
]

const selected = ref("")

const pagination = ref({
  sortBy: 'desc',
  descending: false,
  page: 2,
  rowsPerPage: 4
  // rowsNumber: xx if getting data from a server
})

const pagesNumber = computed(() => Math.ceil(rows.length / pagination.value.rowsPerPage))


</script>

<template>
  <div class="q-pa-sm">
    <q-table grid title="ThingUE列表" :rows="rows" :columns="columns" v-model="selected"
             selection="multiple" hide-pagination :pagination="{rowsPerPage:0}">
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
          <q-card class="q-pt-md q-pb-sm">
            <!--            <q-card-section>-->
            <!--              <q-checkbox dense v-model="props.selected" :label="props.row.name"/>-->
            <!--            </q-card-section>-->
            <!--            <q-separator/>-->
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
      <q-card-section class="q-pa-sm">
        <div class="row no-wrap items-center q-pa-sm">
          <div class="text-h6">实例配置</div>
          <q-space/>
          <q-btn flat color="primary">保存</q-btn>
          <q-btn flat>关闭</q-btn>
        </div>
      </q-card-section>
<!--      <q-separator/>-->
      <q-card-section class="q-pa-none q-pt-sm">
        <q-list dense>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1 q-ml-sm">实例ID：</div>
            </q-item-section>
            <q-item-section side class="q-pl-none">
              <q-input dense outlined v-model="data.name" style="width: 405px"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1 q-ml-sm">路径：</div>
            </q-item-section>
            <q-item-section side>
              <q-input dense outlined v-model="data.path" style="width: 405px">
                <template v-slot:append>
                  <q-icon name="sym_o_file_open" @click="select" class="cursor-pointer"/>
                </template>
              </q-input>
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-section>
      <q-card-actions/>
    </q-card>
  </q-dialog>
</template>

<style scoped>

</style>