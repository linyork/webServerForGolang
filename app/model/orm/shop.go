package orm

var TNShop = "shop"

type Shop struct {
    Id             int
    MainCategory   string
    CategoryId     int
    SubCategory    string
    SubCategoryId  int
    ShopName       string
    Tel            string
    CityId         int
    DistrictId     int
    Address        string
    Metro          string
    LeaveDay       string
    Opening        string
    Description    string
    RateDelicious  string
    RateQuality    string
    RateEnviroment string
    Menu           string
    ShopImage      string
    UpdatedAt      string
}

type ShopData struct{
    Shop `xorm:"extends"`
    Category `xorm:"extends"`
}
