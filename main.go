package main

import (
	"allcaps/api/router"
	"allcaps/pkg/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// cs := internals.InitCronJob()

	r := gin.Default()

	c := helpers.Client
	router.Handlerouter(r, c)
	err := r.Run(":9000")
	if err != nil {
		fmt.Println("Error occured", err)
	}
	// defer cs.Stop()
	// select {}
}
