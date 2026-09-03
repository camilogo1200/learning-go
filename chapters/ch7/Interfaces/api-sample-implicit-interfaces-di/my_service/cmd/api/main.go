package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//# wiring, flags, signal handling.
	logger, err := createLogger()

	if err != nil {
		fmt.Println(err.Error())
	}

	mux := http.NewServeMux()

	_, ok := createRoutes(mux)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 7 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       90 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error())
		}
	}()

}

func createLogger() (*slog.Logger, error) {
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	replacer := strings.NewReplacer(":", "_", " ", "_")
	logFileName := replacer.Replace(time.Now().Format(time.RubyDate))
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
