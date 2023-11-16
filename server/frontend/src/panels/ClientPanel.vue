<script setup>
import {collectLogs, downloadLogs} from "@/api";
import {Notify, useQuasar} from "quasar";
import {v4 as uuidv4} from 'uuid';
import {emitter, wsId} from "@/ws";
import {usePanelStore} from "@/stores";

const panelStore = usePanelStore();

const props = defineProps(['data']);
const $q = useQuasar();

async function handleCollectClick() {
  const traceId = uuidv4()
  let res = await collectLogs({
    wsId: wsId,
    clientId: props.data.id,
    traceId
  });
  if (res.code === 200) {
    const dialog = $q.dialog({
      title: '提示',
      message: '正在收集，请稍等',
      progress: true,
      persistent: true,
      ok: false
    })
    emitter.on('downloadComplete', () => {
      dialog.update({
        message: '收集完成，请在60秒内完成下载',
        progress: false,
        ok: true
      }).onOk(() => {
        doLogsDownload(traceId);
      })
    })
  }
}

function doLogsDownload(traceId) {
  downloadLogs(traceId).then((response) => {
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
          <q-item-label caption>{{ props.data.version }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>客户端安装路径：</q-item-label>
          <q-item-label caption>{{ props.data.workdir }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>CPU：</q-item-label>
          <q-item-label caption v-for="cpu in props.data.cpus">{{ cpu }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>GPU：</q-item-label>
          <q-item-label caption v-for="gpu in props.data.gpus">{{ gpu }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>内存：</q-item-label>
          <q-item-label caption>{{ props.data.memory }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统类型：</q-item-label>
          <q-item-label caption>{{ props.data.osType }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统架构：</q-item-label>
          <q-item-label caption>{{ props.data.osArch }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>系统用户：</q-item-label>
          <q-item-label caption>{{ props.data.systemUser }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item>
        <q-item-section>
          <q-item-label>IP地址：</q-item-label>
          <q-item-label caption v-for="ip in props.data.ips">{{ ip }}</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
    <q-btn color="white" text-color="primary" label="收集客户端日志" @click="handleCollectClick"/>
    <q-btn color="white" text-color="primary" label="关闭" @click="panelStore.closePanel()"/>
  </div>
</template>
