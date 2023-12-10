package router

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

var version string

func SetVersion(ver string) {
	version = ver
}

func GetVersion() string {
	return version
}

func PrintVersion() {
	fmt.Printf(`%s, Compiler: %s %s, Copyright (C) 2023 Bo-Yi Wu, Inc.`,
		version,
		runtime.Compiler,
		runtime.Version())
	fmt.Println()
}

func VersionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-VERSION", version)
		c.Next()
	}
}
