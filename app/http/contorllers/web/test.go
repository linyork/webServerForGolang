package web

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

func Test(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"count": 123123})
}

func SessionCountAndName(c *gin.Context) {
    session := sessions.Default(c)
    var count int
    v := session.Get("count")
    if v == nil {
        count = 0
    } else {
        count = v.(int)
        count++
    }
    session.Set("count", count)

    name := session.Get("name")
    getName := c.Param("name")

    if name == nil && getName == "/" {
        name = "yorkDefault"
        session.Set("name", name)
    } else if name == nil {
        name = strings.Replace(getName, "/", "", -1)
        session.Set("name", name)
    }

     err := session.Save()
     if err != nil {
         c.JSON(http.StatusOK, "session error")
     }
    //c.JSON(http.StatusOK, gin.H{"count": count})
    c.JSON(http.StatusOK, gin.H{"count": count, "name": name})
}