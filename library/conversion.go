package library

import (
    "log"
)

func String2Bool(str string) bool{
    switch str {
    case "1", "t", "T", "true", "TRUE", "True":
        return true
    case "0", "f", "F", "false", "FALSE", "False":
        return false
    default:
        log.Fatalln("Conversion Fall:"+str)
        return false
    }
}
