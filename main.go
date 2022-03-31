// This is a bruh moment

package main

import (
	"log"
	"net/http"
)

const (
	PORT = "8080"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./deploy")))

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}

}
