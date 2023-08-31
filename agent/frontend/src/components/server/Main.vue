<script setup>
import LocalServer from "@/components/server/LocalServerPanel.vue";
import RemoteServer from "@/components/server/RemoteServerPanel.vue";
import {defineProps, onMounted, ref, watch} from "vue";

const tab = ref("local")

const props = defineProps(['tab']);

const emits = defineEmits(['changeTab']);

onMounted(() => {
  tab.value = props.tab
})

watch(tab, () => {
  emits("changeTab", tab.value)
})

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
    <q-tab name="local" label="内置"/>
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