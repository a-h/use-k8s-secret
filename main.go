package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	})

	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	fmt.Println("Server starting on port 8443")
	http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
}
