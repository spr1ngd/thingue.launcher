<script setup>
import {ref} from 'vue';
import {RouterView} from 'vue-router'
import {usePanelStore} from "@/stores"

let panelStore = usePanelStore();

const leftDrawerOpen = ref(true)

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>

<template>
  <q-layout view="hHh lpR fFf">
    <q-header bordered class="text-white" style="background-color: #001529">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="toggleLeftDrawer"/>
        <q-toolbar-title>
          ThingUE Server 控制台
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-drawer v-model="leftDrawerOpen" side="left" elevated behavior="desktop" :width="200">
      <!-- 菜单目录 -->
      <q-list>
        <q-item to="/">
          <q-item-section avatar>
            <q-icon name="home"/>
          </q-item-section>
          <q-item-section>
            首页
          </q-item-section>
        </q-item>
        <q-item to="/client">
          <q-item-section avatar>
            <q-icon name="computer"/>
          </q-item-section>
          <q-item-section>
            客户端
          </q-item-section>
        </q-item>
        <q-item to="/instance">
          <q-item-section avatar>
            <q-icon name="flash_on"/>
          </q-item-section>
          <q-item-section>
            实例
          </q-item-section>
        </q-item>
        <q-item to="/sync">
          <q-item-section avatar>
            <q-icon name="sync"/>
          </q-item-section>
          <q-item-section>
            同步方案
          </q-item-section>
        </q-item>
        <q-expansion-item expand-separator icon="settings" label="设置">
          <q-list class="q-ml-none">
            <q-item to="/relay-setting">
              <q-item-section class="q-ml-lg">
                中继管理
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section class="q-ml-lg">
                密钥
              </q-item-section>
            </q-item>
          </q-list>
        </q-expansion-item>
      </q-list>
    </q-drawer>
    <q-drawer v-model="panelStore.open" side="right" elevated behavior="desktop" :width="panelStore.width" overlay>
      <component v-if="panelStore.open" :is="panelStore.component" :data="panelStore.data"/>
    </q-drawer>
    <q-page-container>
      <RouterView/>
    </q-page-container>
  </q-layout>
</template>

<style lang="sass">
</style>
