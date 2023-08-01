<script setup>
import InstanceList from "@/components/unreal/InstanceList.vue";
import InstanceInfo from "@/components/unreal/InstanceInfo.vue";
import {ref} from "vue";

const panel = ref("list")
</script>

<template>
  <q-tab-panels v-model="panel" animated class="shadow-2 rounded-borders">
    <q-tab-panel name="list">
      <instance-list/>
    </q-tab-panel>

    <q-tab-panel name="info">
      <instance-info/>
    </q-tab-panel>
  </q-tab-panels>
</template>

<style scoped>

</style>