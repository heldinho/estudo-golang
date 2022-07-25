package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func _home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(r.Header)
}

func _init() {
	http.HandleFunc("/", _home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	_init()
}
