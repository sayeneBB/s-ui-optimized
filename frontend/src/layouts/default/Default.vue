<template>
  <v-app style="overflow: auto;">
    <div class="glow-orb glow-orb-1"></div>
    <div class="glow-orb glow-orb-2"></div>
    <drawer :isMobile="isMobile" :displayDrawer="displayDrawer" @toggleDrawer="toggleDrawer" />
    <default-bar :isMobile="isMobile" @toggleDrawer="toggleDrawer" />
    <default-view />
  </v-app>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import DefaultBar from './AppBar.vue'
import Drawer from './Drawer.vue'
import DefaultView from './View.vue'
import { useDisplay } from 'vuetify'

const { smAndDown } = useDisplay()
const displayDrawer = ref(false)

const toggleDrawer = () => {
  displayDrawer.value = !displayDrawer.value
}

const isMobile = computed( ():boolean =>{
  displayDrawer.value = !smAndDown.value
  return smAndDown.value
})
</script>

<style>
.v-card-subtitle {
  text-align: center;
  border-bottom: 1px solid rgb(var(--v-theme-border)) !important;
  min-height: 20px;
  font-size: 0.8rem;
  color: rgb(var(--v-theme-secondary));
}
.v-switch.v-input {
  padding-inline-start: .6rem;
}
</style>