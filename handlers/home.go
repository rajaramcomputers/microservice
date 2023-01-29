package handlers

import (
	"log"
	"net/http"
)

type Home struct {
	l *log.Logger
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (g *Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("From Home"))
}
