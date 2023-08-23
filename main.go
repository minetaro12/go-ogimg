package main

import (
	"fmt"
	"go-ogimg/handlers"
	"log"
	"net/http"
	"os"
)

var (
	httpListen = fmt.Sprintf(":%v", getEnv("PORT", "8000"))
)

func init() {
	// フォントファイルがなければ起動しない
	if !fileExists("font.ttf") {
		log.Fatal("font.ttf not found")
	}
}

func main() {
	http.HandleFunc("/", handlers.Root)

	log.Println("Server Listening on", httpListen)
	log.Fatal(http.ListenAndServe(httpListen, logRequest(http.DefaultServeMux)))
}

func getEnv(key, fallback string) string {
	if val, isFound := os.LookupEnv(key); isFound {
		return val
	}
	return fallback
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
