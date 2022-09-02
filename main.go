package main

import (
	"time"
	"github.com/wianoski/api-vss/services"
)


func main() {
	for range time.Tick(time.Second * 10) {
		services.HeartBeat()
	}
}