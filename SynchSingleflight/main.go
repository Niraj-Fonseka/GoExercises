package main

import (
	"GoExercises/SynchSingleflight/controllers"
	"GoExercises/SynchSingleflight/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

//https://medium.com/@vCabbage/go-avoid-duplicate-requests-with-sync-singleflight-311601b3068b

func main() {

	// r := gin.Default()
	// r.GET("/github", controllers.Github)
	// r.Run() // listen and serve on 0.0.0.0:7070

	var requestGroup singleflight.Group

	r := gin.Default()
	r.GET("/github", func(c *gin.Context) {
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			// Start a goroutine which deletes the "github" key after 250ms.
			go func() {
				time.Sleep(60 * time.Second)
				log.Println("Deleting \"github\" key")
				requestGroup.Forget("github")
			}()

			return models.GithubStatus()
		})

		fmt.Println(err)
		status := v.(string)

		log.Printf("/github handler requst: status %q, shared result %t", status, shared)

		c.JSON(200, v.(string))
	})

	r.GET("/health", controllers.Controller{Group: &requestGroup}.GetHealth)

	r.Run() // listen and serve on 0.0.0.0:7070

	// http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
	// 	v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
	// 		// Start a goroutine which deletes the "github" key after 250ms.
	// 		go func() {
	// 			time.Sleep(1 * time.Minute)
	// 			log.Println("Deleting \"github\" key")
	// 			requestGroup.Forget("github")
	// 		}()

	// 		return githubStatus()
	// 	})
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	status := v.(string)

	// 	log.Printf("/github handler requst: status %q, shared result %t", status, shared)

	// 	fmt.Fprintf(w, "GitHub Status: %q", status)
	// })

	//http.ListenAndServe("127.0.0.1:7070", nil)
}

// githubStatus retrieves GitHub's API status
// func githubStatus() (string, error) {
// 	log.Println("Making request to GitHub API")
// 	defer log.Println("Request to GitHub API Complete")

// 	log.Println("Making that db call")
// 	time.Sleep(60 * time.Second)

// 	// resp, err := http.Get("https://status.github.com/api/status.json")
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// 	// defer resp.Body.Close()

// 	// if resp.StatusCode != 200 {
// 	// 	return "", fmt.Errorf("github response: %s", resp.Status)
// 	// }

// 	// r := struct{ Status string }{}

// 	// err = json.NewDecoder(resp.Body).Decode(&r)

// 	return "Stats_is_good", nil
// }
