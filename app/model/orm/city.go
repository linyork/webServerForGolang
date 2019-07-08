package orm

var TNCity = "city"

type City struct {
    Id       int `xorm:"pk"`
    CityName string
}
