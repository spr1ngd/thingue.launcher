<script setup>
import {reactive} from "vue";
import {SelectExePath, UnrealStart} from "../../wailsjs/go/main/App.js";

const data = reactive({
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
})

function select() {
  SelectExePath("").then(result => {
    data.path = result
  })
}

function unrealStart() {
  UnrealStart(data.path,[
    "-AudioMixer",
    "-RenderOffScreen",
    "-ForceRes",
    "-ResX=1920",
    "-ResX=1080",
    "-PixelStreamingURL=ws://127.0.0.1:8080/ws/streamer/abcd",
  ])
}
</script>

<template>
  <div class="q-pa-sm row items-start q-gutter-md">
    <q-card v-if="true">
      <q-card-section>
        <div class="text-h6">ThingUE</div>
      </q-card-section>
      <q-separator/>
      <q-card-actions vertical>
        <q-list dense>
          <q-item>
            <q-item-label>
              <div class="text-subtitle2"></div>
            </q-item-label>
            <q-item-section>
              <q-input dense square outlined v-model="data.path" style="width: 500px">
                <template v-slot:after>
                  <q-btn flat round icon="file_open" @click="select"/>
                </template>
              </q-input>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section avatar>
              <q-btn dense label="启动" color="primary" @click="unrealStart"></q-btn>
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-actions>
      <q-separator />
    </q-card>
  </div>
</template>

<style scoped>

</style>