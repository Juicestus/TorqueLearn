package main

import (
	"log"
	"net/http"
)

const (
	PORT = "8080"
	URL  = "https://github.com/Juicestus/TorqueLearn"
)

func main() {

	GenerateStaticFiles()

	http.Handle("/", http.FileServer(http.Dir("./deploy")))

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
