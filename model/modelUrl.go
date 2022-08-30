package model

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ServerUrl struct {
	Server string
	Key string
	Action string
}


var Url = "http://47.252.16.64"
var UrlPort = 9966

func SetServerApi(ServerUrl string, KeyUrl string, ActionUrl string, Token string, Paramater string) (string){
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", UrlPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		} 
	}()
	RequestUrl := fmt.Sprintf("%s:%d/%s/%s/%s?%s&%s",Url,UrlPort, ServerUrl,KeyUrl,ActionUrl,Token,Paramater)
	//						   "http://47.252.16.64:9966/vss/vehicle/getDeviceStatus.action?token=137e98bc3d0247ce80897b8b887bb029&deviceID=bb345"
	//							     Url:           urlport/serverurl/keyurl/actionurl
	res,err := http.Get(RequestUrl)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	return string(resBody)
}