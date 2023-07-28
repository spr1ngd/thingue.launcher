<script setup>
import {reactive} from "vue";
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

const items = [1, 2, 3, 4, 5]
</script>

<template>
    <q-card>
<!--      <q-card-actions>-->
<!--        <q-btn size="sm" dense round color="primary" icon="add"/>-->
<!--      </q-card-actions>-->
      <q-card-section>
        <q-list>
          <q-item-label header>ThingUE</q-item-label>
          <q-item v-for="i in items">
            <q-input dense square outlined v-model="data.path" style="width: 500px">
              <template v-slot:append>
                <q-btn flat round icon="file_open" @click="select"/>
              </template>
            </q-input>
          </q-item>
        </q-list>
      </q-card-section>
    </q-card>
</template>

<style scoped>

</style>