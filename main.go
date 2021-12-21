package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := fmt.Fprintf(w, "Request: %s", body); err != nil {
			log.Fatal(err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
