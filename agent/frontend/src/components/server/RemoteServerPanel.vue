<script setup>
import {reactive} from "vue";

const remoteServerConfig = reactive({
  addr: "127.0.0.1:8888",
  enable: false
})

const a = [1, 2, 3, 4, 5]

</script>

<template>
  <div class="row no-wrap items-center q-pa-sm">
    <div class="text-subtitle1 q-ml-sm">服务地址列表</div>
    <q-space/>
    <q-btn dense size="sm" color="primary" round icon="add"/>
  </div>
  <q-separator/>
  <q-list separator>
    <q-item v-for="ai in a">
      <q-item-section>
        <q-input dense standout="bg-grey" type="text"
                 v-model="remoteServerConfig.addr">
          <template v-slot:after>
            <q-btn class="gt-xs" size="12px" flat dense round icon="sym_o_delete"/>
          </template>
        </q-input>
      </q-item-section>
    </q-item>
  </q-list>
</template>

<style scoped>

</style>