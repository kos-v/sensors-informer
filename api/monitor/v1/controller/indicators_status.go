package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kos-v/sensors-informer/internal/sensor"
	"github.com/kos-v/sensors-informer/internal/temperature"
	"log"
	"net/http"
	"time"
)

type IndicatorsStatusController struct {
	SReader sensor.Reader
}

func (c *IndicatorsStatusController) Exec(ctx *gin.Context) {
	busList, err := c.SReader.ReadTemperatureSensors()
	if err != nil {
		log.Printf("error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"error":   "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"result": gin.H{
			"indications": c.formatSensorsOut(busList),
			"temperature": gin.H{"unit": temperature.UnitCelsius},
			"timestamp":   time.Now().Unix(),
		},
	})
}

func (c *IndicatorsStatusController) formatSensorsOut(busList []sensor.Bus) []map[string]interface{} {
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
				"value": s.GetInputInfo().GetValueAsFloat(),
			})
		}
		bResult["sensors"] = sResult

		if len(sResult) > 0 {
			result = append(result, bResult)
		}
	}

	return result
}
