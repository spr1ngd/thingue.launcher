<script setup>
import {deleteCloudRes, listCloudRes} from "@/api";
import {onMounted, ref} from "vue";
import {usePanelStore} from "@/stores";

const panelStore = usePanelStore();

const rows = ref([])
const selected = ref([])

const columns = [
  {name: 'name', label: '资源名称', field: 'name', align: 'left'},
  {name: 'lastUpdateAt', label: '最后更新时间', field: 'lastUpdateAt', align: 'left'},
];

async function queryData() {
  const res = await listCloudRes()
  rows.value = res.data
}

async function del() {
  await deleteCloudRes(selected.value.map(select => select.name))
  await queryData()
}

onMounted(() => {
  queryData()
})
</script>
<template>
  <q-table title="同步资源列表" :rows="rows" :columns="columns" row-key="name" class="q-ma-md" selection="multiple"
           v-model:selected="selected">
    <template v-slot:top-left>
      <div class="full-width row flex-center q-gutter-sm">
        <span class="text-h6">同步资源列表</span>
      </div>
    </template>
    <template v-slot:top-right>
      <div class="full-width row flex-center q-gutter-sm">
        <q-btn size="sm" color="primary" round icon="add" @click="panelStore.togglePanel('cloudResPanel',{},380)"/>
        <q-btn size="sm" color="primary" round icon="delete" @click="del"/>
      </div>
    </template>
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td>
          <q-checkbox v-model="props.selected"/>
        </q-td>
        <q-td key="name" :props="props">
          <q-btn padding="none" flat no-caps dense color="primary" @click="panelStore.togglePanel('cloudResPanel',props.row,380)" :label="props.row.name"/>
        </q-td>
        <q-td key="lastUpdateAt" :props="props">{{ props.row.lastUpdateAt }}</q-td>
      </q-tr>
    </template>
  </q-table>
</template>