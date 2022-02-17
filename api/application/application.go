package application

import (
	"github.com/AyokunlePaul/Owo-Eyo/api/utils/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

var owoEyoRouter *gin.Engine

func init() {
	zapLogger := logger.GetLogger()
	owoEyoRouter.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	owoEyoRouter.Use(ginzap.RecoveryWithZap(zapLogger, true))
}
