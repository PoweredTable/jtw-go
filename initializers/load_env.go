package initializers

import (
	"log"
	"os"
)

var DbDSN string
var JwtKey []byte

func LoadEnv() {
	DbDSN = os.Getenv("DB_DSN")
	if DbDSN == "" {
		log.Fatal("DB_DSN environment variable not set")
	}

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatal("JWT_KEY environment variable not set")
	}
	JwtKey = []byte(jwtKey)
}
