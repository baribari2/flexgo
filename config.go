package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PK string
	ME string
)

func getEnv(env string) string {
	env, exists := os.LookupEnv(env)
	if !exists {
		log.Fatalf("Env var %v doesn't exist: ", env)
	}

	return env
}

func init() {
	godotenv.Load()

	PK = getEnv("PK")
	ME = getEnv("ME")
}
