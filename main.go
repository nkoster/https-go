package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const httpPort = "80"
const httpsPort = "443"

func main() {
	log.Printf("Web server listening to ports %s and %s\n", httpPort, httpsPort)
	http.HandleFunc("/", exampleHandler)
	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	err := http.ListenAndServeTLS(":"+httpsPort, "tls.pem", "tls.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "https://"+name, http.StatusMovedPermanently)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
