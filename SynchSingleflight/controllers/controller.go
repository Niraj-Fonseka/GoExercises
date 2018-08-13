package controllers

import (
	"GoExercises/SynchSingleflight/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type Controller struct {
	Group *singleflight.Group
}

// func Github(c *gin.Context) {
// 	var requestGroup singleflight.Group

// 	v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
// 		// Start a goroutine which deletes the "github" key after 250ms.
// 		go func() {
// 			time.Sleep(30 * time.Second)
// 			log.Println("Deleting \"github\" key")
// 			requestGroup.Forget("github")
// 		}()

// 		return models.GithubStatus()
// 	})

// 	fmt.Println(err)
// 	status := v.(string)

// 	log.Printf("/github handler requst: status %q, shared result %t", status, shared)

// 	c.JSON(200, v.(string))
// }

func (ctl Controller) GetHealth(c *gin.Context) {
	//	var requestGroup singleflight.Group

	v, err, shared := ctl.Group.Do("health", func() (interface{}, error) {
		// Start a goroutine which deletes the "github" key after 250ms.
		go func() {
			time.Sleep(30 * time.Second)
			log.Println("Deleting \"health\" key")
			ctl.Group.Forget("health")
		}()

		return models.Health()
	})

	fmt.Println(err)
	status := v.(string)

	log.Printf("/health handler requst: status %q, shared result %t", status, shared)

	c.JSON(200, v.(string))
}
