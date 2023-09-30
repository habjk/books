package main

import (
	"fmt"
	"golang-crud/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	hb()
}

func hb() {
	fmt.Printf("Starting main\n")
	routes := httprouter.New()
	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome to GO.")
	})
	server := http.Server{Addr: "localhost:8888", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfEror(err)

}
