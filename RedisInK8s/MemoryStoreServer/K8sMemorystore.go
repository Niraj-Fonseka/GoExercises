package main

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

type Student struct {
	Info *StudentDetails `json:”info,omitempty”`
	Rank int             `json:”rank,omitempty”`
}
type StudentDetails struct {
	FirstName string
	LastName  string
	Major     string
}

func redisTestHandler() {
	fmt.Println("Running redis test handler")
	conn := redisPool.Get()
	defer conn.Close()

	studentJD := &Student{
		Info: &StudentDetails{
			FirstName: "John",
			LastName:  "Doe",
			Major:     "CSE",
		},
		Rank: 1,
	}

	studentJD2 := &Student{
		Info: &StudentDetails{
			FirstName: "Niraj",
			LastName:  "Fonseka",
			Major:     "Non-of-yo-buss",
		},
		Rank: 1,
	}

	bytes, err := json.Marshal(studentJD)

	_, err = conn.Do("SADD", "instance", string(bytes))
	if err != nil {
		return
	}

	bytes2, _ := json.Marshal(studentJD2)

	_, err = conn.Do("SADD", "instance", string(bytes2))
	if err != nil {
		return
	}

	value, err := redis.Values(conn.Do("SMEMBERS", "instance"))

	for _, val := range value {
		var v, ok = val.([]byte)
		if ok {
			fmt.Println(string(v))
		}
	}

	ClearCompleteCache()
	// time.Sleep(10 * time.Second)
	// valuesdelete, err := redis.Values(conn.Do("SMEMBERS", "instance"))

	// //delete one thing
	// val := valuesdelete[0]
	// var v, ok = val.([]byte)
	// if ok {
	// 	fmt.Println("Value to delete :--")
	// 	deleteStudent := Student{}
	// 	err := json.Unmarshal(v, &deleteStudent)
	// 	if err != nil {
	// 		fmt.Println("Error when unmarshall : ", err)
	// 	}
	// 	fmt.Println("Saved : ", deleteStudent)
	// 	_, err = conn.Do("SREM", "instance", string(v))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// }

	// afterDelete, err := redis.Values(conn.Do("SMEMBERS", "instance"))

	// for _, val := range afterDelete {
	// 	var v, ok = val.([]byte)
	// 	if ok {
	// 		fmt.Println("Value to delete :--")
	// 		fmt.Println(string(v))
	// 		deleteStudent := Student{}
	// 		err := json.Unmarshal(v, deleteStudent)
	// 		fmt.Println("Saved : ", deleteStudent)
	// 		_, err = conn.Do("SREM", "instance", string(v))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 	}
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}

func ClearCompleteCache() error {

	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")

	if err != nil {
		return err
	}

	return nil
}

func main() {

	fmt.Println("Starting up the program")
	redisHost := "0.0.0.0"
	redisPort := "6379"
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	const maxConnections = 10
	redisPool = redis.NewPool(func() (redis.Conn, error) {

		return redis.Dial("tcp", redisAddr)
	}, maxConnections)

	redisTestHandler()
	// http.HandleFunc("/redis", incrementHandler)
	// log.Fatal(http.ListenAndServe(":7000", nil))
}
