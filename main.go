package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type logHandler struct {
	http.Handler
}

func (h logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v - %v", time.Now().Format(time.Stamp), r.URL)
	h.Handler.ServeHTTP(w, r)
}

func main() {
	current, _ := os.Getwd()
	log.Printf("Listening for '%v' at :8080", current)
	panic(http.ListenAndServe(":8080", &logHandler{http.FileServer(http.Dir(current))}))
}
