package v1

import (
    "github.com/gin-gonic/gin"
    "gin/app/http/response"
    "gin/app/model/orm"
    "gin/core/database"
    "log"
    "net/http"
)

func GetDistrictById(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    // 從 get 拿值
    id := c.Param("id")

    //  new 一個 orm District 的 struct
    responseStruct.Data = new(orm.District)

    // 拿 DB 連線後操作 ORM 取得資料
    xorm := database.Mysql()
    has, err := xorm.Id(id).Get(responseStruct.Data)

    // DB 例外處理
    if err != nil {
        responseStruct.HttpCode = http.StatusInternalServerError
        log.Fatalln(err)
    }

    // 指定 Http Code
    if responseStruct.Status = has; has {
        responseStruct.HttpCode = http.StatusOK
    } else {
        responseStruct.HttpCode = http.StatusNotFound
    }

    // response
    c.JSON(http.StatusOK, responseStruct)
}

func PostDistrictById(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    //  new 一個 orm District 的 struct
    responseStruct.Data = new(orm.District)

    // 從 post 拿值
    id := c.PostForm("id")
    if id == "" {
        responseStruct.HttpCode = http.StatusBadRequest
        responseStruct.Err = "Id not found"
    } else {
        // 拿 DB 連線後操作 ORM 取得資料
        xorm := database.Mysql()
        has, err := xorm.Id(id).Get(responseStruct.Data)

        // DB 例外處理
        if err != nil {
            responseStruct.HttpCode = http.StatusInternalServerError
            log.Fatalln(err)
        }

        // 指定 Http Code
        if responseStruct.Status = has; has {
            responseStruct.HttpCode = http.StatusOK
        } else {
            responseStruct.HttpCode = http.StatusNotFound
        }
    }

    // response
    c.JSON(http.StatusOK, responseStruct)
}
