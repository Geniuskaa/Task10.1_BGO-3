package main

import (
	"github.com/Geniuskaa/Task10.1_BGO-3/cmd/app"
	"github.com/Geniuskaa/Task10.1_BGO-3/cmd/card"
	"log"
	"net/http"
	"os"
)

func main() {


	if  err := execute("localhost:9999"); err != nil {
		log.Println(err)
		os.Exit(1)
	}


}


func execute(addr string) (err error) {
	cardSvc := card.NewService()
	mux := http.NewServeMux()
	application := app.NewServer(cardSvc, mux)
	application.Init()

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
	}
	return server.ListenAndServe()
}
