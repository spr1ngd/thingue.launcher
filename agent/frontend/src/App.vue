<script setup>
import {ref} from 'vue'
import Server from '@/components/server/Main.vue'
import Unreal from '@/components/unreal/Main.vue'
import Setting from "@/components/Setting.vue";


const tab = ref("unreal")
</script>

<template>
  <div>
    <div class="row no-wrap">
      <div style="height: 100vh;border-right: 1px solid rgba(0, 0, 0, 0.12);">
        <q-tabs v-model="tab" vertical class="text-primary shadow-0">
          <q-tab name="unreal" icon="svguse:ue-logo-white.svg#icon-1" label="unreal"/>
          <q-tab name="server" icon="sym_o_cloud" label="服务"/>
          <q-tab name="settings" icon="sym_o_settings" label="设置"/>
        </q-tabs>
      </div>
      <div class="col-grow bi-border-left" style="height: 100vh">
        <q-tab-panels
            v-model="tab"
            animated
            swipeable
            vertical
            transition-prev="jump-up"
            transition-next="jump-up"
        >
          <q-tab-panel name="unreal" class="q-pa-none">
            <Unreal/>
          </q-tab-panel>
          <q-tab-panel name="server" class="q-pa-none">
            <Server/>
          </q-tab-panel>
          <q-tab-panel name="settings">
            <Setting/>
          </q-tab-panel>
        </q-tab-panels>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
