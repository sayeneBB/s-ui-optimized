<template>
  <div class="login-bg">
    <div class="glow-orb glow-orb-1"></div>
    <div class="glow-orb glow-orb-2"></div>
    <v-container class="fill-height justify-center align-center">
      <v-row justify="center" align="center" style="width: 100%;">
        <v-col cols="12" sm="8" md="5" lg="4">
          <v-card class="login-card" elevation="0">
            <div class="login-logo-container text-center mb-6">
              <v-img src="@/assets/logo.svg" width="80" class="mx-auto logo-glow"></v-img>
            </div>
            <div class="login-title mb-6" v-text="$t('login.title')"></div>
            <v-card-text>
              <v-form @submit.prevent="login" ref="form">
                <v-text-field 
                  v-model="username" 
                  :label="$t('login.username')" 
                  :rules="usernameRules" 
                  required
                  prepend-inner-icon="mdi-account"
                  class="custom-field mb-4"
                  hide-details="auto"
                ></v-text-field>
                <v-text-field 
                  v-model="password" 
                  :label="$t('login.password')" 
                  :rules="passwordRules" 
                  type="password" 
                  required
                  prepend-inner-icon="mdi-lock"
                  class="custom-field mb-6"
                  hide-details="auto"
                ></v-text-field>
                <v-btn 
                  :loading="loading" 
                  type="submit" 
                  color="primary" 
                  block 
                  class="login-btn py-6 d-flex align-center justify-center" 
                  v-text="$t('actions.submit')"
                ></v-btn>
              </v-form>
              
              <div class="d-flex align-center justify-space-between mt-8">
                <v-select
                  density="compact"
                  hide-details
                  variant="solo-filled"
                  :items="languages"
                  v-model="$i18n.locale"
                  @update:modelValue="changeLocale"
                  class="locale-select"
                  style="max-width: 140px;"
                ></v-select>
                
                <v-menu>
                  <template v-slot:activator="{ props }">
                    <v-btn icon v-bind="props" variant="tonal" color="primary" class="theme-toggle-btn">
                      <v-icon>mdi-theme-light-dark</v-icon>
                    </v-btn>
                  </template>
                  <v-list class="theme-list">
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
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useLocale, useTheme } from 'vuetify'
import { i18n, languages } from '@/locales'
import { useRouter } from 'vue-router'
import HttpUtil from '@/plugins/httputil'

const theme = useTheme()
const locale = useLocale()

const themes = [
  { value: 'light', icon: 'mdi-white-balance-sunny' },
  { value: 'dark', icon: 'mdi-moon-waning-crescent' },
  { value: 'system', icon: 'mdi-laptop' },
]

const username = ref('')
const usernameRules = [
  (value: string) => {
    if (value?.length > 0) return true
    return i18n.global.t('login.unRules')
  },
]

const password = ref('')
const passwordRules = [
  (value: string) => {
    if (value?.length > 0) return true
    return i18n.global.t('login.pwRules')
  },
]

const loading = ref(false)
const router = useRouter()

const login = async () => {
  if (username.value == '' || password.value == '') return
  loading.value = true
  const response = await HttpUtil.post('api/login', { user: username.value, pass: password.value })
  if (response.success) {
    setTimeout(() => {
      loading.value = false
      router.push('/')
    }, 500)
  } else {
    loading.value = false
  }
}
const changeLocale = (l: any) => {
  locale.current.value = l ?? 'zhHans'
  localStorage.setItem('locale', locale.current.value)
}
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
.login-bg {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: var(--body-bg);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}



.login-card {
  background: var(--glass-bg) !important;
  backdrop-filter: blur(30px) saturate(180%) !important;
  -webkit-backdrop-filter: blur(30px) saturate(180%) !important;
  border: 1px solid var(--glass-border) !important;
  box-shadow: var(--glass-shadow) !important;
  border-radius: 28px !important;
  padding: 36px 28px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1) !important;
}

.login-card:hover {
  border-color: var(--glass-border-hover) !important;
  box-shadow: var(--glass-shadow-hover) !important;
  transform: translateY(-4px) !important;
}

.logo-glow {
  filter: var(--logo-glow-filter);
}

.login-title {
  font-family: 'Outfit', 'Inter', sans-serif;
  font-weight: 700 !important;
  background: linear-gradient(135deg, var(--text-primary) 40%, var(--accent-color) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  text-align: center;
  font-size: 1.9rem !important;
  letter-spacing: -0.03em;
}

.custom-field :deep(.v-field) {
  border-radius: 12px !important;
  border: 1px solid var(--glass-border) !important;
  background: var(--input-bg) !important;
  color: var(--text-primary) !important;
  transition: all 0.3s ease !important;
}

.custom-field :deep(.v-field--focused) {
  border-color: var(--accent-color) !important;
  box-shadow: var(--accent-glow) !important;
}

.login-btn {
  border-radius: 12px !important;
  font-weight: 600 !important;
  font-size: 1.05rem !important;
  background: linear-gradient(135deg, var(--accent-color) 0%, #00b0ff 100%) !important;
  border: none !important;
  color: #ffffff !important;
  box-shadow: var(--accent-glow) !important;
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1) !important;
}

.login-btn:hover {
  box-shadow: 0 8px 25px rgba(0, 240, 255, 0.35) !important;
  transform: translateY(-2px) !important;
}

.v-theme--light .login-btn:hover {
  box-shadow: 0 8px 25px rgba(98, 0, 234, 0.35) !important;
}

.theme-toggle-btn {
  background: rgba(0, 240, 255, 0.08) !important;
  border: 1px solid rgba(0, 240, 255, 0.15) !important;
  color: var(--accent-color) !important;
}

.v-theme--light .theme-toggle-btn {
  background: rgba(98, 0, 234, 0.08) !important;
  border: 1px solid rgba(98, 0, 234, 0.15) !important;
}

.theme-list {
  background: var(--dialog-bg) !important;
  backdrop-filter: blur(10px) !important;
  border: 1px solid var(--glass-border) !important;
  border-radius: 12px !important;
  color: var(--text-primary) !important;
}
</style>
