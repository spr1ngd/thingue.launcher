<script setup>
import {reactive} from 'vue'
import {SelectExePath,ServerStart,ServerShutdown,UnrealStart,UnrealStop} from "../../wailsjs/go/main/App.js";

const localServerConfig = reactive({
  addr: "0.0.0.0:8080",
  basePath: "/"
})

const data = reactive({
  path: "D:/Test/ue4-game/game/Binaries/Win64/game.exe",
})

function select() {
  SelectExePath("").then(result => {
    data.path = result
  })
}

function serverStart() {
  ServerStart(localServerConfig.addr, localServerConfig.basePath);
}

function serverShutdown() {
  ServerShutdown()
}

function unrealStart() {
  UnrealStart(data.path,[""])
}
</script>

<template>
  <div class="q-pa-sm row items-start q-gutter-md">
    <q-card v-if="true">
      <q-card-section>
        <div class="row no-wrap items-center">
          <div class="text-h6">本地信令服务</div>
          <q-icon color="red" name="lens" class="q-ma-sm">
            <q-tooltip>已启动</q-tooltip>
          </q-icon>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-actions vertical>
        <q-list dense>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                绑定地址：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input dense square outlined type="text" style="width: 120px" v-model="localServerConfig.addr"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <div class="text-subtitle1">
                服务路径：
              </div>
            </q-item-section>
            <q-item-section avatar>
              <q-input size="sm" dense square outlined style="width: 120px;" type="text" v-model="localServerConfig.basePath"/>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section>
              <q-btn dense label="启动" color="positive" @click="serverStart"></q-btn>
            </q-item-section>
            <q-item-section>
              <q-btn dense label="停止" color="negative" @click="serverShutdown"></q-btn>
            </q-item-section>
            <q-item-section avatar>
              <q-btn flat round icon="open_in_new"/>
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