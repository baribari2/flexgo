package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PKMAIN  string
	PK2     string
	ADDMAIN string
	ADD2    string
	DTOKEN  string
	CHANNEL string
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

	PKMAIN = getEnv("PKMAIN")
	PK2 = getEnv("PK2")
	ADDMAIN = getEnv("ADDMAIN")
	ADD2 = getEnv("ADD2")
	DTOKEN = getEnv("DTOKEN")
	CHANNEL = getEnv("CHANNEL")
}
