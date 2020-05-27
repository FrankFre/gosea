package main

import (
	"github.com/FrankFre/gosea/status"
	"net/http"
)

func main() {

	mux := http.NewServeMux() // neues Objekt, hier ohne "make"
	mux.HandleFunc("/health", status.Health)


	srv := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	srv.ListenAndServe() // mit Fehlerbehandlung siehe Docs


}


