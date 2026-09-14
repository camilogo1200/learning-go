package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

type configurationLoader struct {
	errs []error
}

func (cl *configurationLoader) requiredString(name string) string {
	value, err := getEnvVar(name)
	if err != nil {
		cl.errs = append(cl.errs, err)
		return ""
	}
	return value
}

func (cl *configurationLoader) stringOrDefault(name string, defaultValue string) string {
	value, err := getEnvVarOrDefault(name, defaultValue)
	if err != nil {
		cl.errs = append(cl.errs, err)
		return defaultValue
	}
	return value
}

func (cl *configurationLoader) durationOrDefault(name string, durationDefault time.Duration) time.Duration {
	value, err := getEnvDurationVarOrDefault(name, durationDefault)
	if err != nil {
		cl.errs = append(cl.errs, err)
		return durationDefault
	}
	return value
}

func (cl *configurationLoader) boolOrDefault(name string, boolDefault bool) bool {
	value, err := getEnvBooleanVarOrDefault(name, boolDefault)
	if err != nil {
		cl.errs = append(cl.errs, err)
		return boolDefault
	}
	return value
}

func LoadConfig() (ServerConfiguration, error) {

	loader := configurationLoader{}

	srvCfg := ServerConfiguration{
		HTTPAddr:          loader.stringOrDefault("HTTP_ADDR", ":8080"),
		DatabaseURL:       loader.requiredString("DATABASE_URL"),
		ShutdownTimeout:   loader.durationOrDefault("SHUTDOWN_TIMEOUT", 10*time.Second),
		OutboxRelayOn:     loader.boolOrDefault("OUTBOX_RELAY", false),
		LogLevel:          loader.stringOrDefault("LOG_LEVEL", "info"),
		ReadTimeout:       loader.durationOrDefault("READ_TIMEOUT", 10*time.Second),
		ReadHeaderTimeout: loader.durationOrDefault("READ_HEADER_TIMEOUT", 7*time.Second),
		WriteTimeout:      loader.durationOrDefault("WRITE_TIMEOUT", 30*time.Second),
		IdleTimeout:       loader.durationOrDefault("IDLE_TIMEOUT", 10*time.Second),
	}

	err := errors.Join(loader.errs...)

	if err != nil {
		return ServerConfiguration{}, fmt.Errorf(
			"invalid server configuration: %w",
			err)
	}

	return srvCfg, nil
}

func getEnvVar(name string) (string, error) {
	if name == "" {
		err := &ErrInvalidConfiguration{
			Key:    name,
			Value:  "",
			Reason: "environment variable name cannot be empty",
			When:   time.Now(),
		}
		return "", err
	}
	value, exists := os.LookupEnv(name)
	if !exists {
		err := &ErrInvalidConfiguration{
			Key:    name,
			Value:  value,
			Reason: "environment variable does not exists",
			When:   time.Now(),
		}
		return "", err
	}
	return value, nil
}

func getEnvVarOrDefault(name string, defaultValue string) (string, error) {
	if name == "" {
		err := &ErrInvalidConfiguration{
			Key:    name,
			Value:  "",
			Reason: "environment variable name cannot be empty",
			When:   time.Now(),
		}
		return "", err
	}
	value, exists := os.LookupEnv(name)
	if !exists {
		return defaultValue, nil
	}
	return value, nil
}

func getEnvBooleanVarOrDefault(name string, defaultBoolean bool) (bool, error) {
	value, exists := os.LookupEnv(name)
	if !exists {
		return defaultBoolean, nil
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, &ErrInvalidConfiguration{
			Key:    name,
			Value:  value,
			Reason: "Value must be a valid boolean.",
			When:   time.Now(),
		}
	}

	return parsed, nil
}

func getEnvDurationVarOrDefault(name string, duration time.Duration) (time.Duration, error) {
	value, exists := os.LookupEnv(name)
	if !exists {
		return duration, nil
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, &ErrInvalidConfiguration{
			Key:    name,
			Value:  value,
			Reason: "Value must be a valid duration.",
			When:   time.Now(),
		}
	}
	return duration, nil
}

// maybe this should be used to set  the environment variables, and call the LoadConfig() method

func loadConfigurationFile(filename string) ServerConfiguration {
	serverConfig, err := LoadConfig()
	if err != nil {
	}
	return serverConfig
}
