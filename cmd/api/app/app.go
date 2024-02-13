package app

import (
	"fmt"
	"time"

	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/agustinrabini/api-items-project/cmd/api/dependencies"

	"github.com/agustinrabini/go-toolkit/gingonic/handlers"
	"github.com/agustinrabini/go-toolkit/goutils/logger"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	handler, errBuildDepend := dependencies.BuildDependencies()
	if errBuildDepend != nil {
		fmt.Printf("Error Build Dependencies")
		waitAndPanic(errBuildDepend)
	}
	router := ConfigureRouter()
	MapUrlsToControllers(router, handler)

	if errRouter := router.Run(config.ConfMap.APIRestServerPort); errRouter != nil {
		logger.Errorf("Error starting router", errRouter)
		waitAndPanic(errRouter)
	}
}

func ConfigureRouter() *gin.Engine {
	logger.InitLog(config.ConfMap.LoggingPath, config.ConfMap.LoggingFile, config.ConfMap.LoggingLevel)
	return handlers.DefaultJopitRouter()

}

func waitAndPanic(err error) {
	time.Sleep(2 * time.Second) // needs one second to send the log
	panic(err)
}
