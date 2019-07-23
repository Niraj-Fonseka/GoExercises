package main 


import (
	"fmt"
	"time"
)

func main(){

	startTime := "2019-06-21T13:56:00Z"
	startTimeInTime := ESTtoUTC(startTime , "2006-01-02T15:04:05Z")


	fmt.Printf("From EST  %s to UTC %v \n" , startTime , startTimeInTime)

	
}


func ESTtoUTC(timeStamp string , layout string) time.Time {
	TimeZone, _ := time.LoadLocation("America/New_York")
	newTimeZone, _ := time.ParseInLocation(layout, timeStamp, TimeZone)
	return newTimeZone.UTC()
}


func TimeStringToTimeObject(imageTime string, layout string) (interface{}, error) {
	if toTimeObj, err := time.Parse(layout, imageTime); err != nil {
		return nil, err
	} else {
		return toTimeObj, nil
	}
}


//ChangeTimes converts from EST to UTC and visa versa
func ChangeTimes(timeStamp string, from string, to string , layout string) string {
	if timeStamp == "" {
		return ""
	}
	TimeZone, _ := time.LoadLocation(from)
	if to == "UTC" {
		newTimeZone, _ := time.ParseInLocation(layout, timeStamp, TimeZone)
		return newTimeZone.UTC().Format(time.RFC3339)
	}
	location, _ := time.LoadLocation("America/New_York")
	newTime, _ := time.ParseInLocation(layout, timeStamp,location)
	return newTime.Format(time.RFC3339)

}