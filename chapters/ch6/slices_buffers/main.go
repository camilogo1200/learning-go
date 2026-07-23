package main

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"strings"
)

func main() {
	open_resource_traditional()
}

func open_resource_traditional() (string, error) {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r, err := os.Open("main.go")
	if err != nil {
		logger.Warn("Error opening file", err)
		return "", errors.New("cannot find file on directory")
	}
	defer func(r *os.File) {
		err := r.Close()
		if err != nil {
			logger.Error("error closing file", err)
		}
	}(r)

	bufferSize := 100
	buffer := make([]byte, bufferSize)
	var text []string
	for {
		countDataRead, readErr := r.Read(buffer)
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				logger.Debug("EOF", readErr)
				break
			} else {
				return "", readErr
			}
		}
		text = append(text, processData(buffer[:countDataRead]))
	}
	return strings.Join(text, " "), nil
}

func processData(bytes []byte) string {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	st := string(bytes)
	logger.Info(st)
	return st
}
