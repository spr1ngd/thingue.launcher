<script setup>
import { onMounted, ref, reactive, inject } from 'vue';
import { statusToText } from '@/utils';
import { queryAgent, sendInstanceControl, sendPakControl } from '@/api';
import { Notify } from 'quasar';


const rows = ref([])

// const stompClient = inject('stompClient');

const columns = [
  { name: 'id', label: '节点编号', field: 'id', align: 'left' },
  { name: 'hostname', label: '主机名', field: 'hostname', align: 'left' },
  { name: 'ips', label: 'IP地址', field: 'ips', align: 'left' },
  {
    name: 'instanceCount',
    label: '实例数量',
    field: (row) => row.instances.length,
    align: 'left'
  }
];

const subColumns = [
  { name: 'status', label: '状态', field: (row) => statusToText(row.status), align: 'left' },
  { name: 'players', label: '连接数', field: (row) => (row.players ? row.players.length : 0), align: 'left' }
];

function start(id) {
  sendInstanceControl(id, 'START');
}

function stop(id) {
  sendInstanceControl(id, 'STOP');
}

function handleChange(row) {
  let newPak = row.pak;
  if (newPak) {
    sendPakControl({
      nodeName: row.name,
      type: "load",
      pakName: newPak
    }).then((r) => {
      if (r.code === 500) {
        row.pak = '';
      }else{
        row.pak = `${newPak}(切换中)`;
      }
    });
    
  } else {
    sendPakControl({
      nodeName: row.name,
      type: "unload"
    }).then((r) => {
      if (r.code === 500) {
        row.pak = '';
      }else{
        row.pak = `卸载中`;
      }
    });
  }
}

function handleClear(value) {
  if ('卸载中'.includes(value)) {
    Notify.create({ type: 'warning', position: 'top', message: '无效操作' });
  }
  console.log(value);
}

onMounted(async () => {
  const data = await queryAgent()
  console.log(JSON.stringify(data.data.list))
  rows.value = data.data.list;
});
</script>
<template>
  <q-table title="节点列表" :rows="rows" :columns="columns" row-key="sessionId">
    <template v-slot:header="props">
      <q-tr :props="props">
        <q-th auto-width />
        <q-th v-for="col in props.cols" :key="col.name" :props="props">
          {{ col.label }}
        </q-th>
      </q-tr>
    </template>

    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td auto-width>
          <q-btn size="sm" color="primary" round dense @click="props.expand = !props.expand" :icon="props.expand ? 'remove' : 'add'" />
        </q-td>
        <q-td :props="props" :key="props.cols[0].name">
          <q-btn
            flat
            dense
            no-caps
            color="primary"
            :label="props.cols[0].value"
            @click="$emit('someEvent', props.row, props.row.sessionId, 'agent')"
          ></q-btn>
        </q-td>
        <q-td v-for="col in props.cols.slice(1)" :key="col.name" :props="props">
          {{ col.value }}
        </q-td>
      </q-tr>
      <!-- 子列表 -->
      <q-tr v-show="props.expand" :props="props">
        <q-td colspan="100%">
          <q-table hide-pagination :columns="subColumns" :rows="props.row.instances">
            <template v-slot:header="props">
              <q-tr :props="props">
                <q-th align="left">实例编号</q-th>
                <q-th align="left">实例名称</q-th>
                <q-th v-for="col in props.cols" :key="col.name" :props="props">
                  {{ col.label }}
                </q-th>
                <q-th>站点</q-th>
                <q-th auto-width>操作</q-th>
              </q-tr>
            </template>

            <template v-slot:body="props">
              <q-tr :props="props">
                <q-td>{{ props.row.id }}</q-td>
                <q-td>
                  <q-btn flat no-caps dense color="primary" :href="`player.html?name=${props.row.name}`" target="_blank" :label="props.row.name" />
                </q-td>
                <q-td v-for="col in props.cols" :key="col.name" :props="props">
                  {{ col.value }}
                </q-td>
                <q-td auto-width>
                  <div class="q-gutter-md" style="min-width: 135px">
                    <q-select
                      dense
                      options-dense
                      clearable
                      :options="['宜宾换流站', '雁门关换流站', '延庆换流站', '中都换流站']"
                      v-model="props.row.pak"
                      @clear="handleClear"
                      @update:model-value="handleChange(props.row)"
                    />
                  </div>
                </q-td>
                <q-td auto-width>
                  <div class="q-pa-md q-gutter-sm">
                    <q-btn
                      size="sm"
                      color="primary"
                      round
                      dense
                      icon="settings"
                      @click="$emit('someEvent', props.row, props.row.sessionId, 'instance')"
                    ></q-btn>
                    <q-btn size="sm" color="positive" round dense icon="play_arrow" @click="start(props.row.id)"></q-btn>
                    <q-btn size="sm" color="negative" round dense icon="stop" @click="stop(props.row.id)"></q-btn>
                    <q-btn size="sm" color="info" round dense icon="terminal" :href="`grafana/d/a1e60a7c-f226-4614-9300-4efef0d1c62f/thingue?orgId=1&var-node=${props.row.name}&viewPanel=1`" target="_blank"></q-btn>
                  </div>
                </q-td>
              </q-tr>
            </template>
          </q-table>
        </q-td>
      </q-tr>
      <!-- 子列表 -->
    </template>
  </q-table>
</template>
