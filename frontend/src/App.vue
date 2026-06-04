<template>
  <v-overlay
    :model-value="loading"
    persistent
    content-class="text-center"
    class="align-center justify-center"
  >
    <v-progress-circular
      indeterminate
      size="64"
    ></v-progress-circular>
    <br />
    {{ $t('loading') }}
  </v-overlay>
  <Message />
  <router-view />
</template>

<script lang="ts" setup>
import Message from '@/components/message.vue'
import { inject, ref, Ref } from 'vue'
import { useTheme } from 'vuetify'

const loading:Ref = inject('loading')?? ref(false)
const theme = useTheme()

const initTheme = () => {
  const savedTheme = localStorage.getItem('theme') ?? 'system'
  if (savedTheme === 'system') {
    const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    theme.global.name.value = isDark ? 'dark' : 'light'
  } else {
    theme.global.name.value = savedTheme
  }
}

// Monitor system theme changes
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
  const savedTheme = localStorage.getItem('theme') ?? 'system'
  if (savedTheme === 'system') {
    theme.global.name.value = e.matches ? 'dark' : 'light'
  }
})

// Run theme initialization immediately
initTheme()

// Change page title
document.title = "S-UI " + document.location.hostname
</script>

<style>
.v-overlay .v-list-item,
.v-field__input {
  direction: ltr;
}
</style>