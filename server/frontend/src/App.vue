<script setup>
import { ref } from 'vue';
import InstanceList from '@/components/InstanceList.vue';
import ClientInfo from '@/components/ClientInfo.vue';
import InstanceInfo from '@/components/InstanceInfo.vue';

const rightDrawerOpen = ref(false);

const rowProp = ref({});
const sessionIdProp = ref('');

let currentTab = '';
const tabs = {
  'client': ClientInfo,
  'instance': InstanceInfo
};

function showInfo(row, sessionId, type) {
  /**
   * 已开时
   *  不同切换
   *  相同关闭
   * 未开时打开
   */
  if (rightDrawerOpen.value) {
    if (type !== currentTab || rowProp.value !== row) {
      rowProp.value = row;
      currentTab = type;
      sessionIdProp.value = sessionId;
    } else {
      rightDrawerOpen.value = false;
    }
  } else {
    rightDrawerOpen.value = true;
    rowProp.value = row;
    currentTab = type;
    sessionIdProp.value = sessionId;
  }
}
</script>

<template>
  <q-layout view="hHh lpR fFf">
    <q-drawer v-model="rightDrawerOpen" :width="350" side="right" behavior="desktop" elevated>
      <component :is="tabs[currentTab]" :row="rowProp" :sessionId="sessionIdProp" @close="rightDrawerOpen = false"></component>
    </q-drawer>
    <q-page-container>
      <InstanceList @someEvent="showInfo" />
    </q-page-container>
  </q-layout>
</template>

<style scoped></style>
