<template>
  <v-navigation-drawer
    v-model="showDrawer"
    :temporary="isMobile"
    :expand-on-hover="!isMobile"
    :rail="!isMobile"
    :permanent="!isMobile"
    class="glass-drawer"
    @click="isMobile ? $emit('toggleDrawer') : null"
  >
    <v-list-item
      height="63"
      prepend-avatar="@/assets/logo.svg"
      title="S-UI"
      class="logo-item"
    >
      <template v-slot:prepend>
        <v-img src="@/assets/logo.svg" width="30" class="logo-glow mr-3" />
      </template>
      <template v-slot:append v-if="isMobile">
        <v-icon icon="mdi-close" />
      </template>
    </v-list-item>

    <v-list density="compact" nav class="py-4">
      <v-list-item link
        v-for="item in menu"
        :key="item.title"
        :to="item.path"
        class="nav-item"
        active-class="nav-item-active"
        :active="router.currentRoute.value.path == item.path">
        <template v-slot:prepend>
          <v-icon :icon="item.icon" size="20"></v-icon>
        </template>
        <v-list-item-title class="nav-title" v-text="$t(item.title)"></v-list-item-title>
      </v-list-item>
    </v-list>
    
    <template v-slot:append>
      <v-list-item 
        prepend-icon="mdi-logout" 
        :title="$t('menu.logout')" 
        class="logout-item"
        @click="Logout"
      ></v-list-item>
    </template>
  </v-navigation-drawer>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import router from '@/router'
import { logout } from '@/plugins/httputil'

const props = defineProps(['isMobile','displayDrawer'])

const showDrawer = computed((): boolean => {
  return props.displayDrawer
})

const menu = [
  { title: 'pages.home', icon: 'mdi-home',  path: '/' },
  { title: 'pages.inbounds', icon: 'mdi-cloud-download',  path: '/inbounds' },
  { title: 'pages.clients', icon: 'mdi-account-multiple',  path: '/clients' },
  { title: 'pages.outbounds', icon: 'mdi-cloud-upload',  path: '/outbounds' },
  { title: 'pages.endpoints', icon: 'mdi-cloud-tags',  path: '/endpoints' },
  { title: 'pages.services', icon: 'mdi-server',  path: '/services' },
  { title: 'pages.tls', icon: 'mdi-certificate',  path: '/tls' },
  { title: 'pages.basics', icon: 'mdi-application-cog',  path: '/basics' },
  { title: 'pages.rules', icon: 'mdi-routes',  path: '/rules' },
  { title: 'pages.dns', icon: 'mdi-dns',  path: '/dns' },
  { title: 'pages.admins', icon: 'mdi-account-tie',  path: '/admins' },
  { title: 'pages.settings', icon: 'mdi-cog',  path: '/settings' },
]

const Logout = async () => {
  logout()
}
</script>

<style scoped>
.glass-drawer {
  background: rgb(var(--v-theme-surface)) !important;
  border-right: 1px solid rgb(var(--v-theme-border)) !important;
}

.logo-item {
  border-bottom: 1px solid rgb(var(--v-theme-border)) !important;
  background: transparent !important;
}

.logo-glow {
  filter: none !important;
}

.nav-item {
  margin: 4px 8px !important;
  border-radius: 6px !important;
  transition: all 0.15s ease !important;
  color: rgb(var(--v-theme-secondary)) !important;
}

.nav-item:hover {
  background: rgba(var(--v-theme-primary), 0.04) !important;
  color: rgb(var(--v-theme-primary)) !important;
}

.nav-item-active {
  background: rgba(var(--v-theme-primary), 0.06) !important;
  color: rgb(var(--v-theme-primary)) !important;
  font-weight: 600 !important;
}

.nav-item-active :deep(.v-icon) {
  color: rgb(var(--v-theme-primary)) !important;
}

.nav-title {
  font-family: 'Inter', system-ui, sans-serif;
  font-size: 0.85rem !important;
  font-weight: 500 !important;
}

.logout-item {
  margin: 8px !important;
  border-radius: 6px !important;
  color: rgb(var(--v-theme-error)) !important;
  transition: all 0.15s ease !important;
}

.logout-item:hover {
  background: rgba(var(--v-theme-error), 0.08) !important;
  color: rgb(var(--v-theme-error)) !important;
}
</style>