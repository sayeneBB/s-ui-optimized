<script lang="ts" setup>
import { HumanReadable } from '@/plugins/utils'
import { computed } from 'vue'

const props = defineProps({
  tilesData: <any>{},
  type: String
})

const data = computed(() => {
  const d = props.tilesData
  if (!d.mem && !d.cpu) return { percent: 0, text: '-' }
  switch (props.type) {
    case 'g-cpu':
      return { percent: d.cpu, text: Math.ceil(d.cpu) + "%" }
    case 'g-mem':
      return gaugeData(d.mem)
    case 'g-dsk':
      return gaugeData(d.dsk)
    case 'g-swp':
      return gaugeData(d.swp)
  }
  return { percent: 0, text: '-'}
})

const gaugeData = (d:any) :any => {
  if (!d) return { percent: 0, text: '-' }
  const curr = HumanReadable.sizeFormat(d.current,0).split(' ')
  const total = HumanReadable.sizeFormat(d.total,0).split(' ')
  if (curr[1] == total[1]) curr[1] = ''
  return {
    percent: Math.ceil(d.current*100/d.total),
    text: curr[0] + (curr[1]?? ' ') + " / " +  total[0] + (total[1]?? '')
  }
}

const gaugeColor = computed(() => {
  if (data.value.percent > 90) return 'error'
  if (data.value.percent > 70) return 'warning'
  return 'primary'
})
</script>

<template>
  <div class="linear-gauge py-4 px-2 w-100">
    <div class="d-flex justify-space-between align-center mb-3">
      <span class="text-caption text-secondary uppercase-label">Usage Status</span>
      <span class="text-body-1 font-weight-bold mono-text text-primary" v-html="data.text"></span>
    </div>
    <v-progress-linear
      :model-value="data.percent"
      :color="gaugeColor"
      height="8"
      rounded
      class="bg-track-border"
    ></v-progress-linear>
    <div class="d-flex justify-space-between mt-2 text-caption text-secondary mono-text">
      <span>{{ data.percent }}%</span>
      <span>100%</span>
    </div>
  </div>
</template>

<style scoped>
.linear-gauge {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
  font-family: 'Inter', system-ui, sans-serif;
}

.uppercase-label {
  text-transform: uppercase;
  font-size: 0.75rem !important;
  font-weight: 600;
  letter-spacing: 0.05em;
}

.mono-text {
  font-family: 'JetBrains Mono', 'SF Mono', Consolas, monospace !important;
  font-size: 0.9rem !important;
}

.bg-track-border {
  background-color: rgba(var(--v-theme-secondary), 0.12) !important;
}
</style>
