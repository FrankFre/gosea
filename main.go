package main

import (
	"github.com/FrankFre/gosea/status" // ein beliebiger Pfad lokal, mit go init
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//Sesion vom 29.5.20
func main() {

	/*logfile, err := os.Create("messages.log")
	if err != nil {
		log.Fatal("error opening log file: %s", err.Error())
	}

	_ = logfile
	defer func() {
		log.Print("closing log file")
		logfile.Close()
	}()*/

	logger := log.New(os.Stdout, "gosea ", log.LstdFlags)

	sigChan := make(chan os.Signal) //Channel mit 1 Signal
	defer close(sigChan)            //bewusstes Schließen des Channel
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux() // neues Objekt, hier ohne "make"
	mux.HandleFunc("/health", status.Health)

	srv := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	defer srv.Close()

	go func() {
		err := srv.ListenAndServe()
		//if err != nil && !errors.Is(err, http.ErrServerClosed) {
		if err != nil {
			logger.Fatalf("error starting server: %s", err.Error())
		}
	}()

	logger.Print("starting service")

	<-sigChan //Pfeil heist lESEN

	logger.Print("stopping service")

}
