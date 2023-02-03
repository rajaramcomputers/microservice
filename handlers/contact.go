package handlers

import (
	"log"
	"net/http"
)

type Contact struct {
	l *log.Logger
}

func NewContact(l *log.Logger) *Contact {
	return &Contact{l}
}

func (g *Contact) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("From the Contact\n"))
}
