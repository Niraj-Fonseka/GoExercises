package models

import (
	"log"
	"time"
)

// githubStatus retrieves GitHub's API status
func GithubStatus() (string, error) {
	log.Println("Making request to GitHub API")
	defer log.Println("Request to GitHub API Complete")

	log.Println("Making that db call")
	time.Sleep(20 * time.Second)

	// resp, err := http.Get("https://status.github.com/api/status.json")
	// if err != nil {
	// 	return "", err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return "", fmt.Errorf("github response: %s", resp.Status)
	// }

	// r := struct{ Status string }{}

	// err = json.NewDecoder(resp.Body).Decode(&r)

	return "Stats_is_good", nil
}

func Health() (string, error) {
	log.Println("Making request to Health API")
	defer log.Println("Request to Health API Complete")

	log.Println("Making that db call healt")
	time.Sleep(10 * time.Second)

	return "Health is ok", nil
}
