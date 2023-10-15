package main

import (
	"github/suda7kabo/myrouter2"
	"log"
	"net/http"
)

func main() {
	r := myrouter2.NewRouter()
	r.GET("/", http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	}))

	r.GET("/hoge", http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hoge"))
	}))
	log.Fatal(http.ListenAndServe(":8080", r))
}
