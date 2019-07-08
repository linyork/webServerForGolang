package web

import (
    "github.com/gin-gonic/gin"
    "gin/app/model/orm"
    "gin/core/database"
    "gin/library"
    "log"
    "net/http"
)

func Index(c *gin.Context) {
    c.HTML(http.StatusOK, "index/index", gin.H{
        "title": "Index",
    })
}
func City(c *gin.Context) {

    //  new 一個 orm city 的 slice
    result := make([]orm.City, 0)
    data := &result
    // DB查詢
    count, err := database.Mysql().Table("city").FindAndCount(&result)

    // DB 例外處理
    if err != nil {
        log.Fatalln(err)
    }

    // response
    c.HTML(http.StatusOK, "city/list", gin.H{
        "title": "City",
        "data": data,
        "count": count,
    })
}

func Category(c *gin.Context) {
    result := make([]orm.Category, 0)
    data := &result
    count, err := database.Mysql().Table("category").FindAndCount(&result)

    if err != nil {
        log.Fatalln(err)
    }
    c.HTML(http.StatusOK, "category/list", gin.H{
        "title": "Category",
        "data": data,
        "count": count,
    })
}

func District(c *gin.Context) {
    result := make([]orm.District, 0)
    data := &result
    count, err := database.Mysql().Table("district").FindAndCount(&result)

    if err != nil {
        log.Fatalln(err)
    }
    c.HTML(http.StatusOK, "district/list", gin.H{
        "title": "District",
        "data": data,
        "count": count,
    })
}

func Shop(c *gin.Context) {
    // 常數
    maxPageShopCount := 10

    result := make([]orm.Shop, 0)
    data := &result

    // get data count
    count, err := database.Mysql().Table("shop").Count()
    if err != nil {
        log.Fatalln(err)
    }

    pageText := c.Param("page")
    page := library.GetPage(pageText, maxPageShopCount)
    
    getOffset := ( page - 1 ) * maxPageShopCount

    err = database.Mysql().Table("shop").Limit(maxPageShopCount, getOffset).Find(&result)
    if err != nil {
        log.Fatalln(err)
    }

    // pagination
    maxPagination := int(count) / maxPageShopCount
    if int(count) % maxPageShopCount > 0 {
        maxPagination = maxPagination + 1
    }
    
    pagination := library.GetPagination(page, maxPagination)

    c.HTML(http.StatusOK, "shop/list", gin.H{
        "title": "Shop",
        "data": data,
        "count": count,
        "pagination": pagination,
    })
}

func SubCategory(c *gin.Context) {
    result := make([]orm.SubCategory, 0)
    data := &result
    count, err := database.Mysql().Table("sub_category").FindAndCount(&result)

    if err != nil {
        log.Fatalln(err)
    }
    c.HTML(http.StatusOK, "sub_category/list", gin.H{
        "title": "SubCategory",
        "data": data,
        "count": count,
    })
}