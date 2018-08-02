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
		"http://slowwly.robertomurray.co.uk/delay/10000/url/http://www.mocky.io/v2/5b628ffb3000005a0064ffea",
		"http://slowwly.robertomurray.co.uk/delay/2000/url/ http://www.mocky.io/v2/5b629066300000620964ffef",
		"http://slowwly.robertomurray.co.uk/delay/2000/url/http://www.mocky.io/v2/5b629089300000710064fff1",
	}
	jsonResponses := make(chan Running)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for key, url := range urls {
		go func(url string) {
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

			list.countersLock.Unlock()
		}
	}()

	go func() {
		for index := 0; index < 3; index++ {
			fmt.Println(list.counters[index])
		}
	}()

	wg.Wait()
}
