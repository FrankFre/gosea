package main

import (
	"errors"
	"github.com/FrankFre/gosea/api"
	"github.com/FrankFre/gosea/posts"
	"github.com/FrankFre/gosea/status" // ein beliebiger Pfad lokal, mit go init
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// initialize logger
	logfile, err := os.Create("messages.log")
	if err != nil {
		log.Fatalf("error opening log file: %s", err.Error())
	}

	defer func() {
		log.Print("closing log file")
		logfile.Close()
	}()

	logger := log.New(os.Stdout, "gosea ", log.LstdFlags)

	// init signal handling
	sigChan := make(chan os.Signal) //Channel mit 1 Signal
	defer close(sigChan)            //bewusstes Schließen des Channel
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// create services
	postsService := posts.NewWithSea()
	apiService := api.New(postsService)

	mux := http.NewServeMux() // neues Objekt, hier ohne "make"
	mux.HandleFunc("/health", status.Health)
	mux.HandleFunc("/api", apiService.Posts)

	srv := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("error starting server: %s", err.Error())
		}
	}()

	logger.Print("starting service")

	<-sigChan //Pfeil heist lESEN

	logger.Print("stopping service")

}
