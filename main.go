package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := os.Getenv("SECRET1_FILE")
		data, err := os.ReadFile(fileName)
		if err != nil {
			msg := fmt.Sprintf("failed to read file %q with error: %v", fileName, err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	fmt.Println("Server starting on port 8443")
	http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
}
