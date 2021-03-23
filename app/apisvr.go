package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine
var APISvr *http.Server

func NewAPISvr(runmode string, httpport int, readTimeout time.Duration, writeTimeout time.Duration) {

	gin.SetMode(runmode)

	// GinEngine = gin.Default()
	GinEngine = gin.New()
	GinEngine.Use(gin.Logger())
	GinEngine.Use(gin.Recovery())

	endPoint := fmt.Sprintf(":%d", httpport)
	maxHeaderBytes := 1 << 20
	APISvr = &http.Server{
		Addr:           endPoint,
		Handler:        GinEngine,
		MaxHeaderBytes: maxHeaderBytes,

		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

}

func RunAPISvr() {
	go func() {
		err := APISvr.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("启动 http 服务失败，%v", err))
		}
	}()

}

func StopAPISvr() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	APISvr.Shutdown(ctx)
}

/*
	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless

	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes

	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}

	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
*/
