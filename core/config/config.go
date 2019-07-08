package config

import (
    "fmt"
)

var (
    thisEnv  map[string]string
    ENV      *ConfENV
    DB       *ConfDB
    Template *ConfTemplate
)

func init() {

    ENV = &ConfENV{}
    ENV.loadEnv()

    DB = &ConfDB{}
    DB.loadDBConf()

    Template = &ConfTemplate{}
    Template.Path = defaultViewPath

    fmt.Print(DB)
}
