package main

import (
	"gorm.io/gorm"
	"task-sesi8/configs"
	_ "task-sesi8/docs"
	"task-sesi8/routers"
)

var (
	db *gorm.DB = configs.ConnectDB()
)

// @title Orders API
// @description Sample API Spec for Orders
// @version v1.0
// @termsOfService /
// @BasePath /
// @host localhost:8080
// @contact.name M Hilmi Muharromi
// @contact.email hilmimuharrom@gmail.com
func main() {
	configs.MigrateDatabase(db)
	configs.DisconnectDB(db)
	routers.SetupRouter()
}
