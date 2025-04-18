package main

import (
	"github.com/joho/godotenv"
	"github.com/yusadeol/go-budgeteer/internal/infra/http"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/handler"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpServer := http.NewChiServer()

	httpServer.Register(http.MethodGet, "/auth", handler.NewAuth().GenerateToken)

	err = httpServer.Listen(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
