<script setup>
import {onMounted, ref} from "vue";
import {CreateRemoteServer, DeleteRemoteServer, ListRemoteServer, SaveRemoteServer} from "@wails/go/api/serverApi";

const data = ref([])

onMounted(async () => {
  await list()
})

async function list() {
  data.value = await ListRemoteServer()
}

async function save(obj) {
  await SaveRemoteServer(obj)
  await list()
}

async function create() {
  await CreateRemoteServer({})
  await list()
}

async function del(id) {
  await DeleteRemoteServer(id)
  await list()
}

</script>

<template>
  <div class="row no-wrap items-center q-pa-sm">
    <div class="text-subtitle1 q-ml-sm">服务地址列表</div>
    <q-space/>
    <q-btn dense size="sm" color="primary" round icon="add" @click="create()"/>
  </div>
  <q-separator/>
  <q-list separator>
    <q-item v-for="row in data">
      <q-item-section>
        <q-input dense standout="bg-grey" type="text"
                 v-model="row.Url" @change="save(row)">
          <template v-slot:after>
            <q-btn class="gt-xs" size="12px" flat dense round icon="sym_o_delete">
              <q-menu>
                <div class="q-pa-sm">
                  确定要删除？
                  <q-btn dense size="sm" label="确认" color="blue" v-close-popup  @click="del(row.ID)"/>
                </div>
              </q-menu>
            </q-btn>
          </template>
        </q-input>
      </q-item-section>
    </q-item>
  </q-list>
</template>

<style scoped>

</style>