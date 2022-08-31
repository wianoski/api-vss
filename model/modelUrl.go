package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/wianoski/api-vss/other"

)

type ServerUrl struct {
	Server string
	Key string
	Action string
}


var UrlPort = 9966

func SetServerApi(ServerUrl string, KeyUrl string, ActionUrl string, Token string, Paramater string) (string){
	url := other.GetEnvVariable("URL")
	RequestUrl := fmt.Sprintf("%s:%d/%s/%s/%s?token=%s&%s",url,UrlPort, ServerUrl,KeyUrl,ActionUrl,Token,Paramater)
	res,_ := http.Get(RequestUrl)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	return string(resBody)
}