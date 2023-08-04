<script setup>
import {defineEmits} from "vue";
import {OpenFileDialog} from "@wails/go/app/App.js";
import {Notify} from "quasar";

const emit = defineEmits(['openListPanel'])
const props = defineProps(['data']);

function select() {
  OpenFileDialog("选择文件", "ThingUE (*.exe)", "*.exe").then(result => {
    props.data.path = result
  }).catch(err => {
    // todo
    // Notify.create({
    //   message: '取消选择'
    // })
  })
}
</script>

<template>
  <q-card>
    <q-card-section class="q-pa-sm">
      <div class="row no-wrap items-center q-pa-sm">
        <div class="text-h6">实例配置</div>
        <q-space/>
        <div class="q-gutter-md">
          <q-btn color="primary">保存</q-btn>
          <q-btn @click="emit('openListPanel')">关闭</q-btn>
        </div>
      </div>
    </q-card-section>
    <q-card-section class="q-pa-none q-pt-sm">
      <q-list>
        <q-item>
          <q-item-section avatar>
            <q-item-label>实例ID</q-item-label>
            <q-input dense outlined square v-model="props.data.name"/>
          </q-item-section>
        </q-item>
        <q-item>
          <q-item-section>
            <q-item-label>路径</q-item-label>
            <q-input dense outlined square v-model="props.data.path">
              <template v-slot:append>
                <q-icon name="sym_o_file_open" @click="select" class="cursor-pointer"/>
              </template>
            </q-input>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card-section>
    <q-card-actions/>
  </q-card>
</template>