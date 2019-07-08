package config

type ConfENV struct {
    ENV string
}

type ConfApp struct {
    Name string
    Debug bool
    Domain string
}

type ConfDB struct {
    DataBase string
    Host     string
    Password string
    User     string
    Port     string
}

type ConfTemplate struct {
    Path string
}