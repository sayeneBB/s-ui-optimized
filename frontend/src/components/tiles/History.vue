<template>
  <Line v-if="loaded" :data="chartData" :options="<any>chartOptions" />
</template>

<script lang="ts" setup>
import { ref, computed, watch } from 'vue'
import { Line } from 'vue-chartjs'
import { useTheme } from 'vuetify'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Filler,
} from 'chart.js'
import { HumanReadable } from '@/plugins/utils'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Filler
)

const props = defineProps(['tilesData', 'type'])

const theme = useTheme()
const loaded = ref(false)
const labels = new Array(20).fill('')
const oldValues = ref<any>({ net: {}, dio: {} })
const rawData = ref<any>({ datasets: [] })

// Define colors dynamically based on active theme
const themeColors = computed(() => {
  const isDark = theme.global.current.value.dark
  return {
    text: isDark ? 'rgba(255, 255, 255, 0.5)' : 'rgba(0, 0, 0, 0.5)',
    grid: isDark ? 'rgba(255, 255, 255, 0.06)' : 'rgba(0, 0, 0, 0.06)',
    accentPrimary: isDark ? '#3B82F6' : '#2563EB', // Blue
    accentSecondary: isDark ? '#10B981' : '#059669', // Green
    bgPrimary: isDark ? 'rgba(59, 130, 246, 0.08)' : 'rgba(37, 99, 235, 0.04)',
    bgSecondary: isDark ? 'rgba(16, 185, 129, 0.08)' : 'rgba(5, 150, 105, 0.04)'
  }
})

const baseOptions = computed(() => {
  const colors = themeColors.value
  return {
    animation: { duration: 250 },
    responsive: true,
    maintainAspectRatio: false,
    interaction: {
      intersect: false,
      mode: 'index',
    },
    plugins: {
      tooltip: {
        enabled: true,
        backgroundColor: theme.global.current.value.dark ? '#18181B' : '#FFFFFF',
        titleColor: theme.global.current.value.dark ? '#FAFAFA' : '#09090B',
        bodyColor: theme.global.current.value.dark ? '#A1A1AA' : '#71717A',
        titleFont: { family: 'Inter, system-ui, sans-serif', weight: 'bold' },
        bodyFont: { family: 'Inter, system-ui, sans-serif' },
        borderColor: colors.grid,
        borderWidth: 1,
        padding: 8,
        cornerRadius: 6,
      },
      legend: { display: false }
    },
    scales: {
      x: {
        grid: { display: false },
        ticks: {
          display: false,
          font: { family: 'Inter, system-ui, sans-serif' }
        }
      },
      y: {
        grid: { color: colors.grid },
        border: { dash: [4, 4] },
        beginAtZero: true,
        ticks: {
          color: colors.text,
          font: { size: 9, family: 'JetBrains Mono, SF Mono, monospace' }
        }
      }
    }
  }
})

const chartOptions = computed(() => {
  const opts = JSON.parse(JSON.stringify(baseOptions.value))
  
  // Assign tooltip labels formatting manually since functions get lost during deep copy
  opts.plugins.tooltip.callbacks = {
    label: (context: any) => {
      let label = context.dataset.label || ''
      if (label) label += ': '
      if (context.parsed.y !== null) {
        if (props.type === 'h-net' || props.type === 'h-dio') {
          label += (context.parsed.y === 0 ? '0 B' : HumanReadable.sizeFormat(context.parsed.y, 1)) + '/s'
        } else if (props.type === 'hp-net') {
          label += (context.parsed.y === 0 ? '0 p' : HumanReadable.packetFormat(context.parsed.y, 1)) + '/s'
        } else {
          label += context.parsed.y.toFixed(1) + '%'
        }
      }
      return label
    }
  }

  switch (props.type) {
    case 'h-net':
      opts.scales.y.ticks.callback = (label: any) => {
        return label == 0 ? "0" : HumanReadable.sizeFormat(label, 0) + '/s'
      }
      break
    case 'hp-net':
      opts.scales.y.ticks.callback = (label: any) => {
        return label == 0 ? "0" : HumanReadable.packetFormat(label, 0) + '/s'
      }
      break
    case 'h-dio':
      opts.scales.y.ticks.callback = (label: any) => {
        return label == 0 ? "0" : HumanReadable.sizeFormat(label, 0) + '/s'
      }
      break
    default:
      opts.scales.y.min = 0
      opts.scales.y.max = 100
      opts.scales.y.ticks.callback = (label: any) => label + '%'
      break
  }
  return opts
})

