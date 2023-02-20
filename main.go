package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"rajaramsystems.com/microservice/handlers/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)

	//this handler was created by Dr Ram
	//homeh := handlers.NewHome(l)
	//ch := handlers.NewContact(l)
	//

	//from the lecture with respect to product handler

	ph := handlers.NewProducts(l)
	//goh := handlers.NewGoals(l)

	sm := mux.NewRouter()
	sm.Handle("/", ph)
	//sm.Handle("/goodbye", gh)
	//sm.Handle("/home", homeh)
	//sm.Handle("/contact", ch)
	//sm.Handle("/products", ph)
	//sm.Handle("/goals", goh)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Retreived terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

}
