package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

func main() {
	router := httprouter.New()
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprintln(w, "Error : ", i)
	}

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Method not allowed")
	})

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL not found")
	})

	router.GET("/ups", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Upss")
	})

	router.POST("/customers", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Success create customer")
	})

	router.GET("/customers/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "customers with id " + p.ByName("id")
		fmt.Fprintln(w, text)
	})

	router.ServeFiles("/storage/*filepath", http.FS(directory))

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
