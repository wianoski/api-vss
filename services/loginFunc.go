package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	
)

type Server struct{
	Server string
}

type API struct{
	Key string
}

type Action struct {
	Action string
}

type Data struct {
	Data   struct {
		Token string `json:"token"`
		Pid   string `json:"pid"`
	} `json:"data"`
}

const url = "http://47.252.16.64"
const urlPort = 9966
const user_name = "BlueB"
const user_hash = "3d22abdc5e8c9ab21ba13b540a8875f0"
var outcome Data
var NameServer = Server{"vss"}
var ApiUser = API{"user"}
var ActionLogin = Action{"apiLogin.action?"}

func GetToken() (string, string){
	requestURL := fmt.Sprintf("%s:%d/%s/%s/%susername=%s&password=%s", url,urlPort, NameServer.Server,ApiUser.Key, ActionLogin.Action, user_name, user_hash)
	res, err := http.Get(requestURL)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	er := json.Unmarshal([]byte(resBody), &outcome)
	
	if er != nil{
		fmt.Println(er)
	}

	Tokens := outcome.Data.Token
	Pids := outcome.Data.Pid
	
	// fmt.Printf("Token: %+v\n", outcome.Data.Token)
	// fmt.Printf("PID: %+v\n", outcome.Data.Pid)
	return Tokens, Pids
}