package jwt

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "strings"
    "time"
)

func Generate(id string) *jwt {
    // 新增 jwt token struct
    jwtStruct := new(jwt)
    // 設定 header payload
    jwtStruct.setHeader()
    jwtStruct.setPayload(id)
    // 建立 token
    jwtStruct.createToken()
    // 本地建立指定 verify
    jwtStruct.Status = true
    // return
    return jwtStruct
}

func GetId(jwtToken string) *jwt {
    // 新增 jwt token struct
    jwtStruct := new(jwt)
    // 設定 token
    jwtStruct.setToken(jwtToken)
    // 驗證
    if jwtStruct.verify(); jwtStruct.Status {
        // 建立 header payload
        jwtStruct.createHeader()
        jwtStruct.createPayload()
    }
    // return
    return jwtStruct
}

func (j *jwt) setHeader() {
    j.Header = header{
        Alg: "SHA256",
        Typ: "JWT",
    }
}

func (j *jwt) createHeader() {
    // 切割 token
    split := strings.Split(j.Token, ".")
    headerByte := structDecode(split[0])
    // 解析 header
    if err := json.Unmarshal(headerByte, &j.Header); err != nil {
        fmt.Println(`failed header Decode`, err)
    }
}

func (j *jwt) setPayload(id string) {
    j.Payload = payload{
        Id:        id,
        Timestamp: time.Now().Unix(),
    }
}

func (j *jwt) createPayload() {
    // 切割 token
    split := strings.Split(j.Token, ".")
    payloadByte := structDecode(split[1])
    // 解析 payload
    if err := json.Unmarshal(payloadByte, &j.Payload); err != nil {
        fmt.Println(`failed payload Decode`, err)
    }
}

func (j *jwt) setToken(t string) {
    j.Token = t
}

func (j *jwt) createToken() {
    // 轉json格式
    jsonTokenHeader, _ := json.Marshal(j.Header)
    jsonTokenPayload, _ := json.Marshal(j.Payload)
    // 加密
    unsignedToken := structEncode(jsonTokenHeader) + "." + structEncode(jsonTokenPayload)
    signature := encodeHS256(unsignedToken)
    // 組合
    j.Token = unsignedToken + "." + signature
}

func (j *jwt) verify() {
    j.Status = true
}

func structEncode(s []byte) string {
    return base64.RawURLEncoding.EncodeToString(s)
}

func structDecode(str string) []byte {
    b, err := base64.RawURLEncoding.DecodeString(str)
    if err != nil {
        fmt.Println(`failed base64 Decode`, err)
    }
    return b
}

func encodeHS256(u string) string {
    // new 一個 HMAC 結構體使用 jwt 的 key
    h := hmac.New(sha256.New, []byte(key))
    // 寫入資料
    h.Write([]byte(u))
    // Get result and encode as hexadecimal string
    return hex.EncodeToString(h.Sum(nil))
}

func decodeHS256(u string) string {
    return "TO DO"
}