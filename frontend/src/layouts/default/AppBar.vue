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
    const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    theme.global.name.value = isDark ? 'dark' : 'light'
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
  background: var(--glass-appbar-bg) !important;
  backdrop-filter: blur(20px) saturate(180%) !important;
  -webkit-backdrop-filter: blur(20px) saturate(180%) !important;
  border-bottom: 1px solid var(--glass-border) !important;
}

.app-bar-title {
  font-family: 'Outfit', 'Inter', sans-serif;
  font-weight: 600 !important;
  font-size: 1.15rem !important;
  letter-spacing: -0.01em;
  color: var(--text-primary) !important;
}

.action-btn {
  background: rgba(255, 255, 255, 0.03) !important;
  border: 1px solid var(--glass-border) !important;
  transition: all 0.3s ease !important;
  color: var(--text-primary) !important;
}

.v-theme--light .action-btn {
  background: rgba(0, 0, 0, 0.03) !important;
}

.action-btn:hover {
  background: rgba(0, 240, 255, 0.08) !important;
  border-color: var(--accent-color) !important;
  color: var(--accent-color) !important;
}

.v-theme--light .action-btn:hover {
  background: rgba(98, 0, 234, 0.08) !important;
}

.menu-list {
  background: var(--dialog-bg) !important;
  backdrop-filter: blur(10px) !important;
  border: 1px solid var(--glass-border) !important;
  border-radius: 12px !important;
  color: var(--text-primary) !important;
}
</style>
