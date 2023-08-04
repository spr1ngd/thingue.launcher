<script setup>
import {ref} from 'vue'
import LocalServer from "@/components/server/LocalServerPanel.vue";
import RemoteServer from "@/components/server/RemoteServerPanel.vue";

const tab = ref("local")

</script>

<template>
  <q-tabs
      v-model="tab"
      dense
      align="left"
      active-color="primary"
      indicator-color="primary"
      narrow-indicator
  >
    <q-tab name="local" label="本地"/>
    <q-tab name="remote" label="远程"/>
  </q-tabs>
  <q-separator/>
  <q-tab-panels
      v-model="tab"
      animated
      swipeable
      transition-prev="jump-up"
      transition-next="jump-up"
  >
    <q-tab-panel name="local" class="q-pa-md">
      <local-server/>
    </q-tab-panel>
    <q-tab-panel name="remote" class="q-pa-sm">
      <remote-server/>
    </q-tab-panel>
  </q-tab-panels>
</template>

<style scoped>

</style>