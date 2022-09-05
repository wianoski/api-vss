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



func SetServerApi(ServerUrl string, KeyUrl string, ActionUrl string, Token string, Paramater string) (string){
	url := other.GetEnvVariable("URL")
	urlPort := other.GetEnvVariable("URL_PORT")
	RequestUrl := fmt.Sprintf("%s:%s/%s/%s/%s?token=%s&%s",url,urlPort, ServerUrl,KeyUrl,ActionUrl,Token,Paramater)
	res,_ := http.Get(RequestUrl)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	return string(resBody)
}