<script setup>
import {usePanelStore} from "@/stores"

let panelStore = usePanelStore();

function clientPanel() {
  panelStore.togglePanel("clientPanel", {val: 1})
}

function instancePanel() {
  panelStore.togglePanel("instancePanel", {val: 2}, 380)
}
</script>
<template>
  <div>HomeView</div>
  <q-btn @click="clientPanel">clientPanel</q-btn>
  <q-space/>
  <q-btn @click="instancePanel">instancePanel</q-btn>
</template>