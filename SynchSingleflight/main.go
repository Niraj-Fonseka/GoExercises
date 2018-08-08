package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

//https://medium.com/@vCabbage/go-avoid-duplicate-requests-with-sync-singleflight-311601b3068b

func main() {
	var requestGroup singleflight.Group

	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			// Start a goroutine which deletes the "github" key after 250ms.
			go func() {
				time.Sleep(1 * time.Minute)
				log.Println("Deleting \"github\" key")
				requestGroup.Forget("github")
			}()

			return githubStatus()
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		status := v.(string)

		log.Printf("/github handler requst: status %q, shared result %t", status, shared)

		fmt.Fprintf(w, "GitHub Status: %q", status)
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

// githubStatus retrieves GitHub's API status
func githubStatus() (string, error) {
	log.Println("Making request to GitHub API")
	defer log.Println("Request to GitHub API Complete")

	log.Println("Making that db call")
	time.Sleep(1 * time.Second)

	resp, err := http.Get("https://status.github.com/api/status.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	r := struct{ Status string }{}

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Status, err
}
