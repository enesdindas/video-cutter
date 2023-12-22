package router

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	ffmpeg_wrapper "ns/video-cutter/ffmpeg_wrapper"
	"os"
	"path/filepath"

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
	r.SetTrustedProxies(nil)
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
	request := c.Request
	file, fileHeader, err := request.FormFile("videoFile")
	if err != nil {
		log.Logger.Debug().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	jsonOutput, err := ffmpeg_wrapper.GetJsonOutput(file)
	// start := "00:00:00"
	// end := "00:00:10"
	// writer := bytes.NewBuffer(nil)
	// if cutVideoErr := ffmpeg_wrapper.CutVideo(file, writer, start, end); cutVideoErr != nil {
	// 	log.Logger.Debug().Msg(cutVideoErr.Error())
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": cutVideoErr.Error()})
	// 	return
	// }
	if err != nil {
		log.Logger.Debug().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Debug().Msg(jsonOutput)
	var buf bytes.Buffer
	io.Copy(&buf, file)
	fileName := fileHeader.Filename
	fileSize := fileHeader.Size
	fileExt := filepath.Ext(fileName)
	if fileExt != ".mp4" {
		log.Logger.Debug().Msg("File extension is not mp4")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File extension is not mp4"})
		return
	}
	log.Debug().Msg("File Name:" + fileHeader.Filename)
	c.JSON(http.StatusOK, gin.H{
		"filename": fileName,
		"filesize": fileSize,
	})
}
