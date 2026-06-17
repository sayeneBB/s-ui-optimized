<template>
  <div class="login-bg">
    <v-container class="fill-height justify-center align-center">
      <v-row justify="center" align="center" style="width: 100%;">
        <v-col cols="12" sm="8" md="5" lg="4">
          <v-card class="login-card">
            <div class="login-logo-container text-center mb-6">
              <v-img src="@/assets/logo.svg" width="64" class="mx-auto"></v-img>
            </div>
            <div class="login-title mb-6" v-text="$t('login.title')"></div>
            <v-card-text class="pa-0">
              <v-form @submit.prevent="login" ref="form">
                <v-text-field 
                  v-model="username" 
                  :label="$t('login.username')" 
                  :rules="usernameRules" 
                  required
                  prepend-inner-icon="mdi-account"
                  class="mb-4"
                  hide-details="auto"
                ></v-text-field>
                <v-text-field 
                  v-model="password" 
                  :label="$t('login.password')" 
                  :rules="passwordRules" 
                  type="password" 
                  required
                  prepend-inner-icon="mdi-lock"
                  class="mb-6"
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
                  variant="outlined"
                  :items="languages"
                  v-model="$i18n.locale"
                  @update:modelValue="changeLocale"
                  class="locale-select"
                  style="max-width: 140px;"
                ></v-select>
                
                <v-menu>
                  <template v-slot:activator="{ props }">
                    <v-btn icon v-bind="props" variant="outlined" color="secondary" class="theme-toggle-btn">
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
.login-bg {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgb(var(--v-theme-background));
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.login-card {
  background: rgb(var(--v-theme-surface)) !important;
  border: 1px solid rgb(var(--v-theme-border)) !important;
  border-radius: 8px !important;
  padding: 32px 24px;
  box-shadow: none !important;
}

.login-title {
  font-family: 'Inter', system-ui, sans-serif;
  font-weight: 700 !important;
  color: rgb(var(--v-theme-primary));
  text-align: center;
  font-size: 1.6rem !important;
  letter-spacing: -0.02em;
}

.login-btn {
  border-radius: 6px !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  height: 44px !important;
  text-transform: none !important;
}

.theme-toggle-btn {
  border-color: rgb(var(--v-theme-border)) !important;
}

.theme-list {
  background: rgb(var(--v-theme-surface)) !important;
  border: 1px solid rgb(var(--v-theme-border)) !important;
  border-radius: 8px !important;
}
</style>
