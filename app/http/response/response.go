package response

import (
    "net/http"
)

type Response struct {
    Status   bool
    HttpCode int64
    Data     interface{}
    Err      string
}

func GetDefaultStruct() Response {
    r := Response{
        Status:   true,
        HttpCode: http.StatusOK,
        Data:     "",
        Err:      ""}
    return r
}
