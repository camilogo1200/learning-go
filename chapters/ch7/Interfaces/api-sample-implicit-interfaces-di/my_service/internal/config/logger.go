package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

func CreateLogger() (*slog.Logger, error) {
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	replacer := strings.NewReplacer(":", "_", " ", "_")
	logFileName := replacer.Replace(time.Now().Format(time.RFC3339))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		message := fmt.Sprintf("Failed to open log file = [%s].\n", logFileName)
		return nil, errors.New(message)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Printf("Failed to close log file = [%s].\n", logFileName)
		}
	}(logFile)

	jsonHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	multiLogHandler := slog.NewMultiHandler(consoleHandler, jsonHandler)
	logger := slog.New(multiLogHandler)
	slog.SetDefault(logger)

	return logger, nil
}
