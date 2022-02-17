package main

import (
	"github.com/AyokunlePaul/Owo-Eyo/api/utils/logger"
	"github.com/AyokunlePaul/Owo-Eyo/blockchain/block"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var blockChain = block.Blockchain{}

func main() {
	route := gin.New()
	zapLogger := logger.GetLogger()
	route.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	route.Use(ginzap.RecoveryWithZap(zapLogger, true))

	blockChain.CreateBlock("", "Genesis")

	route.GET("/mine/:name", func(context *gin.Context) {
		data := context.Param("name")
		previousBlock := blockChain.GetPreviousBlock()
		newBlock := blockChain.CreateBlock(previousBlock.Hash, data)
		context.JSON(http.StatusOK, newBlock)
	})

	route.GET("/all", func(context *gin.Context) {
		context.JSON(http.StatusOK, blockChain.Chain)
	})

	logger.Error("application start error", route.Run(":8080"))
}
