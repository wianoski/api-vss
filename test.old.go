package test

import (

    // Import godotenv

	"os"
	"log"
	"fmt"
  	"github.com/joho/godotenv"
)


// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
    // os package

  // godotenv package
  dotenv := goDotEnvVariable("KUNCI")

  fmt.Printf("godotenv : %s = %s \n", "KUNCI", dotenv)
}