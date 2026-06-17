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
  background: radial-gradient(circle at 50% 50%, #0d111a 0%, #080b11 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.glow-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(120px);
  opacity: 0.15;
  mix-blend-mode: screen;
  pointer-events: none;
  animation: orbFloat 20s infinite alternate ease-in-out;
}

.glow-orb-1 {
  width: 450px;
  height: 450px;
  background: #00f0ff;
  top: -10%;
  left: 10%;
}

.glow-orb-2 {
  width: 550px;
  height: 550px;
  background: #bf5af2;
  bottom: -15%;
  right: 5%;
  animation-delay: -7s;
}

@keyframes orbFloat {
  0% { transform: translate(0, 0) scale(1); }
  50% { transform: translate(60px, -80px) scale(1.15); }
  100% { transform: translate(-40px, 40px) scale(0.9); }
}

.login-card {
  background: rgba(16, 20, 32, 0.4) !important;
  backdrop-filter: blur(30px) saturate(180%) !important;
  -webkit-backdrop-filter: blur(30px) saturate(180%) !important;
  border: 1px solid rgba(255, 255, 255, 0.08) !important;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.35), 0 0 25px rgba(0, 240, 255, 0.05) !important;
  border-radius: 28px !important;
  padding: 36px 28px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1) !important;
}

.login-card:hover {
  border-color: rgba(0, 240, 255, 0.35) !important;
  box-shadow: 0 20px 45px rgba(0, 0, 0, 0.45), 0 0 35px rgba(0, 240, 255, 0.18) !important;
  transform: translateY(-4px) !important;
}

.logo-glow {
  filter: drop-shadow(0 0 12px rgba(0, 240, 255, 0.55));
}

.login-title {
  font-family: 'Outfit', 'Inter', sans-serif;
  font-weight: 700 !important;
  background: linear-gradient(135deg, #ffffff 40%, #00F0FF 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  text-align: center;
  font-size: 1.9rem !important;
  letter-spacing: -0.03em;
}

.custom-field :deep(.v-field) {
  border-radius: 12px !important;
  border: 1px solid rgba(255, 255, 255, 0.08) !important;
  background: rgba(10, 14, 22, 0.55) !important;
  transition: all 0.3s ease !important;
}

.custom-field :deep(.v-field--focused) {
  border-color: rgba(0, 240, 255, 0.5) !important;
  box-shadow: 0 0 12px rgba(0, 240, 255, 0.15) !important;
}

.login-btn {
  border-radius: 12px !important;
  font-weight: 600 !important;
  font-size: 1.05rem !important;
  background: linear-gradient(135deg, #00f0ff 0%, #0072ff 100%) !important;
  border: none !important;
  color: #ffffff !important;
  box-shadow: 0 6px 20px rgba(0, 240, 255, 0.25) !important;
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1) !important;
}

.login-btn:hover {
  box-shadow: 0 8px 25px rgba(0, 240, 255, 0.45) !important;
  transform: translateY(-2px) !important;
}

.theme-toggle-btn {
  background: rgba(0, 240, 255, 0.08) !important;
  border: 1px solid rgba(0, 240, 255, 0.15) !important;
}

.theme-list {
  background: rgba(16, 20, 32, 0.9) !important;
  backdrop-filter: blur(10px) !important;
  border: 1px solid rgba(255, 255, 255, 0.08) !important;
  border-radius: 12px !important;
}
</style>
