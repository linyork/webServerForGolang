package v1

import (
    "github.com/gin-gonic/gin"
    "gin/app/http/response"
    "gin/app/model/orm"
    "gin/core/database"
    "log"
    "net/http"
    "strconv"
)

func GetCityById(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    // 從 get 拿值
    id := c.Param("id")

    //  new 一個 orm City 的 struct
    responseStruct.Data = new(orm.City)

    // 拿 DB 連線後操作 ORM 取得資料
    has, err := database.Mysql().ID(id).Get(responseStruct.Data)

    // DB 例外處理
    if err != nil {
        responseStruct.HttpCode = http.StatusInternalServerError
        log.Fatalln(err)
    }

    // 指定 Http Code
    if responseStruct.Status = has; !has {
        responseStruct.HttpCode = http.StatusNotFound
    }

    // response
    c.JSON(http.StatusOK, responseStruct)
}

func PostCityById(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    // 從 post 拿值 轉 int
    if id, idErr := strconv.Atoi(c.PostForm("id")); idErr != nil || id == 0 {
        responseStruct.HttpCode = http.StatusBadRequest
        responseStruct.Err = "Id not found"
    } else {
        //  new 一個 orm city 的 struct 裡頭設 ID 為 post 值
        responseStruct.Data = &orm.City{Id: id}

        // 拿 DB 連線後操作 ORM 取得資料
        has, err := database.Mysql().Get(responseStruct.Data)

        // DB 例外處理
        if err != nil {
            responseStruct.HttpCode = http.StatusInternalServerError
            log.Fatalln(err)
        }

        // 指定 Http Code
        if responseStruct.Status = has; !has {
            responseStruct.HttpCode = http.StatusNotFound
        }
    }
    // response
    c.JSON(http.StatusOK, responseStruct)
}

func GetCitys(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    //  new 一個 orm city 的 slice
    result := make([]orm.City, 0)
    responseStruct.Data = &result

    // DB查詢
    count, err := database.Mysql().Table("city").FindAndCount(&result)

    // DB 例外處理
    if err != nil {
        responseStruct.Status = false
        responseStruct.HttpCode = http.StatusInternalServerError
        log.Fatalln(err)
    }
    // 指定 Http Code
    if responseStruct.Status = count != 0; count == 0{
        responseStruct.HttpCode = http.StatusNotFound
    }

    // response
    c.JSON(http.StatusOK, responseStruct)
}
