package main

import (
	"log"
	"net/http"
	"os/exec"
)

const (
	PORT = "8080"
	URL  = "https://github.com/Juicestus/TorqueLearn"
)

func Execute(command string, args ...string) {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}

func Fetch() {
	exists, err := FileExists("./local")
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		Execute("cd", "./local", "&&", "git", "pull")
	} else {
		Execute("git", "clone", URL, "./local")
	}
}

func main() {

	Fetch()

	http.Handle("/", http.FileServer(http.Dir("./deploy")))

	// if err := http.ListenAndServe(":"+PORT, nil); err != nil {
	// 	log.Fatal(err)
	// }

}
