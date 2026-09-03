package main

import "time"

type Configuration struct {
	HTTPAddr        string        `env:"HTTP_ADDR" envDefault:":8080"`
	DatabaseURL     string        `env:"DATABASE_URL,required"`
	shutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
	OutboxRelayOn   bool          `env:"OUTBOX_RELAY" envDefault:"false"`
}
