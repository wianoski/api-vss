package services

type LoginWs struct {
	Action  string `json:"action"`
	Payload struct {
		Username string `json:"username"`
		Pid      string `json:"pid"`
		Token    string `json:"token"`
	} `json:"payload"`
}

type Heartbeat struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

type WS struct {
	Payload struct{
		DeviceID  string `json:"deviceID"`
		EventType string `json:"eventType"`
		Location  struct {
			Dtu        string `json:"dtu"`
			Direct     string `json:"direct"`
			Satellites string `json:"satellites"`
			Speed      string `json:"speed"`
			Altitude   string `json:"altitude"`
			Longitude  string `json:"longitude"`
			Latitude   string `json:"latitude"`
		} `json:"location"`
		Payload struct {
			Det struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Tp   string `json:"tp"`
			} `json:"det"`
			Dtu    string `json:"dtu"`
			Ec     string `json:"ec"`
			Et     string `json:"et"`
			St     string `json:"st"`
			UUID   string `json:"uuid"`
			NodeID string `json:"nodeID"`
		} `json:"payload"`
	} `json:"payload"`
}

type SendHB struct {
	Guid string `json:"guid"`
	Timestamp string `json:"timestamp"`
	Direction string `json:"direction"`
	Sats string `json:"sats"`
	Speed string `json:"speed"`
	Alt string `json:"alt"`
	Long string `json:"long"`
	Lat string `json:"lat"`
}