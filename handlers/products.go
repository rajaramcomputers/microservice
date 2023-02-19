package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"rajaramsystems.com/microservice/handlers/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		//expect the ID in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		//p.l.Println("got id", id)

		p.UpdateProducts(id, rw, r)
		return

	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusBadRequest)
	}
	//p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
