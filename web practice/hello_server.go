package main

import (
	"log"
	"net/http"
)

type hello struct {
}

func (hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}

func main() {
	var h hello
	http.Handle("/", h)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
