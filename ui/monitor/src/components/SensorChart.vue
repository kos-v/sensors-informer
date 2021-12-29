<script>
import { Line, mixins } from 'vue-chartjs'

export default {
  extends: Line,
  name: "SensorChart",
  props: ["value", "timestamp"],
  mixins: [mixins.reactiveData],
  data: function () {
    return {
      labelsStorage: [],
      valuesStorage: [],
      options: {
        elements: {
          point:{
            radius: 0
          }
        },
        legend: {
          display: false
        },
        plugins: {
          datalabels: {
            display: false,
          },
        },
        scales: {
          xAxes: [{display: false}],
          yAxes: [{
            ticks: {
              suggestedMin: 20,
              suggestedMax: 100
            }
          }],
        },
        tooltips: {
          enabled: false
        },
      }
    }
  },
  mounted () {
    this.renderChart(this.chartData, this.options)
  },
  watch: {
    timestamp: function () {
      let valuesStorage = this.valuesStorage;
      valuesStorage.push(this.value);
      let labelsStorage = this.labelsStorage;
      labelsStorage.push(this.valuesStorage.length);

      let popLen = this.valuesStorage.length - 60;
      if (popLen > 0) {
        valuesStorage = this.valuesStorage.slice(popLen);
        labelsStorage = this.labelsStorage.slice(popLen)
      }

      this.valuesStorage = valuesStorage;
      this.labelsStorage = labelsStorage;


      this.chartData = {
        labels: this.labelsStorage,
        datasets: [{
          borderColor: '#17a2b8',
          data: this.valuesStorage,
          backgroundColor: '#b1ecfa',
        }]
      }
    }
  }
}
</script>