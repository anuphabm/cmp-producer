package database

import (
	"cmpdata/config"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	var err error
	envMode := os.Getenv("RUN_MODE")
	user := config.Appconfig.GetString(fmt.Sprintf("%s.database.username", envMode))
	pwd := config.Appconfig.GetString(fmt.Sprintf("%s.database.password", envMode))
	host := config.Appconfig.GetString(fmt.Sprintf("%s.database.host", envMode))
	port := config.Appconfig.GetInt(fmt.Sprintf("%s.database.port", envMode))
	name := config.Appconfig.GetString(fmt.Sprintf("%s.database.name", envMode))

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, pwd)
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	logMode := config.Appconfig.GetBool(fmt.Sprintf("%s.database.logmode", envMode))
	DB.LogMode(logMode)
}
