package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	action := os.Args[1]
	commandIsRecognized := false

	allow := [2]string{"up", "down"}
	for _, a := range allow {
		if action == a {
			commandIsRecognized = true
		}
	}

	if commandIsRecognized == false {
		log.Fatal("Use up or down command")
	}

	dsn := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	cmd := exec.Command(
		"migrate",
		"-path", "db/migrations",
		"-database", dsn,
		action,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
