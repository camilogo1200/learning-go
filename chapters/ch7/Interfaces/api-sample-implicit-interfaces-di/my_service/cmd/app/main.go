package main

import (
	"errors"
	"fmt"
	"learning-go-ch7-api-example-implicit-intefaces-di/my_service/internal/config"
	"net/http"
	"os"
)

func main() {

	srvConfiguration, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load configuration:\n%v\n", err)
		os.Exit(1)
	}

	//# wiring, flags, signal handling.
	logger, err := config.CreateLogger()

	if err != nil {
		fmt.Println(err.Error())
	}

	mux := http.NewServeMux()
	//_, ok := createRoutes(mux)

	//if ok != nil {
	//
	//}

	server := &http.Server{
		Addr:              srvConfiguration.HTTPAddr,
		Handler:           mux,
		ReadTimeout:       srvConfiguration.ReadTimeout,
		ReadHeaderTimeout: srvConfiguration.ReadHeaderTimeout,
		WriteTimeout:      srvConfiguration.WriteTimeout,
		IdleTimeout:       srvConfiguration.IdleTimeout,
	}

	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error())
		}
	}()

}
