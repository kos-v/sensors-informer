package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	conf "github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/sensor"
	"github.com/kos-v/sensors-informer/internal/temperature"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := flag.String("addr", ":80", "TCP address for the server to listen on, in the form \"host:port\"")
	specificConfig := flag.String("config", "", "The path to a specific configuration file")
	tmplDir := flag.String("tmplDir", "./web/monitor/templates", "Path to html templates")
	staticsDir := flag.String("staticsDir", "./web/monitor/dist", "Path to static files")
	flag.Parse()

	config, err := conf.LoadConfig(*specificConfig)
	if err != nil {
		log.Fatalf("erro: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob(*tmplDir + "/*")
	router.Static("/dist", *staticsDir)

	sr := &sensor.CommandReader{Command: config.LmSensors.Command}
	router.GET("/api/v1/indicators/status", func(c *gin.Context) {
		busList, err := sr.ReadTemperatureSensors()
		if err != nil {
			log.Printf("error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"code":    http.StatusInternalServerError,
				"error":   "server error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"code":    http.StatusOK,
			"result": gin.H{
				"indications": formatSensorsOut(busList),
				"temperature": gin.H{"unit": temperature.UnitCelsius},
				"timestamp":   time.Now().Unix(),
			},
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(*addr)
}

func formatSensorsOut(busList []sensor.Bus) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, b := range busList {
		bResult := map[string]interface{}{
			"name": b.Name,
			"type": sensor.SensorTypeTemperature,
		}

		sResult := []map[string]interface{}{}
		for _, s := range b.Sensors {
			if !s.HasInputInfo() {
				continue
			}

			sResult = append(sResult, map[string]interface{}{
				"name":  s.Name,
				"value": s.GetInputInfo().GetValueAsInt(),
			})
		}
		bResult["sensors"] = sResult

		if len(sResult) > 0 {
			result = append(result, bResult)
		}
	}

	return result
}
