<script setup>
import UnrealListPanel from "@/components/unreal/UnrealListPanel.vue";
import UnrealSettingPanel from "@/components/unreal/UnrealSettingsPanel.vue";
import {ref} from "vue";

const currentPanel = ref("list")

const settingsData = ref({})

const emit = defineEmits(["gotoServer"])

function switchSettingsPanel(data) {
  settingsData.value = data
  currentPanel.value = "setting"
}

function switchListPanel() {
  currentPanel.value = "list"
}

function gotoServer(tab) {
  emit("gotoServer", tab)
}

</script>

<template>
  <q-tab-panels v-model="currentPanel" animated class="shadow-2">
    <q-tab-panel name="list" class="q-pa-none">
      <unreal-list-panel @open-settings-panel="switchSettingsPanel" @goto-server="gotoServer"/>
    </q-tab-panel>
    <q-tab-panel name="setting">
      <unreal-setting-panel @open-list-panel="switchListPanel" :data="settingsData"/>
    </q-tab-panel>
  </q-tab-panels>
</template>

<style scoped>

</style>