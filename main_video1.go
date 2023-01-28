package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello Ram")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Data %s\n", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye Ram")
	})

	http.ListenAndServe(":9090", nil)

}
