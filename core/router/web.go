package router

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "gin/app/http/contorllers/web"
    "gin/core/template"
)

func setWebRouter(ginEngine *gin.Engine) *gin.Engine {
    store := cookie.NewStore([]byte("secret"))
    //常數
    templateDir := "template/admin"
    ginEngine.HTMLRender = template.LoadTemplate(templateDir)

    ginEngine.Use(sessions.Sessions("session", store))
    ginEngine.GET("/sessioncount/*name", web.SessionCountAndName)
    ginEngine.GET("/", web.Test)

    adminGroup := ginEngine.Group("/admin")
    {
        adminGroup.GET("/", web.Index)
        adminGroup.GET("/city/", web.City)
        adminGroup.GET("/category/", web.Category)
        adminGroup.GET("/district/", web.District)
        adminGroup.GET("/shop/", web.Shop)
        adminGroup.GET("/shop/:page", web.Shop)
        adminGroup.GET("/sub_category/", web.SubCategory)
    }

    return ginEngine
}