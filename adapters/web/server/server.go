package server

import (
	"github/arlenmendes/hexagonal-arq-studies/adapters/web/handler"
	"github/arlenmendes/hexagonal-arq-studies/application"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webservice struct {
	Service application.ProductServiceInterface
}

func MakeNewWebservice() *Webservice {
	return &Webservice{}
}

func (w *Webservice) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	serve := &http.Server{
		Addr:              ":9000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
