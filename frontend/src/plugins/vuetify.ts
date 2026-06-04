/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles/main.css'

import colors from 'vuetify/util/colors'
import { fa, en, vi, zhHans, zhHant, ru } from 'vuetify/locale'

// Composables
import { createVuetify } from 'vuetify'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  defaults: {
    VRow: { density: 'compact' },
    VTextField: {
      variant: 'solo-filled',
    },
    VSelect: {
      variant: 'solo-filled',
    },
    VCombobox: {
      variant: 'solo-filled',
    },
    VTextarea: {
      variant: 'solo-filled',
    },
  },
  theme: {
    defaultTheme: 'dark',
    themes: {
      light: {
        colors: {
          primary: '#00B0FF',
          background: '#F4F6F9',
          surface: '#FFFFFF',
          error: '#FF5252',
        },
      },
      dark: {
        dark: true,
        colors: {
          primary: '#00F0FF',     // 激光霓虹蓝 (Neon Cyan)
          secondary: '#bf5af2',   // 激光霓虹紫 (Laser Purple)
          background: '#080b11',  // 深空底色 (Deep Slate Black)
          surface: '#101420',     // 玻璃卡片底板 (Translucent Glass Slate)
          success: '#00E676',     // 呼吸状态绿
          warning: '#FF9100',     // 呼吸警示橙
          error: '#FF5252',       // 报错珊瑚红
          info: '#00B0FF',        // 天空蔚蓝
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
