<template>
  <div class="d-grid">
    <div class="card border-dark mb-3" v-for="group in groups" v-if="group.sensors.length>0">
      <div class="card-header border-dark">{{ group.name }}</div>
      <div class="card-body">
        <div class="row row-cols-5">
          <div class="col" v-for="s in group.sensors">
            <Sensor
                v-bind:name="s.name"
                v-bind:timestamp="timestamp"
                v-bind:value="s.value"></Sensor>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Sensor from './Sensor'

export default {
  name: "SensorGroup",
  props: ["groups", "timestamp"],
  components: {
    Sensor
  }
}
</script>