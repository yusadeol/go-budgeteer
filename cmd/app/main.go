package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yusadeol/go-budgeteer/internal/app"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenGenerator := adapter.NewJWTGenerator()

	input := app.NewGenerateTokenInput(os.Getenv("JWT_KEY"), "Yuri Oliveira")

	output, err := app.NewGenerateToken(tokenGenerator).Execute(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output.Token)
}
