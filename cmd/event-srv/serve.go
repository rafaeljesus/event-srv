package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rafaeljesus/event-srv/pkg/checker"
	"github.com/rafaeljesus/event-srv/pkg/config"
	"github.com/rafaeljesus/event-srv/pkg/handlers"
	"github.com/rafaeljesus/event-srv/pkg/kafka-bus"
	m "github.com/rafaeljesus/event-srv/pkg/middleware"
	"github.com/rafaeljesus/event-srv/pkg/models"
	"github.com/rafaeljesus/event-srv/pkg/repos"
	"github.com/spf13/cobra"
)

func Serve(cmd *cobra.Command, args []string) {
	log.WithField("version", version).Info("Event Service starting...")

	env, err := config.LoadEnv()
	failOnError(err, "Failed to load config!")

	level, err := log.ParseLevel(strings.ToLower(env.LogLevel))
	failOnError(err, "Failed to get log level!")
	log.SetLevel(level)

	ds, err := datastore.New(env.DatastoreURL)
	failOnError(err, "Failed to init dababase connection!")
	defer ds.Close()

	e, err := kafkabus.NewEmitter(kafkabus.EmitterConfig{
		Url:      globalConfig.BrokerURL,
		Attempts: globalConfig.ProducerAttempts,
		Timeout:  globalConfig.ProducerTimeout,
	})
	failOnError(err, "Failed to init kafka connection!")
	defer e.Close()

	checkers := map[string]checker.Checker{
		"api":     checker.NewApi(),
		"elastic": checker.NewElastic(globalConfig.DatastoreURL),
		"kafka":   checker.NewKafka(globalConfig.BrokerURL),
	}
	eventRepo := repos.NewEvent(ds)

	statusHandler := handlers.NewStatusHandler(checkers)
	eventsHandler := handlers.NewEventsHandler(eventRepo, e)

	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(&m.LoggerRequest{}))
	r.Use(middleware.Recoverer)

	r.Get("/status", statusHandler.StatusIndex)

	r.Route("/events", func(r chi.Router) {
		r.Get("/", eventsHandler.EventIndex)
		r.Post("/", eventsHandler.EventsCreate)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", globalConfig.Port), r))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithError(err).Panic(msg)
	}
}
