<script setup>
import {onMounted, ref} from "vue";
import {getPaksName, processStateToText} from "@/utils";
import {controlProcess, getInstanceList, sendPakControl} from "@/api";
import {usePanelStore} from "@/stores";
import {Notify} from "quasar";

const panelStore = usePanelStore();

const rows = ref([])

const columns = [
  {name: 'status', label: '进程状态', field: (row) => processStateToText(row.stateCode), align: 'center'},
  {name: 'status', label: 'Streamer状态', field: (row) => row.streamerConnected ? '已连接' : '未连接', align: 'center'},
  {name: 'players', label: '连接数', field: (row) => (row.playerIds ? row.playerIds.length : 0), align: 'center'},
];

async function queryData() {
  const res = await getInstanceList()
  rows.value = res.data.list
}

function start(sid) {
  controlProcess(sid, 'START');
}

function stop(sid) {
  controlProcess(sid, 'STOP');
}

function handleChange(row) {
  let newPak = row.pakName;
  if (newPak) {
    sendPakControl({
      sid: row.sid,
      type: "load",
      pak: row.paks.filter(pak => pak.name === newPak)[0].value
    }).then((r) => {
      if (r.code === 500) {
        row.pakName = '';
      } else {
        row.pakName = `${newPak}(切换中)`;
      }
    });
  } else {
    sendPakControl({
      sid: row.sid,
      type: "unload"
    }).then((r) => {
      if (r.code === 500) {
        row.pakName = '';
      } else {
        row.pakName = `卸载中`;
      }
    });
  }
}

function handleClear(value) {
  if ('卸载中'.includes(value)) {
    Notify.create({type: 'warning', position: 'top', message: '无效操作'});
  }
  console.log(value);
}

onMounted(() => {
  queryData()
})
</script>
<template>
  <q-table title="实例列表" :rows="rows" :columns="columns" row-key="id" class="q-ma-md">
    <template v-slot:header="props">
      <q-tr :props="props">
        <q-th>实例编号</q-th>
        <q-th>实例名称</q-th>
        <q-th v-for="col in props.cols" :key="col.name" :props="props">
          {{ col.label }}
        </q-th>
        <q-th>站点</q-th>
        <q-th auto-width>操作</q-th>
      </q-tr>
    </template>
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td>{{ props.row.cid }}</q-td>
        <q-td>
          <q-btn padding="none" flat no-caps dense color="primary" :href="`player.html?sid=${props.row.sid}`"
                 target="_blank" :label="props.row.name"/>
        </q-td>
        <q-td v-for="col in props.cols" :key="col.name" :props="props">
          {{ col.value }}
        </q-td>
        <q-td auto-width>
          <div class="q-gutter-md" style="min-width: 135px">
            <q-select dense options-dense clearable :options="getPaksName(props.row.paks)"
                      v-model="props.row.pakName" @clear="handleClear"
                      @update:model-value="handleChange(props.row)">
              <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-italic text-grey">
                    <span>当前实例非壳加载模</span>
                    <span>式 ，没有可用选项</span>
                  </q-item-section>
                </q-item>
              </template>
            </q-select>
          </div>
        </q-td>
        <q-td auto-width>
          <div class="q-pa-md q-gutter-sm">
            <q-btn size="sm" color="primary" round dense icon="settings"
                   @click="panelStore.togglePanel('instancePanel', props.row)"></q-btn>
            <q-btn size="sm" color="positive" round dense icon="play_arrow"
                   @click="start(props.row.sid)"></q-btn>
            <q-btn size="sm" color="negative" round dense icon="stop" @click="stop(props.row.sid)"></q-btn>
          </div>
        </q-td>
      </q-tr>
    </template>
  </q-table>
</template>