package main

import (
	"time"
	// "log"
	"fmt"

	"github.com/wianoski/api-vss/type_test"
)
var Token, PID string = type_test.GetToken()

func main() {
	fmt.Printf("Token: %s\nPID: %s \n", Token,PID)
	time.Sleep(100 * time.Millisecond)	
}