package handlers

import (
	"log"
	"net/http"

	"rajaramsystems.com/microservice/handlers/data"
)

type Goals struct {
	l *log.Logger
}

func NewGoals(l *log.Logger) *Goals {
	return &Goals{l}
}

func (g *Goals) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		g.GetGoals(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (g *Goals) GetGoals(rw http.ResponseWriter, r *http.Request) {
	gp := data.GetGoals()
	err := gp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
