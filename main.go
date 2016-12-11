package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/trace-srv/event_bus"
	"github.com/rafaeljesus/trace-srv/handlers"
	"github.com/rafaeljesus/trace-srv/models"
	"github.com/spf13/viper"
)

const (
	trace_srv_db   = "TRACE_SRV_DB"
	trace_srv_port = "TRACE_SRV_PORT"
	trace_srv_bus  = "TRACE_SRV_BUS"
)

func main() {
	viper.AutomaticEnv()

	db, err := models.newDB(viper.GetString(trace_srv_db))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to init database connection!")
	}

	event_bus, err := event_bus.NewEventBus(viper.GetString(trace_srv_bus))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to init event bus!")
	}

	event_bus.On("events", handlers.EventCreated)

	env := &handlers.Env{db, event_bus}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	r := e.Group("/v1")
	r.GET("/healthz", env.HealthIndex)
	r.GET("/events", env.EventsIndex)
	r.POST("/events", env.EventsCreate)

	log.WithFields(log.Fields{
		"port": viper.GetString(trace_srv_port),
	}).Info("Starting Trace Service")

	e.Run(fasthttp.New(":" + viper.GetString(trace_srv_port)))
}
