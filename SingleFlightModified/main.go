package main

import (
	"GoExercises/SingleFlightModified/singleflight"
	"fmt"
	"time"
)

type SingleFlightGroup struct {
	RequestGroup *singleflight.Group
}

func main() {
	fmt.Println("Starting SingleFlightModified application")
	var requestGroup singleflight.Group

	fmt.Println("Press Ctrl + C to terminate the program")
	fmt.Scanln()

}

// func AddToRequestGroup(key string) {
// 	value, err, _ := typ.RequestGroup.Do(key, func() (interface{}, error) {
// 		go func() {
// 			time.Sleep(60 * time.Second)
// 			log.Println("Deleting key ", key)
// 			typ.RequestGroup.Forget(key)
// 		}()

// 		return finalData, nil
// 	})

// }

func (r *SingleFlightGroup) RequestOne(seconds time.Duration) string {
	fmt.Println("Starting RequestOne function")
	defer DoneExecutingFunction("RequestOne Exit")
	time.Sleep(seconds)
	return "RequestOne -> Done"
}

func (r *SingleFlightGroup) RequestTwo(seconds time.Duration) string {
	fmt.Println("Starting RequestTwo function")
	defer DoneExecutingFunction("RequestTwo Exit")

	time.Sleep(seconds)
	return "RequestTwo -> Done"
}

func (r *SingleFlightGroup) RequestThree(seconds time.Duration) string {
	fmt.Println("Starting RequestTwo function")
	defer DoneExecutingFunction("RequestThree Exit")

	time.Sleep(seconds)
	return "RequestThree -> Done"
}

func DoneExecutingFunction(terminateMessage string) {
	fmt.Printf("Done Executing : %s \n", terminateMessage)
}
