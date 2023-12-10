package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"ns/video-cutter/router"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Version control for notify.
var version = "No Version Provided"

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
		mode        string
	)

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "V", false, "Print version information.")
	flag.StringVar(&port, "p", "8080", "port number for")
	flag.StringVar(&port, "port", "8080", "port number for")
	flag.StringVar(&mode, "m", "release", "Set server mode.")
	flag.StringVar(&mode, "mode", "release", "Set server mode.")
	flag.BoolVar(&ping, "ping", false, "ping server")

	flag.Usage = usage
	flag.Parse()
	router.SetVersion(version)

	// Show version and exit
	if showVersion {
		router.PrintVersion()
		os.Exit(0)
	}

	if ping {
		if err := pinger(&port); err != nil {
			fmt.Errorf("failed to ping server: %v", err)
		}
		fmt.Println("ping success")
		os.Exit(0)
	}

	router.RunHTTPServer(&port, &mode)
}

func usage() {
	fmt.Printf("%s\n", usageStr)
}

// handles pinging the endpoint and returns an error if the
// agent is in an unhealthy state.
func pinger(port *string) error {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"http://localhost:"+*port+"/ping",
		nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status code")
	}
	return nil
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
