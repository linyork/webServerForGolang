package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/satori/go.uuid"
    "log"
)

func RequestIdMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("X-Request-Id", uuid.Must(uuid.NewV4()).String())
        c.Next()
    }
}

func Api() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Request in")
        c.Next()
        log.Println("Response out")
    }
}
