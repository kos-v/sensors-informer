package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	v1 "github.com/kos-v/sensors-informer/api/monitor/v1"
	conf "github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/sensor"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":80", "TCP address for the server to listen on, in the form \"host:port\"")
	specificConfig := flag.String("config", "", "The path to a specific configuration file")
	tmplDir := flag.String("tmplDir", "./ui/monitor/templates", "Path to html templates")
	staticsDir := flag.String("staticsDir", "./ui/monitor/dist", "Path to static files")
	flag.Parse()

	config, err := conf.LoadConfig(*specificConfig)
	if err != nil {
		log.Fatalf("erro: %v", err)
	}

	router := gin.Default()
	api := v1.API{
		Router:  router,
		SReader: &sensor.CommandReader{Command: config.LmSensors.Command},
	}
	api.Handle()

	router.LoadHTMLGlob(*tmplDir + "/*")
	router.Static("/dist", *staticsDir)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(*addr)
}
