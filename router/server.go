package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunHTTPServer(port *string, mode *string) {
	r := routerEngine(mode)
	r.POST("/cutter", postCutter)
	r.Run(fmt.Sprintf(":%s", *port))
}

func routerEngine(mode *string) *gin.Engine {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *mode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	isTerm := isatty.IsTerminal(os.Stdout.Fd())
	if isTerm {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stdout,
				NoColor: false,
			},
		)
	}

	// set server mode
	gin.SetMode(*mode)

	r := gin.New()

	// Global middleware
	r.Use(logger.SetLogger(
		logger.WithUTC(true),
		logger.WithSkipPath([]string{
			"/ping",
		}),
	))
	r.Use(gin.Recovery())
	r.Use(VersionMiddleware())

	r.GET("/ping", heartbeatHandler)
	r.HEAD("/ping", heartbeatHandler)
	r.GET("/version", versionHandler)
	r.GET("/", rootHandler)

	return r
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome!",
	})
}

func heartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"source":  "https://github.com/enesdindas/video-cutter",
		"version": GetVersion(),
	})
}

func postCutter(c *gin.Context) {
	fmt.Println("postCutter")
	fmt.Println(c.Request.Body)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
