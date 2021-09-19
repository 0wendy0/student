package orm

import (
	"log"
	"student/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitDatabase() {
	driver := config.Database.Driver
	url := config.Database.Url
	var err error
	Db, err = gorm.Open(driver, url)
	if err != nil {
		log.Fatal("数据库连接失败")
	}
	sqlDb := Db.DB()
	sqlDb.SetMaxIdleConns(config.Database.MaxIdle)
	sqlDb.SetMaxOpenConns(config.Database.MaxOpen)
	sqlDb.SetConnMaxLifetime(time.Second * config.Database.MaxLifetime)
}
