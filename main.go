package main

import (
	"fmt"
	"log"
	"os"

	"github.com/minetaro12/go-ogimg/handlers"

	"github.com/gofiber/fiber/v2"
)

var (
	httpListen = fmt.Sprintf(":%v", getEnv("PORT", "8000"))
)

func init() {
	// フォントファイルがなければ起動しない
	if !fileExists("font.otf") {
		log.Fatal("font.ttf not found")
	}
}

func main() {
	app := fiber.New()
	app.Get("/", handlers.Root)

	log.Println("Server Listening on", httpListen)
	err := app.Listen(httpListen)
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
