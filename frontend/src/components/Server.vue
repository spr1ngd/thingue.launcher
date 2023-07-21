<script setup>
import {reactive} from 'vue'
import {SelectExePath} from "../../wailsjs/go/main/App.js";

const data = reactive({
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
})

function open() {
  SelectExePath("").then(result => {
    data.path = result
  })
}
</script>

<template>
  <div class="q-pa-sm row items-start q-gutter-md">
    <q-input dense square outlined v-model="data.path"></q-input>
    <q-card v-if="true">
      <q-card-section>
        <div class="row no-wrap items-center">
          <div class="text-h6">本地信令服务</div>
<!--          <q-icon color="red" name="lens" class="q-ma-sm">-->
<!--            <q-tooltip>已启动</q-tooltip>-->
<!--          </q-icon>-->
        </div>
      </q-card-section>
      <q-separator />
      <q-card-actions vertical>
        <q-list dense>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                服务端口：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input dense square outlined type="number" style="width: 120px"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                服务路径：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input size="sm" dense square outlined style="width: 120px;" type="text"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <q-btn label="启动" color="primary"></q-btn>
            </q-item-section>
            <q-item-section avatar>
              <q-btn flat round icon="open_in_new" />
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-actions>
    </q-card>
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
              <q-input dense square outlined v-model="data.path"></q-input>
            </q-item-section>
            <q-item-section avatar>
              <q-btn flat round icon="file_open" @click="open"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <q-btn label="启动" color="primary"></q-btn>
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