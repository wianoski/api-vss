package other

import (
	"os"
	"log"
  	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("key.env")
  
	if err != nil {
	  log.Fatalf("[err]Error loading .env file")
	}
  
	return os.Getenv(key)
  }