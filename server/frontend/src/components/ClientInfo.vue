<script setup>
import {collectLogs} from "@/api";
import {Notify} from "quasar";

const props = defineProps(['row']);

async function handleCollectClick() {
  collectLogs(props.row.id).then((response) => {
    if (response.data.type === 'application/zip') {
      let filename = 'logs.zip'; // 默认文件名
      const contentDisposition = response.headers['content-disposition'];
      if (contentDisposition) {
        const regex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/;
        const matches = regex.exec(contentDisposition);
        if (matches != null && matches[1]) {
          filename = decodeURIComponent(matches[1].replace(/['"]/g, ''));
        }
      }
      const blob = new Blob([response.data], {type: 'application/octet-stream'});
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', filename);
      document.body.appendChild(link);
      link.click();
    } else {
      response.data.text().then(data => {
        let json = JSON.parse(data);
        Notify.create({type: 'warning', position: 'top', message: json.msg});
      })
    }
  }).catch((error) => {
    console.error('文件下载失败：', error);
  });
}

</script>
<template>
  <div class="q-pa-md q-gutter-md">
    <q-list>
      <q-item-label header>
        <div class="text-h5">客户端信息</div>
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
    <q-btn color="white" text-color="primary" label="收集实例日志" @click="handleCollectClick"/>
    <q-btn color="white" text-color="primary" label="关闭" @click="$emit('close')"/>
  </div>
</template>
