package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kos-v/sensors-informer/api/monitor/v1/controller"
	"github.com/kos-v/sensors-informer/internal/sensor"
)

type API struct {
	Router  *gin.Engine
	SReader sensor.Reader
}

func (a *API) Handle() {
	group := a.Router.Group("/api/v1")
	group.GET("/indicators/status", func(ctx *gin.Context) {
		c := controller.IndicatorsStatusController{SReader: a.SReader}
		c.Exec(ctx)
	})
}
