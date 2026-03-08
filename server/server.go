package server

import (
	"fmt"
	"go-prac/routers"
	"log"
	"net/http"
)

func Start() {
	// mux is like a router
	masterMux := http.NewServeMux()

	masterMux.HandleFunc("/health", healthHandler)
	masterMux.HandleFunc("/", homeHandler)
	masterMux.Handle("/user/", http.StripPrefix("/user", routers.UserRouter))

	port := ":8080"

	fmt.Println("Server is started at port", port)
	// if can get more than one statement in go, one usally to create a variable
	// and second for check condition for true and false
	// this is help in not populiting the global scope and keep err check related variables in
	// block scope, It is great approach to handle the error handling mesh of golang if the
	// variable is not required globally or in outer scope
	if err := http.ListenAndServe(port, masterMux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// use http.ResponseWriter to write the response and details of it like (Header, Status Code, Body)
	// header().set to set an header
	w.Header().Set("Content-Type", "application/json")
	// WriteHeader to write status code
	w.WriteHeader(http.StatusOK)
	// Write data to tcp connection
	w.Write([]byte(`{"message":"Welcome 👋"}`))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
