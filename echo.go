package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form.Get("command"))
}

func main() {

	port := flag.Int("port", 8000, "tcp port to listen")
	flag.Parse()

	http.HandleFunc("/echo", echo)
	log.Printf("Listening on :%d", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
