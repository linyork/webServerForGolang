package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
)

func VerifyToken() gin.HandlerFunc {
    return func(c *gin.Context) {

        log.Println("Verify Token")
    }
}
