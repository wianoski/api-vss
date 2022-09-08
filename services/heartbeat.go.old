package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	// "log"
	"fmt"

	"github.com/wianoski/api-vss/model"
	"github.com/wianoski/api-vss/other"
	"github.com/wianoski/api-vss/static"
)

type Device struct{
	Device []struct{
		DeviceGuid string `json:"deviceguid"`
		DeviceName string `json:"deviceName"`
		Long float64 `json:"longitude"`
		Lat float64 `json:"latitude"`
		Sat string `json:"satellites"`
		Speed string `json:"speed"`
		Direct string `json:"direct"`
		Accstate interface{} `json:"accstate"`
	}`json:"data"`
}

type Payload struct{
	DeviceGuid string `json:"guid"`
	DeviceName string `json"deviceName"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json":"latitude"`
	Satelitte string `json:"satellites"`
	Speed string `json:"speed"`
	Direct string `json:"direct"`
	AccState interface{} `json"accstate"`
}


var deviceData Device
var Token, PID string = GetToken()

func HeartBeat(){
		server := "vss"
		key := static.ActionTypes(2)
		action := "getDeviceStatus.action"
		param := "deviceID=bb345"

		tbServer := other.GetEnvVariable("URL_TB")

		var getDevice string = model.SetServerApi(server,key,action,Token, param)
		fmt.Printf("Action: %s\n", action)
		er := json.Unmarshal([]byte(getDevice), &deviceData)
		if er != nil{
			fmt.Println(er)
		}

		DeviceN := deviceData.Device
		for _,dd := range DeviceN{
			fmt.Printf("DeviceGuid: %s\nDevice Name: %s\n",dd.DeviceGuid,dd.DeviceName)
			if dd.Accstate == nil {
				fmt.Println("Acc Status: Off")
			}else{
				fmt.Printf("Acc Status: On\n")
			}
			payloadSend := &Payload{
				DeviceGuid: dd.DeviceGuid,
				DeviceName: dd.DeviceName,
				Longitude: dd.Long,
				Latitude: dd.Lat,
				Satelitte: dd.Sat,
				Speed: dd.Speed,
				Direct: dd.Direct,
				AccState: dd.Accstate,
			}
			buf := new(bytes.Buffer)
			json.NewEncoder(buf).Encode(payloadSend)
			requestPost, _ := http.NewRequest("POST", tbServer, buf)
			client := &http.Client{}
			res, e := client.Do(requestPost)
			if e != nil {
				fmt.Println("error")
			}
		
			defer res.Body.Close()
		
			fmt.Println("heartbeat response Status:", res.Status)
			
		}
		
		time.Sleep(100 * time.Millisecond)	

}