const chartData = computed(() => {
  const colors = themeColors.value
  if (!rawData.value.datasets || rawData.value.datasets.length === 0) {
    return { labels, datasets: [] }
  }
  
  // Map theme-adaptive styles to existing datasets
  return {
    labels,
    datasets: rawData.value.datasets.map((dataset: any, idx: number) => {
      const isSecondary = idx === 0 && rawData.value.datasets.length > 1
      return {
        ...dataset,
        backgroundColor: isSecondary ? colors.bgSecondary : colors.bgPrimary,
        borderColor: isSecondary ? colors.accentSecondary : colors.accentPrimary,
        pointBackgroundColor: isSecondary ? colors.accentSecondary : colors.accentPrimary,
        borderWidth: 1.5,
        pointRadius: 0,
        pointHoverRadius: 4,
        tension: 0.3,
        fill: true
      }
    })
  }
})

const updateData1 = (value1: number) => {
  const newData = [...(rawData.value.datasets?.[0]?.data || []), value1]
  if (newData.length > 20) newData.shift()
  rawData.value = {
    datasets: [{ label: '', data: newData }]
  }
  loaded.value = true
}

const updateData2 = (value1: number, value2: number) => {
  const newData1 = [...(rawData.value.datasets?.[0]?.data || []), value1]
  const newData2 = [...(rawData.value.datasets?.[1]?.data || []), value2]
  if (newData1.length > 20) newData1.shift()
  if (newData2.length > 20) newData2.shift()
  rawData.value = {
    datasets: [
      { label: 'Upload', data: newData1 },
      { label: 'Download', data: newData2 }
    ]
  }
  loaded.value = true
}

watch(() => props.tilesData, (v: any) => {
  if (!v) return
  switch (props.type) {
    case 'h-cpu':
      if (v.cpu !== undefined) updateData1(v.cpu)
      break
    case 'h-mem':
      if (v.mem && v.mem.current !== undefined && v.mem.total) {
        updateData1(v.mem.current * 100 / v.mem.total)
      }
      break
    case 'h-net':
      if (v.net && oldValues.value.net && oldValues.value.net.sent !== undefined) {
        const downSpeed = (v.net.recv - oldValues.value.net.recv) / 2
        const upSpeed = (v.net.sent - oldValues.value.net.sent) / 2
        updateData2(upSpeed, downSpeed)
      }
      if (v.net) oldValues.value.net = v.net
      break
    case 'hp-net':
      if (v.net && oldValues.value.net && oldValues.value.net.psent !== undefined) {
        const downSpeed = (v.net.precv - oldValues.value.net.precv) / 2
        const upSpeed = (v.net.psent - oldValues.value.net.psent) / 2
        updateData2(upSpeed, downSpeed)
      }
      if (v.net) oldValues.value.net = v.net
      break
    case 'h-dio':
      if (v.dio && oldValues.value.dio && oldValues.value.dio.read !== undefined) {
        const downSpeed = (v.dio.read - oldValues.value.dio.read) / 2
        const upSpeed = (v.dio.write - oldValues.value.dio.write) / 2
        updateData2(upSpeed, downSpeed)
      }
      if (v.dio) oldValues.value.dio = v.dio
      break
  }
}, { deep: true })
</script>