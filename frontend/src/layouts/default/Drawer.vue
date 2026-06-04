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
  background: var(--glass-drawer-bg) !important;
  backdrop-filter: blur(20px) saturate(180%) !important;
  -webkit-backdrop-filter: blur(20px) saturate(180%) !important;
  border-right: 1px solid var(--glass-border) !important;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1) !important;
}

.logo-item {
  border-bottom: 1px solid var(--glass-border) !important;
  background: rgba(0, 0, 0, 0.05) !important;
}

.v-theme--dark .logo-item {
  background: rgba(0, 0, 0, 0.15) !important;
}

.logo-glow {
  filter: var(--logo-glow-filter);
}

.nav-item {
  margin: 4px 10px !important;
  border-radius: 12px !important;
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1) !important;
  color: var(--text-secondary) !important;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.04) !important;
  color: var(--accent-color) !important;
  transform: translateX(4px);
}

.v-theme--light .nav-item:hover {
  background: rgba(0, 0, 0, 0.03) !important;
}

.nav-item-active {
  background: var(--nav-active-bg) !important;
  border-left: 3px solid var(--accent-color) !important;
  color: var(--accent-color) !important;
  box-shadow: var(--accent-glow) !important;
}

.nav-item-active :deep(.v-icon) {
  color: var(--accent-color) !important;
  filter: drop-shadow(0 0 4px var(--accent-color));
}

.nav-title {
  font-family: 'Outfit', 'Inter', sans-serif;
  font-size: 0.9rem !important;
  font-weight: 500 !important;
}

.logout-item {
  margin: 8px 10px !important;
  border-radius: 12px !important;
  color: rgba(255, 82, 82, 0.75) !important;
  transition: all 0.3s ease !important;
}

.logout-item:hover {
  background: rgba(255, 82, 82, 0.08) !important;
  color: #FF5252 !important;
  transform: scale(0.98);
}
</style>