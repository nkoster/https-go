package main

import (
	"log"
	"net/http"
	"os"
)

const httpPort = "80"
const httpsPort = "443"

func main() {
	log.Printf("Web server listening to ports %s and %s\n", httpPort, httpsPort)
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
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
