<template>
  <div id="app">
    <Header/>
    <div class="container-fluid">
      <SensorGroup
          v-bind:groups="groups"
          v-bind:timestamp="timestamp"></SensorGroup>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './components/Header'
import SensorGroup from './components/SensorGroup'

export default {
  name: 'App',
  components: {
    Header,
    SensorGroup
  },
  data: function () {
    return {
      groups: [],
      timestamp: 0,
    }
  },
  methods: {
    updateSensors() {
      let self = this;
      axios.get('/api/v1/indicators/status')
          .then(function (res) {
            if (!res.data.success || !Array.isArray(res.data.result.indications)) {
              return
            }

            self.groups = res.data.result.indications;
            self.timestamp = res.data.result.timestamp;
          })
          .catch(function (error) {
            console.log(error);
          })
    }
  },
  mounted: function () {
    let self = this;
    setInterval(function () {
      self.updateSensors()
    }, 3000)
  }
}
</script>

<style>
@import '~bootstrap/dist/css/bootstrap.css';
</style>
