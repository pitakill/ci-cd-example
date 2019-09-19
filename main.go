package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		panic("$PORT must be set")
	}

	http.HandleFunc("/", root)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(":-) version 5"))
}
