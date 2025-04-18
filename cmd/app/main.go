package main

import (
	"github.com/joho/godotenv"
	"github.com/yusadeol/go-budgeteer/internal/infra/http"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/route"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpServer := http.NewChiServer()

	routerSetup := route.NewRouterSetup(httpServer)

	authRegistrar := route.NewAuthRegistrar()
	routerSetup.Register(authRegistrar)

	routerSetup.Apply()

	err = httpServer.Listen(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
