package config

func (c *ConfDB) loadDBConf() {
    c.DataBase = thisEnv["DB_DATABASE"]
    c.Host = thisEnv["DB_HOST"]
    c.Password = thisEnv["DB_PASSWORD"]
    c.User = thisEnv["DB_USERNAME"]
    c.Port = thisEnv["DB_PORT"]
}

