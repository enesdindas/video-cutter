package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Version control for notify.
var (
	version = "No Version Provided"
	commit  = "No Commit Provided"
)

var usageStr = `
Usage: video-cutter [options]

Server Options:
    -p, --port <port>                Use port for clients (default: 8088)
    --ping                           healthy check command for container
    -h, --help                       Show this message
    -V, --version                    Show version
`

func main() {
	var (
		ping        bool
		port        string
		showVersion bool
	)

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "V", false, "Print version information.")
	flag.StringVar(&port, "p", "", "port number for gorush")
	flag.StringVar(&port, "port", "", "port number for gorush")
	flag.BoolVar(&ping, "ping", false, "ping server")

	flag.Usage = usage
	flag.Parse()
	SetupRouter().Run(":8080")
}

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/cutter", postCutter)

	return router
}

func postCutter(c *gin.Context) {
	fmt.Println("postCutter")
	fmt.Println(c.Request.Body)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
