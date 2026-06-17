import pluginVue from 'eslint-plugin-vue'
import vueTsConfigs from '@vue/eslint-config-typescript'

export default [
  {
    ignores: ['dist/**'],
  },
  ...pluginVue.configs['flat/essential'],
  ...vueTsConfigs(),
  {
    rules: {
      'vue/multi-word-component-names': 'off',
      '@typescript-eslint/no-explicit-any': 'off',
      '@typescript-eslint/no-unused-vars': 'off',
      '@typescript-eslint/no-empty-object-type': 'off',
      '@typescript-eslint/no-unused-expressions': 'off',
      'vue/no-v-text-v-html-on-component': 'off',
      'vue/require-v-for-key': 'off',
      'no-unused-vars': 'off',
      'no-undef': 'off',
      'no-empty': 'off',
      'no-constant-condition': 'off',
      'no-control-regex': 'off',
      'vue/no-mutating-props': 'off',
      'vue/valid-v-for': 'off',
      'vue/no-side-effects-in-computed-properties': 'off',
      'vue/no-reserved-component-names': 'off',
      'vue/valid-v-slot': 'off',
      'vue/no-use-v-if-with-v-for': 'off',
      'vue/no-unused-components': 'off',
      'vue/no-unused-vars': 'off',
      'prefer-const': 'off',
    }
  }
]
