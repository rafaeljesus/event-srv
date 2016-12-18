package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/event-srv/event_bus"
	"github.com/rafaeljesus/event-srv/handlers"
	"github.com/rafaeljesus/event-srv/models"
	"github.com/spf13/viper"
)

const (
	event_srv_db   = "EVENT_SRV_DB"
	event_srv_port = "EVENT_SRV_PORT"
	event_srv_bus  = "EVENT_SRV_BUS"
)

func main() {
	viper.AutomaticEnv()

	db, err := models.NewDB(viper.GetString(event_srv_db))
	if err != nil {
		log.WithError(err).Fatal("Failed to init database connection!")
	}

	event_bus, err := event_bus.NewEventBus(viper.GetString(event_srv_bus))
	if err != nil {
		log.WithError(err).Fatal("Failed to init event bus!")
	}

	env := &handlers.Env{db, event_bus}

	if err := event_bus.On("events", env.EventCreated); err != nil {
		log.WithError(err).Fatal("Failed to listen for events topic!")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	r := e.Group("/v1")
	r.GET("/healthz", env.HealthIndex)
	r.GET("/events", env.EventsIndex)
	r.POST("/events", env.EventsCreate)

	log.WithField("port", viper.GetString(event_srv_port)).Info("Starting event Service")

	e.Run(fasthttp.New(":" + viper.GetString(event_srv_port)))
}
