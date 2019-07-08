package v1

import (
    "github.com/gin-gonic/gin"
    "gin/app/http/response"
    "gin/core/jwt"
    "net/http"
)

func Login(c *gin.Context) {
    // 新增 response struct
    responseStruct := response.GetDefaultStruct()

    // 從 post 拿值 轉 int
    user := c.PostForm("user")
    pwd := c.PostForm("pwd")
    if user == "" || pwd == "" {
        responseStruct.HttpCode = http.StatusBadRequest
        responseStruct.Err = "empty parameter"
    // 驗證帳號密碼 成功
    } else if verify := tempVerify(user, pwd); verify{
        jwtStruct := jwt.Generate(user)
        testJwtStruct := jwt.GetId(jwtStruct.Token)

        responseStruct.Data = testJwtStruct
    // 驗證帳號密碼 失敗
    } else{
        responseStruct.HttpCode = http.StatusUnauthorized
        responseStruct.Err = "verification failed"
    }
    // response
    c.JSON(http.StatusOK, responseStruct)
}
func tempVerify(user string, pwd string) bool{
    if user == "york" && pwd == "123456" {
        return true
    }
    return false
}