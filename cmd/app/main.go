package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/route"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpServer := adapter.NewChiServer()

	routerSetup := route.NewRouterSetup(httpServer)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	dbConnection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	userRegistrar := route.NewUserRegistrar(dbConnection)
	routerSetup.Register(userRegistrar)

	routerSetup.Apply()

	err = httpServer.Listen(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
