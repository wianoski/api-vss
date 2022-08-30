package main

import (
	"encoding/json"
	"time"
	// "log"
	"fmt"

	"github.com/wianoski/api-vss/model"
	"github.com/wianoski/api-vss/service"
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
var Token, PID string = type_test.GetToken()

func main() {
	// fmt.Printf("Token: %s\nPID: %s \n", Token,PID)

	var getDevice string = model.SetServerApi("vss","vehicle","getDeviceStatus.action",Token, "deviceID=bb345")
	// fmt.Println(getDevice)
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