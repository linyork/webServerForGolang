package router

import (
    "github.com/gin-gonic/gin"
    "gin/app/http/contorllers/v1"
    "gin/app/http/middleware"
)

func setApiRouter(ginEngine *gin.Engine) {
    apiGroup := ginEngine.Group("/api")
    apiGroup.Use(middleware.Api())

    v1Group := apiGroup.Group("/v1")
    {
        // login reply
        v1Group.POST("/login", v1.Login)

        // city
        v1Group.GET("/citys", v1.GetCitys)
        v1Group.POST("/city", v1.PostCityById)

        // category
        v1Group.GET("/category/:id", v1.GetCategoryById)
        v1Group.POST("/category/:id", v1.PostCategoryById)

        // sub category
        v1Group.GET("/sub_category/:id", v1.GetSubCategoryById)
        v1Group.POST("/sub_category/:id", v1.PostSubCategoryById)

        // district
        v1Group.GET("/district/:id", v1.GetDistrictById)
        v1Group.POST("/district/:id", v1.PostDistrictById)

        // shop
        v1Group.POST("/shop", v1.PostShopById)


    }

    v1GroupVerify := apiGroup.Group("/v1")
    {
        v1GroupVerify.Use(middleware.VerifyToken())
        v1GroupVerify.GET("/city/:id", v1.GetCityById)
    }

}
