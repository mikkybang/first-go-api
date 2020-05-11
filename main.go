package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

)


func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte (`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "put called"}`))
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/" post).Methods(http.MethodPost)
	r.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":7070", nil))
}