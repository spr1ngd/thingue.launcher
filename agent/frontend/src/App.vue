<script setup>
import {ref} from 'vue'
import Server from '@/components/server/Server.vue'
import Unreal from '@/components/Unreal.vue'
import Setting from "@/components/Setting.vue";


const tab = ref("unreal")
const splitterModel = ref(10)
</script>

<template>
  <div>
    <q-splitter v-model="splitterModel" style="height: 100vh;">
      <template v-slot:before>
        <q-tabs v-model="tab" vertical class="text-primary shadow-2">
          <q-tab name="unreal" icon="svguse:ue-logo-white.svg#icon-1" label="unreal"/>
          <q-tab name="server" icon="sym_o_cloud" label="服务"/>
          <q-tab name="settings" icon="sym_o_settings" label="设置"/>
        </q-tabs>
      </template>

      <template v-slot:after>
        <q-tab-panels
            v-model="tab"
            animated
            swipeable
            vertical
            transition-prev="jump-up"
            transition-next="jump-up"
        >
          <q-tab-panel name="unreal" class="q-pa-none">
            <!--            <div class="text-h4 q-mb-md">Mails</div>-->
            <Unreal/>
          </q-tab-panel>

          <q-tab-panel name="server" class="q-pa-none">
            <!--            <div class="text-h4 q-mb-md">Alarms</div>-->
            <Server/>
          </q-tab-panel>

          <q-tab-panel name="settings">
            <!--            <div class="text-h4 q-mb-md">Movies</div>-->
            <Setting/>
          </q-tab-panel>
        </q-tab-panels>
      </template>

    </q-splitter>
  </div>
</template>

<style scoped>
</style>
