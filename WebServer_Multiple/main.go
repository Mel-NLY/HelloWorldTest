package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/inc", incHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//To convert server for routing use this main instead
/*
func main(){
	err := http.ListenAndServe(":8080", handlers())
	if err != nil {
		log.Fatal(err)
	}
}

func handlers() http.Handler{
	h := http.NewServeMux()
	h.HandleFunc("/inc", incHandler)
	return h
}
*/

func incHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	if input == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	val, err := strconv.Atoi(input)
	if err != nil {
		http.Error(w, "not a number: "+input, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, val+1)
}
