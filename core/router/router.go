package router

import (
    "github.com/gin-gonic/gin"
    "gin/app/http/middleware"
)

func SetRouter(ginEngine *gin.Engine) {
    ginEngine.Use(middleware.RequestIdMiddleware())
    setApiRouter(ginEngine)
    setWebRouter(ginEngine)
}
