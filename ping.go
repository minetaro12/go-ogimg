package main

import "net/http"

func pingHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
}
