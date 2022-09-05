package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/wianoski/api-vss/static"
	"github.com/wianoski/api-vss/other"
	
)


type Data struct {
	Data   struct {
		Token string `json:"token"`
		Pid   string `json:"pid"`
	} `json:"data"`
}



var outcome Data

func GetToken() (string, string){

	server := "vss"
	key := static.ActionTypes(1)
	action := "apiLogin.action?"

	url := other.GetEnvVariable("URL")
	urlPort := other.GetEnvVariable("URL_PORT")
	user_name := other.GetEnvVariable("USER_NAME")
	user_hash := other.GetEnvVariable("USER_HASH")


	requestURL := fmt.Sprintf("%s:%s/%s/%s/%susername=%s&password=%s", url,urlPort, server,key, action, user_name, user_hash)
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

	return Tokens, Pids
}