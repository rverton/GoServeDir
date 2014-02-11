package main

import (
	"flag"
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

	port := flag.String("port", "8080", "Port to listen to. Default 8080")
	path := flag.String("path", current, "Path to serve. Default is current dir")
	flag.Parse()

	log.Printf("Listening for '%v' at :%v", *path, *port)

	panic(http.ListenAndServe(fmt.Sprintf(":%v", *port), &logHandler{http.FileServer(http.Dir(*path))}))
}
