# gin webserver 本機開發環境建置

請將該專案下載到 GOPATH 的 src 之下

## 設定 host

```bash
# gin
127.0.0.1   local.gin.tw
```

## GoPath GoRoot 相關 (本機操作)

GoRoot 為 Go 編譯元件安裝 Go 之後預設路徑為

* `GOROOT="/usr/local/Cellar/go/{version}/libexec"`

GoPath 放哪裡都可以預設為

* `GOPATH="/Users/{userName}/go"`

GoBin 為 GoPath 底下的 bin 資料夾

* `GOPATH="/Users/{userName}/go/bin"`

Go 本身以package為單位從github下載來的插件預計會在

* `/Users/{userName}/go/src/{packageName}`

* 案件本身也應該在該路徑底下`/Users/{userName}/go/src/{projectName}`

如使用 iterm2 可使用以下指令 或將 該指令放入預載入的檔案內

```bash
export GOPATH="${HOME}/go"
export GOBIN="${GOPATH}/bin"
export PATH="${PATH}:${GOBIN}"
```

如使用 windows 系統 請在環境變數設定以上相關變數

## go get 相關 (本機操作)

首先先移駕至 gin 轉案目錄下

 * `/Users/{userName}/go/src/gin`

下載 go 的套件管理器 govendor 

 * `go get github.com/kardianos/govendor`

再來下載不需重新編譯的小插件

 * `go get github.com/codegangsta/gin`

執行下載 vendor.json 裡的 套件

 * `govendor sync`

## Core Docker 相關 (Docker)

更新至最新版本 core-docker 會有 gin 的專案設定直接啟動即可

`ps: core-docker 該 project 中的 .env 檔, 記得將該專案的路徑貼過去`