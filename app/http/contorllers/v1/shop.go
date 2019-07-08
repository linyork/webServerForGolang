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

func PostShopById(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    // 從 post 拿值 轉 int
    if id, idErr := strconv.Atoi(c.PostForm("id")); idErr != nil || id == 0 {
        responseStruct.HttpCode = http.StatusBadRequest
        responseStruct.Err = "Id not found"
    } else {
        //  new 一個 orm city 的 struct 裡頭設 ID 為 post 值
        shopData := make([]orm.ShopData, 0)
        responseStruct.Data = &shopData

        // 拿 DB 連線後操作 ORM 取得資料
        count, err := database.Mysql().
            Table(orm.TNShop).
            Join("INNER", orm.TNCategory, "shop.category_id = category.id").
            Where("shop.id = ? ", c.PostForm("id")).
            FindAndCount(&shopData)

        // DB 例外處理
        if err != nil {
            responseStruct.Status = false
            responseStruct.HttpCode = http.StatusInternalServerError
            log.Fatalln(err)
        }
        // 指定 Http Code
        if responseStruct.Status = count != 0; count == 0 {
            responseStruct.HttpCode = http.StatusNotFound
        }
    }
    // response
    c.JSON(http.StatusOK, responseStruct)
}
