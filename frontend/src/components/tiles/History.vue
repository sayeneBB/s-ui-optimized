<template>
  <Line v-if="loaded" :data="data" :options="<any>options" />
</template>

<script lang="ts">
import { ref } from 'vue'
import { Line } from 'vue-chartjs'
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
ChartJS.defaults.font.family = 'Outfit, Inter, sans-serif'
ChartJS.defaults.color = 'rgba(255, 255, 255, 0.65)'

export default {
  components: {
    Line
  },
  props: ['tilesData','type'],
  data() {
    return {
      loaded: false,
      labels: new Array(20).fill(''),
      oldValues: <any>{net: {}, dio: {}},
      options1: {
        animation: {
          duration: 300
        },
        responsive: true,
        maintainAspectRatio: false,
        interaction: {
          intersect: false,
          mode: 'index',
        },
        plugins: {
          tooltip: {
            enabled: true,
            backgroundColor: 'rgba(13, 17, 26, 0.85)',
            titleColor: '#00F0FF',
            bodyColor: '#ffffff',
            borderColor: 'rgba(255, 255, 255, 0.1)',
            borderWidth: 1,
            padding: 8,
            cornerRadius: 8,
          },
          legend: {
              display: false,
          }
        },
        scales: {
          x: {
            grid: {
              display: false
            },
            ticks: {
              display: false
            }
          },
          y: {
            min: 0,
            max: 100,
            grid: {
              color: 'rgba(255, 255, 255, 0.05)',
            },
            border: {
              dash: [4, 4]
            },
            beginAtZero: true,
            ticks: {
                beginAtZero: true,
                color: 'rgba(255, 255, 255, 0.5)',
                font: {
                  size: 10
                }
            }
          }
        }
      },
      optionsNet: {
        animation: {
          duration: 300
        },
        responsive: true,
        maintainAspectRatio: false,
        interaction: {
          intersect: false,
          mode: 'index',
        },
        plugins: {
          tooltip: {
            enabled: true,
            backgroundColor: 'rgba(13, 17, 26, 0.85)',
            titleColor: '#00F0FF',
            bodyColor: '#ffffff',
            borderColor: 'rgba(255, 255, 255, 0.1)',
            borderWidth: 1,
            padding: 8,
            cornerRadius: 8,
          },
          legend: {
              display: false,
          }
        },
        scales: {
          x: {
            grid: {
              display: false
            },
            ticks: {
              display: false
            }
          },
          y: {
            grid: {
              color: 'rgba(255, 255, 255, 0.05)',
            },
            border: {
              dash: [4, 4]
            },
            beginAtZero: true,
            ticks: {
              color: 'rgba(255, 255, 255, 0.5)',
              font: {
                size: 10
              },
              callback: (label:any, index: number) => { return parseInt(label).toString() },
              count: 6
            }
          }
        }
      },
      data: ref(<any>{})
    }
  },
  computed: {
    options() {
      switch (this.$props.type){
        case "h-net":
          this.optionsNet.scales.y.ticks.callback = (label:any, index: number) => {
            return label == 0 ? "0" : HumanReadable.sizeFormat(label,0)
          }
          return this.optionsNet
        case "hp-net":
          this.optionsNet.scales.y.ticks.callback = (label:any, index: number) => {
            return label == 0 ? "0" : HumanReadable.packetFormat(label,0)
          }
          return this.optionsNet
        case "h-dio":
          this.optionsNet.scales.y.ticks.callback = (label:any, index: number) => {
            return label == 0 ? "0" : HumanReadable.sizeFormat(label,0)
          }
          return this.optionsNet
      }
      return this.options1
    }
  },
  methods: {
    updateData1(value1: number) {
      const newData = <number[]>[]
      if (this.data.datasets){
        newData.push(...this.data.datasets[0].data,value1)
      }
      if (newData.length>20) newData.shift()
      this.data = {
        labels: this.labels,
        datasets: [
          {
            label: '',
            backgroundColor: 'rgba(0, 240, 255, 0.08)',
            borderColor: 'rgba(0, 240, 255, 0.85)',
            borderWidth: 2,
            pointRadius: 1.5,
            pointHoverRadius: 4,
            pointBackgroundColor: '#00F0FF',
            tension: 0.45,
            fill: true,
            data: newData
          }
        ],
      }
      this.loaded = true
    },
    updateData2(value1: number, value2:number) {
      const newData1 = <number[]>[]
      const newData2 = <number[]>[]
      if (this.data.datasets){
        newData1.push(...this.data.datasets[0].data,value1)
        newData2.push(...this.data.datasets[1].data,value2)
      }
      if (newData1.length>20) newData1.shift()
      if (newData2.length>20) newData2.shift()
      this.data = {
        labels: this.labels,
        datasets: [
          {
            label: 'Upload',
            backgroundColor: 'rgba(191, 90, 242, 0.05)',
            borderColor: 'rgba(191, 90, 242, 0.8)',
            borderWidth: 2,
            pointRadius: 1.5,
            pointHoverRadius: 4,
            pointBackgroundColor: '#bf5af2',
            tension: 0.45,
            fill: true,
            data: newData1
          },
          {
            label: 'Download',
            backgroundColor: 'rgba(0, 240, 255, 0.08)',
            borderColor: 'rgba(0, 240, 255, 0.85)',
            borderWidth: 2,
            pointRadius: 1.5,
            pointHoverRadius: 4,
            pointBackgroundColor: '#00F0FF',
            tension: 0.45,
            fill: true,
            data: newData2
          }
        ],
      }
      this.loaded = true
    }
  },
  watch: {
    tilesData(v:any) {
      switch (this.$props.type) {
        case 'h-cpu':
          this.updateData1(v.cpu)
          break
        case 'h-mem':
          this.updateData1(v.mem.current*100/v.mem.total)
          break
        case 'h-net':
          if (this.oldValues.net.sent) {
            const downSpeed = (v.net.recv-this.oldValues.net.recv)/2  // Each 2 sec
            const upSpeed = (v.net.sent-this.oldValues.net.sent)/2  // Each 2 sec
            this.updateData2(upSpeed,downSpeed)
          }
          this.oldValues.net = v.net
          break
        case 'hp-net':
          if (this.oldValues.net.psent) {
            const downSpeed = (v.net.precv-this.oldValues.net.precv)/2  // Each 2 sec
            const upSpeed = (v.net.psent-this.oldValues.net.psent)/2  // Each 2 sec
            this.updateData2(upSpeed,downSpeed)
          }
          this.oldValues.net = v.net
          break
        case 'h-dio':
          if (this.oldValues.dio.read) {
            const downSpeed = (v.dio.read-this.oldValues.dio.read)/2  // Each 2 sec
            const upSpeed = (v.dio.write-this.oldValues.dio.write)/2  // Each 2 sec
            this.updateData2(upSpeed,downSpeed)
          }
          this.oldValues.dio = v.dio
          break
      }
    }
  }
}
</script>