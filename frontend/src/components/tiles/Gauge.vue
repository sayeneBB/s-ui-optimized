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
    text: curr[0] + "<sup>" + (curr[1]?? ' ') + "</sup>/" +  total[0] + "<sup>" + (total[1]?? '') + "</sup>"
  }
}

const cssTransformRotateValue = computed(() => {
  const percentageAsFraction = data.value.percent / 100
  const halfPercentage = percentageAsFraction / 2

  return `${halfPercentage}turn`
})

const gaugeColor = computed(() => {
  if (data.value.percent > 90) return 'error'
  if (data.value.percent > 70) return 'warning'
  return 'primary'
})
</script>

<template>
  <div class="gauge__outer">
    <div class="gauge__inner">
      <div
        class="gauge__fill" 
        :style="{ 
          transform: `rotate(${cssTransformRotateValue})`,
          background: `rgb(var(--v-theme-${gaugeColor}))`
          }">
      </div>
      <div class="gauge__cover"><span dir="ltr" v-html="data.text"></span></div>
    </div>
  </div>
</template>

<style scoped>
.gauge__outer {
  width: 100%;
  max-width: 250px;
  margin-top: 15px;
}

.gauge__inner {
  width: 100%;
  height: 0;
  padding-bottom: 50%;
  background: rgba(255, 255, 255, 0.03) !important;
  position: relative;
  border-top-left-radius: 100% 200%;
  border-top-right-radius: 100% 200%;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.gauge__fill {
  position: absolute;
  top: 100%;
  left: 0;
  width: inherit;
  height: 100%;
  transform-origin: center top;
  transform: rotate(0turn);
  transition: transform 0.4s cubic-bezier(0.16, 1, 0.3, 1);
  box-shadow: 0 0 15px rgba(0, 240, 255, 0.3);
}

.gauge__cover {
  width: 78%;
  height: 156%;
  background: rgba(16, 20, 32, 0.95) !important;
  backdrop-filter: blur(8px) !important;
  position: absolute;
  top: 22%;
  left: 50%;
  transform: translateX(-50%);
  border-radius: 50%;
  border: 1px solid rgba(255, 255, 255, 0.05);

  /* Text */
  display: flex;
  align-items: center;
  justify-content: center;
  padding-bottom: 25%;
  box-sizing: border-box;
  font-family: 'Outfit', 'Inter', sans-serif;
  font-weight: 700;
  font-size: 22px;
  color: #ffffff;
}

:deep(sup) {
  font-size: 11px;
  opacity: 0.7;
  vertical-align: super;
}
</style>
