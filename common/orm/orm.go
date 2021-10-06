package orm

import (
	"log"
	"student_server/config"
	"time"

	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {
	var err error
	Db, err = gorm.Open(config.Database.Dialect, &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDb, _ := Db.DB()
	sqlDb.SetMaxIdleConns(config.Database.MaxIdle)
	sqlDb.SetMaxOpenConns(config.Database.MaxOpen)
	sqlDb.SetConnMaxLifetime(time.Second * config.Database.MaxLifetime)
}
