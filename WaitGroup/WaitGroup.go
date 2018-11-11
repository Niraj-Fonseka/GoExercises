package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	fmt.Println("Running the waitgroup application")
	go RunGoWaitGroup()
	//go RunGoWaitGroup()

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("waiting for user input")
	input.Scan()
}

type Running struct {
	Order    int
	Response string
}

type MapForMutlipleReturns struct {
	counters     map[int]string
	countersLock sync.RWMutex
}

func RunGoWaitGroup() {
	urls := []string{
		"http://slowwly.robertomurray.co.uk/delay/5000/url/http://www.mocky.io/v2/5b629f8d300000490065002b",
		"http://slowwly.robertomurray.co.uk/delay/2000/url/http://www.mocky.io/v2/5b629fc4300000620065002c",
		"http://slowwly.robertomurray.co.uk/delay/500/url/http://www.mocky.io/v2/5b62a00d300000620065002f",
	}
	jsonResponses := make(chan Running)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for key, url := range urls {
		go func(url string) {
			fmt.Println("Routine for : ", url)
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			} else {
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				} else {
					run := Running{
						Order:    key,
						Response: string(body),
					}
					jsonResponses <- run
				}
			}
		}(url)
	}

	var list MapForMutlipleReturns

	list.counters = make(map[int]string)
	go func() {
		fmt.Println("Running the fetching responses")
		index := 0
		for response := range jsonResponses {
			list.countersLock.Lock()
			fmt.Println(index)
			fmt.Println(response.Response)
			list.counters[index] = response.Response
			index++
			list.countersLock.Unlock()

		}
	}()

	go func() {
		fmt.Println("Keep checking struture")
		for index := 0; index < 3; index++ {
			fmt.Println(list.counters[index])
		}
	}()

	wg.Wait()
}
