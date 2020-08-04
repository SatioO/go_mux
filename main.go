package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {

}
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {}
func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	w.Write(response)
}
func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {}
func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.getProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.createProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func main() {
	a := App{}
	a.Run(":3000")
}
