package main

import (
    "github.com/gin-gonic/gin"
    "gin/core/router"
    "gin/core/template"
    "log"
)

func main() {
    /**
     * Set Flags
     * 設定 log flags
     */
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    /**
     * Gin Engine
     * 伺服器引擎
     */
    ginEngine := gin.Default()

    /**
     * Template
     * 模板
     */
    template.SetHTMLRender(ginEngine)

    /**
     * Router
     * 路由器
     */
    router.SetRouter(ginEngine)

    /**
     * Start Gin Engine
     * 伺服器引擎啟動
     */
    if err := ginEngine.Run(":3000"); err != nil {
        log.Fatalln(err)
    }
}
