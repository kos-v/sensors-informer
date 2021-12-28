<template>
  <div class="py-2">
    <div class="card">
      <div class="card-header bg-white border-0">
        <span class="card-title">{{ name }}</span>
        <div class="card-tools">
          <span><span class="badge text-white" v-bind:class="indicatorClass()">{{ value }}°С</span></span>
        </div>
      </div>
      <div class="card-body">
        <SensorChart
            v-bind:chart-id="chartId"
            v-bind:height="196"
            v-bind:timestamp="timestamp"
            v-bind:value="value"></SensorChart>
      </div>
    </div>
  </div>
</template>

<script>
import SensorChart from './SensorChart'

export default {
  name: "Sensor",
  components: {
    SensorChart
  },
  props: ["name", "value", "timestamp"],
  data: function () {
    return {
      chartId: null
    }
  },
  methods: {
    indicatorClass() {
      if (this.value < 60) {
        return "bg-success";
      } else if (this.value < 75) {
        return "bg-warning";
      }
      return "bg-danger";
    }
  },
  mounted() {
    this.chartId = "chart" + this._uid;
  }
}
</script>

<style scoped>
.card-tools {
  float: right;
}
</style>