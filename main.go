package main

import (
	"fmt"
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
	fmt.Println("Listening at :8080")
	panic(http.ListenAndServe(":8080", &logHandler{http.FileServer(http.Dir(current))}))
}
