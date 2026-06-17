<template>
  <v-app-bar :elevation="0" class="glass-app-bar">
    <v-btn icon v-if="isMobile" @click="$emit('toggleDrawer')" class="ml-2">
      <v-icon>mdi-menu</v-icon>
    </v-btn>
    <span v-else style="width: 24px"></span>
    
    <v-app-bar-title class="app-bar-title text-center">
      {{ $t(<string>route.name) }}
    </v-app-bar-title>
    
    <div class="d-flex align-center mr-4">
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn icon v-bind="props" class="action-btn mx-1">
            <v-icon>mdi-translate</v-icon>
          </v-btn>
        </template>
        <v-list class="menu-list">
          <v-list-item
            v-for="lang in languages"
            :key="lang.value"
            @click="changeLocale(lang.value)"
            :active="isActiveLocale(lang.value)"
          >
            <v-list-item-title>{{ lang.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
      
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn icon v-bind="props" class="action-btn mx-1">
            <v-icon>mdi-theme-light-dark</v-icon>
          </v-btn>
        </template>
        <v-list class="menu-list">
          <v-list-item
            v-for="th in themes"
            :key="th.value"
            @click="changeTheme(th.value)"
            :prepend-icon="th.icon"
            :active="isActiveTheme(th.value)"
          >
            <v-list-item-title>{{ $t(`theme.${th.value}`) }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </div>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { useLocale, useTheme } from 'vuetify'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { languages } from '@/locales'

defineProps(['isMobile'])

const route = useRoute()
const { locale: i18nLocale } = useI18n()
const vuetifyLocale = useLocale()
const theme = useTheme()

const changeLocale = (l: string) => {
  i18nLocale.value = l
  vuetifyLocale.current.value = l
  localStorage.setItem('locale', l)
  window.location.reload()
}
const isActiveLocale = (l: string) => i18nLocale.value === l
const themes = [
  { value: 'light', icon: 'mdi-white-balance-sunny' },
  { value: 'dark', icon: 'mdi-moon-waning-crescent' },
  { value: 'system', icon: 'mdi-laptop' },
]

const changeTheme = (th: string) => {
  localStorage.setItem('theme', th)
  if (th === 'system') {
    theme.global.name.value = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  } else {
    theme.global.name.value = th
  }
}
const isActiveTheme = (th: string) => {
  const current = localStorage.getItem('theme') ?? 'system'
  return current == th
}
</script>

<style scoped>
.glass-app-bar {
  background: rgb(var(--v-theme-surface)) !important;
  border-bottom: 1px solid rgb(var(--v-theme-border)) !important;
}

.app-bar-title {
  font-family: 'Inter', system-ui, sans-serif;
  font-weight: 600 !important;
  font-size: 1.05rem !important;
  letter-spacing: -0.01em;
}

.action-btn {
  background: transparent !important;
  border: 1px solid rgb(var(--v-theme-border)) !important;
  color: rgb(var(--v-theme-secondary)) !important;
  transition: all 0.2s ease !important;
}

.action-btn:hover {
  background: rgb(var(--v-theme-border)) !important;
  color: rgb(var(--v-theme-primary)) !important;
}

.menu-list {
  background: rgb(var(--v-theme-surface)) !important;
  border: 1px solid rgb(var(--v-theme-border)) !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1) !important;
}
</style>
