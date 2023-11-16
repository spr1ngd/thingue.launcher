<script setup>
import {onMounted, ref} from "vue";
import {getClientList} from "@/api";
import {usePanelStore} from "@/stores";

const panelStore = usePanelStore();

const rows = ref([])
const columns = [
  {name: 'id', label: '客户端编号', field: 'id', align: 'center'},
  {name: 'hostname', label: '主机名', field: 'hostname', align: 'center'},
  {name: 'ips', label: 'IP地址', field: 'ips', align: 'center'},
  {
    name: 'instanceCount',
    label: '实例数量',
    field: (row) => row.instances.length,
    align: 'center'
  },
];

async function queryData() {
  const data = await getClientList()
  rows.value = data.data.list;
}

onMounted(() => {
  queryData()
})
</script>
<template>
  <q-table title="客户端列表" :rows="rows" :columns="columns" row-key="id" class="q-ma-md">
    <template v-slot:header="props">
      <q-tr :props="props">
        <q-th v-for="col in props.cols" :key="col.name" :props="props">
          {{ col.label }}
        </q-th>
        <q-th/>
      </q-tr>
    </template>
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td v-for="col in props.cols" :key="col.name" :props="props">
          {{ col.value }}
        </q-td>
        <q-td :align="'center'">
          <q-btn size="sm" dense color="primary" @click="panelStore.togglePanel('clientPanel', props.row)">
            更多信息
          </q-btn>
        </q-td>
      </q-tr>
    </template>
  </q-table>
</template>