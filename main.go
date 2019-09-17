package main

import "net/http"

var addr = ":8080"

func main() {
	http.HandleFunc("/", root)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(":-) version 2"))
}
