<script setup>
import { computed } from 'vue'
import { bytesToSize } from '@/utils';
const props = defineProps(['row']);

const humanMembers = computed(() => {
  return bytesToSize(props.row.totalmem)
})

function handleCollectClick() {}
</script>
<template>
  <div class="q-pa-md q-gutter-md">
    <q-list>
      <q-item-label header>
        <div class="text-h5">节点详细信息</div>
      </q-item-label>
      <q-item>
        <q-item-section>
          <q-item-label>客户端版本：</q-item-label>
          <q-item-label caption>{{ props.row.version }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>客户端安装路径：</q-item-label>
          <q-item-label caption>{{ props.row.workdir }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>CPU：</q-item-label>
          <q-item-label caption v-for="cpu in props.row.cpus">{{ cpu }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>GPU：</q-item-label>
          <q-item-label caption v-for="gpu in props.row.gpus">{{ gpu }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>内存：</q-item-label>
          <q-item-label caption>{{ props.row.memory }}</q-item-label>
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
          <q-item-label>系统架构：</q-item-label>
          <q-item-label caption>{{ props.row.osArch }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统用户：</q-item-label>
          <q-item-label caption>{{ props.row.systemUser }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>IP地址：</q-item-label>
          <q-item-label caption v-for="ip in props.row.ips">{{ ip }}</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
    <q-btn color="white" text-color="primary" label="收集节点日志" @click="handleCollectClick" />
    <q-btn color="white" text-color="primary" label="关闭" @click="$emit('close')" />
  </div>
</template>
