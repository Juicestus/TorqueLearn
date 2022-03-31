package main

import (
	"log"
	"net/http"
)

const (
	PORT = "8080"
	URL  = "https://github.com/Juicestus/TorqueLearn"
)

func ServeStaticFiles(port string) {
	http.Handle("/", http.FileServer(http.Dir("./deploy")))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./deploy")))

	// if err := http.ListenAndServe(":"+PORT, nil); err != nil {
	// 	log.Fatal(err)
	// }

}
