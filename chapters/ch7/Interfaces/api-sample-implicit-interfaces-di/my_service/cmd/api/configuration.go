package main

import (
	"fmt"
	"os"
	"time"
)

type ServerConfiguration struct {
	HTTPAddr          string
	DatabaseURL       string
	ShutdownTimeout   time.Duration
	OutboxRelayOn     bool
	LogLevel          string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func LoadConfig() (ServerConfiguration, error) {
	var errs []error
	srvCfg := ServerConfiguration{
		HTTPAddr:          getEnvVarOrDefault("HTTP_ADDR", ":8080", &errs),
		DatabaseURL:       getEnvVar("DATABASE_URL", *errs),
		ShutdownTimeout:   getEnvDurationVarOrDefault("SHUTDOWN_TIMEOUT", 10*time.Second, &errs),
		OutboxRelayOn:     getEnvBooleanVarOrDefault("OUTBOX_RELAY", false, &errs),
		LogLevel:          getEnvVarOrDefault("LOG_LEVEL", "info", &errs),
		ReadTimeout:       getEnvDurationVarOrDefault("READ_TIMEOUT", 10*time.Second, &errs),
		ReadHeaderTimeout: getEnvDurationVarOrDefault("READ_HEADER_TIMEOUT", 7*time.Second, &errs),
		WriteTimeout:      getEnvDurationVarOrDefault("WRITE_TIMEOUT", 30*time.Second, &errs),
		IdleTimeout:       getEnvDurationVarOrDefault("IDLE_TIMEOUT", 10*time.Second, &errs),
	}

	if errs != nil {
		fmt.Fprintln(os.Stderr, "config:", errs)
		os.Exit(1)
	}

	return srvCfg, nil
}

func getEnvVar(name string, errors *[]error) string {

}
func getEnvBooleanVarOrDefault(name string, defaultBoolean bool, errors *[]error) bool {
}
func getEnvVarOrDefault(name string, defaultValue string, errors *[]error) string {
}
func getEnvDurationVarOrDefault(name string, duration time.Duration, errors *[]error) time.Duration {
}

func loadConfigurationFile(filename string) (ServerConfiguration, error) {
}
