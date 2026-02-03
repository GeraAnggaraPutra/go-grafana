package app

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339

	_ = os.MkdirAll("./logs", 0755)

	file, err := os.OpenFile(
		"./logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	zlog.Logger = zerolog.New(file).
		With().
		Timestamp().
		Str("service", os.Getenv("GRAFANA_NAME")).
		Logger()
}
