package static

func ActionTypes(statusCode int32) (string){
	switch statusCode{
	case 1:
		return "user"
	case 2:
		return "vehicle"
	case 3:
		return "onoffline"
	case 4:
		return "track"
	case 5:
		return "alarm"
	case 6:
		return "record"
	case 7:
		return "vss"
	case 8:
		return "webdownrecord"
	case 9:
		return "apiPage"
	default:
		return ""
	}
}