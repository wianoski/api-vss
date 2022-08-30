package main

import (
	"encoding/json"
	"time"
	// "log"
	"fmt"

	"github.com/wianoski/api-vss/model"
	"github.com/wianoski/api-vss/services"
)

type Device struct{
	Device []struct{
		DeviceGuid string `json:"deviceguid"`
		DeviceName string `json:"deviceName"`
		Long float64 `json:"longitude"`
		Lat float64 `json:"latitude"`
		Accstate interface{} `json:"accstate"`
	}`json:"data"`
}

var deviceData Device
var Token, PID string = services.GetToken()

func main() {

	server := "vss"
	key := "vehicle"
	action := "getDeviceStatus.action"
	param := "deviceID=bb345"

	var getDevice string = model.SetServerApi(server,key,action,Token, param)
	fmt.Printf("Action: %s\n", action)
	er := json.Unmarshal([]byte(getDevice), &deviceData)
	if er != nil{
		fmt.Println(er)
	}

	DeviceN := deviceData.Device
	for _,dd := range DeviceN{
		fmt.Printf("DeviceGuid: %s\nDevice Name: %s\nLatitude:%f\nLongitude: %f\n",dd.DeviceGuid,dd.DeviceName, dd.Lat, dd.Long)
		if dd.Accstate == nil {
			fmt.Println("Acc Status: Off")
		}else{
			fmt.Printf("Acc Status: %s(ON)\n", dd.Accstate)
		}
	}
	
	time.Sleep(100 * time.Millisecond)	
}