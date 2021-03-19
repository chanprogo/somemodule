package model

import (
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func NewPostgresOrm(host, port, user, password, dbname string, maxOpen, maxIdle int) {

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Orm, err = xorm.NewEngine("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := Orm.Ping(); err != nil {
		panic("数据库连接失败，err：" + err.Error())
	}

	Orm.SetMaxOpenConns(maxOpen) // Orm.SetMaxOpenConns(1 << 10)
	Orm.SetMaxIdleConns(maxIdle) // Orm.SetMaxIdleConns(1 << 10)

	// if gin.Mode() != gin.ReleaseMode {
	// 	Orm.ShowSQL(true)
	// }
}
