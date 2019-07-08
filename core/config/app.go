package config

import (
    "github.com/joho/godotenv"
    "gin/library"
    "log"
    "os"
)

func (c *ConfENV) loadEnv() {
    env := os.Getenv("ENV")
    if env == "" {
        c.ENV = "local"
        env = ".env." + "local"
    } else {
        c.ENV = env
        env = ".env." + env
    }
    err := godotenv.Load(env)
    if err != nil {
        log.Println("Error loading " + env + " file")
    }
    thisEnv, _ = godotenv.Read(env)
}

func (c *ConfApp) loadAppConf() {
    c.Name = thisEnv["APP_NAME"]
    c.Debug = library.String2Bool(thisEnv["APP_DEBUG"])
    c.Domain = thisEnv["DB_PASSWORD"]
}
