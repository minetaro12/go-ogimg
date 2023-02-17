package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	httpListen = fmt.Sprintf(":%v", getEnv("PORT", "8000"))
)

func main() {
	log.Println("Server Listening on", httpListen)
	http.HandleFunc("/ogimg.jpg", ogimgHandle)
	log.Fatal(http.ListenAndServe(httpListen, logRequest(http.DefaultServeMux)))
}

func getEnv(key, fallback string) string {
	if val, isFound := os.LookupEnv(key); isFound {
		return val
	}
	return fallback
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func errorResponse(w http.ResponseWriter) {
	http.Error(w, "Error", http.StatusInternalServerError)
}
