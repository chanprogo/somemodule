package gormmodel

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

// Setup initializes the database instance
func Setup(dbType, user, password, host, name, tableprefix string) {
	var err error
	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tableprefix + defaultTableName
	}

	Db.SingularTable(true)
	Db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	Db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	Db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer Db.Close()
}
