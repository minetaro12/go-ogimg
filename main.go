package main

import (
	"fmt"
	"go-ogimg/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	gin.SetMode("release")
	r := gin.Default()
	r.GET("/", handlers.Root)

	log.Println("Server Listening on", httpListen)
	err := r.Run(httpListen)
	if err != nil {
		log.Fatal(err)
	}
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
