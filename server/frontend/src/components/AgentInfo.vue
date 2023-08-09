<script setup>
import { computed } from 'vue'
import { bytesToSize } from '@/utils';
const props = defineProps(['row']);

const humanMembers = computed(() => {
  return bytesToSize(props.row.totalmem)
})
</script>
<template>
  <div class="q-pa-md q-gutter-md">
    <q-list>
      <q-item-label header>
        <div class="text-h5">Agent详细</div>
      </q-item-label>
      <q-item>
        <q-item-section>
          <q-item-label>节点ID：</q-item-label>
          <q-item-label caption>{{ props.row.nodeId }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>Agent版本：</q-item-label>
          <q-item-label caption>{{ props.row.version }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>安装路径：</q-item-label>
          <q-item-label caption>{{ props.row.path }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>CPU：</q-item-label>
          <q-item-label caption>{{ props.row.cpus }}核</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>内存：</q-item-label>
          <q-item-label caption>{{ humanMembers }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统类型：</q-item-label>
          <q-item-label caption>{{ props.row.osType }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统版本：</q-item-label>
          <q-item-label caption>{{ props.row.osVersion }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统架构：</q-item-label>
          <q-item-label caption>{{ props.row.osArch }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统用户：</q-item-label>
          <q-item-label caption>{{ props.row.osUser }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>IP地址：</q-item-label>
          <q-item-label caption>{{ props.row.ips }}</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
    <q-btn color="white" text-color="primary" label="关闭" @click="$emit('close')" />
  </div>
</template>
