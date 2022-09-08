package services

import (
	// "bytes"
	// "net/http"
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wianoski/api-vss/other"
)


var addr = flag.String("addr", other.GetEnvVariable("URL_WS"), "heartbeat_device")
var Token, PID string = GetToken()

// var loc Location

func SendHeartbeat() {
	// tbServer := other.GetEnvVariable("URL_TB")

	Login := &LoginWs{
		Action: "80000",
		Payload: struct{
			Username string `json:"username"`
			Pid string `json:"pid"`
			Token string `json:"token"`}{
				Username: other.GetEnvVariable("USER_NAME"),
				Pid: PID,
				Token: Token,
			},
	}

	hb := &Heartbeat{
		Action: "80001",
		Payload: "",
	}
	flag.Parse()
	log.SetFlags(0)

	jsonLogin, _ := json.Marshal(Login)
	PayloadLogin := string(jsonLogin)

	jsonPayload, _ := json.Marshal(hb)
	PayloadHeartbeat := string(jsonPayload)

	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt,os.Interrupt)

	target := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("Connecting to: %s\n", target.String())

	c, _, err := websocket.DefaultDialer.Dial(target.String(),nil)
	if err != nil{
		log.Fatal("dial", err)
	}
	
	connLogin := c.WriteMessage(websocket.TextMessage, []byte(PayloadLogin))
	if connLogin != nil {
		log.Println("write:", connLogin)
		return
	}
	connHeartbeat := c.WriteMessage(websocket.TextMessage, []byte(PayloadHeartbeat))
	if connHeartbeat != nil {
		log.Println("write:", connHeartbeat)
		return
	}


	defer c.Close()

	done := make(chan struct{})

	go func() {
		var ws WS
		defer close(done)
		for {
			err := c.ReadJSON(&ws)
			if err != nil {
				log.Println("read:", err)
				return
			}
			SendHeartbeat := &SendHB{
				Guid: ws.Payload.DeviceID,
				Timestamp: ws.Payload.Location.Dtu,
				Direction: ws.Payload.Location.Direct,
				Sats: ws.Payload.Location.Satellites,
				Speed: ws.Payload.Location.Speed,
				Alt: ws.Payload.Location.Altitude,
				Long: ws.Payload.Location.Longitude,
				Lat: ws.Payload.Location.Latitude,
			}

			// uncomment if want to use thingsboard

			// if loc.Payload.DeviceID == "bb345" {
			// 	buf := new(bytes.Buffer)
			// 	json.NewEncoder(buf).Encode(SendHeartbeat)
			// 	requestPost, _ := http.NewRequest("POST", tbServer, buf)
			// 	client := &http.Client{}
			// 	res, e := client.Do(requestPost)
			// 	if e != nil {
			// 		log.Println("error")
			// 	}
			// 	defer res.Body.Close()
			// 	log.Println("heartbeat response Status:", res.Status)
			// }
			
			packet, _ := json.Marshal(SendHeartbeat)
			Data := string(packet)
			log.Printf("recv: %s, alarm: %s", Data, ws.Payload.Payload)
			
			
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
} 