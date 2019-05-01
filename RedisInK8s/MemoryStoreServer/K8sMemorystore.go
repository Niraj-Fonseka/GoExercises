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

func redisHSET() {
	fmt.Println("Running redis HGET")
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
			FirstName: "JD2",
			LastName:  "Doe",
			Major:     "MATH",
		},
		Rank: 1,
	}

	studentJD3 := &Student{
		Info: &StudentDetails{
			FirstName: "JD3",
			LastName:  "Doe",
			Major:     "SC",
		},
		Rank: 1,
	}

	studentJD4 := &Student{
		Info: &StudentDetails{
			FirstName: "JD4",
			LastName:  "Doe",
			Major:     "SC2",
		},
		Rank: 1,
	}

	bytes, err := json.Marshal(studentJD)

	_, err = conn.Do("HSET", studentJD.Info.Major, studentJD.Info.FirstName, string(bytes))
	if err != nil {
		return
	}

	bytes2, _ := json.Marshal(studentJD2)

	_, err = conn.Do("HSET", studentJD2.Info.Major, studentJD2.Info.FirstName, string(bytes2))
	if err != nil {
		return
	}

	bytes3, _ := json.Marshal(studentJD3)

	_, err = conn.Do("HSET", studentJD3.Info.Major, studentJD3.Info.FirstName, string(bytes3))
	if err != nil {
		return
	}

	bytes4, _ := json.Marshal(studentJD4)

	_, err = conn.Do("HSET", studentJD4.Info.Major, studentJD4.Info.FirstName, string(bytes4))
	if err != nil {
		return
	}

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

	redisHSET()
	// http.HandleFunc("/redis", incrementHandler)
	// log.Fatal(http.ListenAndServe(":7000", nil))
}
