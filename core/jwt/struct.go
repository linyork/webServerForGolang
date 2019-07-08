package jwt

var key = "CoryRightByYork!"
//var expireTime =

type jwt struct {
    Token   string  `json:"token"`
    Status  bool    `json:"status"`
    Header  header  `json:"header"`
    Payload payload `json:"payload"`
}

type header struct {
    Alg string `json:"alg"`
    Typ string `json:"type"`
}

type payload struct {
    Id        string `json:"id"`
    Timestamp int64  `json:"timestamp"`
}
