import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles/main.css'

import { fa, en, vi, zhHans, zhHant, ru } from 'vuetify/locale'
import { createVuetify } from 'vuetify'

const getDefaultTheme = () => {
  if (typeof window === 'undefined') return 'dark'
  const t = localStorage.getItem('theme') ?? 'dark'
  if (t === 'system') {
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }
  return t
}

export default createVuetify({
  defaults: {
    VRow: { density: 'compact' },
    VTextField: {
      variant: 'outlined',
      density: 'comfortable',
    },
    VSelect: {
      variant: 'outlined',
      density: 'comfortable',
    },
    VCombobox: {
      variant: 'outlined',
      density: 'comfortable',
    },
    VTextarea: {
      variant: 'outlined',
      density: 'comfortable',
    },
    VBtn: {
      rounded: 'lg',
      elevation: 0,
    },
    VCard: {
      rounded: 'lg',
      elevation: 0,
    }
  },
  theme: {
    defaultTheme: getDefaultTheme(),
    themes: {
      light: {
        dark: false,
        colors: {
          primary: '#18181B', // Slate-900
          secondary: '#71717A',
          background: '#FAFAFA', // Slate-50
          surface: '#FFFFFF',
          error: '#DC2626',
          warning: '#D97706',
          success: '#16A34A',
          info: '#2563EB',
          border: '#E4E4E7', // Slate-200
        },
      },
      dark: {
        dark: true,
        colors: {
          primary: '#FAFAFA', // Slate-50
          secondary: '#A1A1AA',
          background: '#09090B', // Slate-950
          surface: '#18181B', // Slate-900
          error: '#EF4444',
          warning: '#F59E0B',
          success: '#22C55E',
          info: '#3B82F6',
          border: '#27272A', // Slate-800
        },
      },
    },
  },
  locale: {
    locale: localStorage.getItem("locale") ?? 'zhHans',
    fallback: 'zhHans',
    messages: { en, fa, vi, zhHans, zhHant, ru },
  },
})
