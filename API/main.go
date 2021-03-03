package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func main() {

	log.Println("Server starting")
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Request for /")
		s := RandomString(12)
		log.Println("Returning: ", s)
		io.WriteString(w, s)
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Request for /big-payload")
		s := RandomString(1500)
		log.Println("Returning rand string 1500")
		io.WriteString(w, s)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/big-payload", h2)
	log.Println("Listening on 0.0.0.0:80")
	log.Fatal(http.ListenAndServe(":80", nil))

}
