package library

import (
    "log"
    "strconv"
    "strings"
)

// 取得分頁數字
func GetPage(pageText string, maxCount int) int {

    page := 1
    if pageText != "" {
        pageString := strings.Replace(pageText, "page", "", -1)
        pageInt, err := strconv.Atoi(pageString)
        if err != nil {
            log.Fatalln(err)
        }
        page = pageInt
    }

    return page
}

// 參考chocolat分頁
func GetPagination(page int, maxPagination int) map[string]interface{} {
    // 常數
    paginationCount := 5

    pageData := make(map[string]interface{})
    pageData["first"] = 1
    pageData["last"] = maxPagination
    pageData["prev"] = 0
    pageData["next"] = 0
    pages := make([]int, 0)

    for i := 0; i < paginationCount; i++ {
        if page + i <= maxPagination && page + i >= 1 {
            pages = append(pages, page + i)
        }
    }
    // prepend
    for i := 1; len(pages) < paginationCount ; i++ {
        pages = append([]int{page - i}, pages...)
    }
    pageData["pages"] = pages

    if page == 1 {
        pageData["next"] = page + 1
    } else if page == maxPagination {
        pageData["prev"] = page - 1
        pageData["next"] = 0
    } else {
        pageData["prev"] = page - 1
        pageData["next"] = page + 1
    }

    return pageData
}