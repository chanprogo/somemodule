package mysqlsql

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// AllOperat 发送数据库查询语句
var AllOperat chan string

// MyData ...
type MyData struct {
	Name  string
	State int
}

func saveToChan() {
	myData := new(MyData)
	myData.Name = "Chan" + strconv.FormatInt(time.Now().Unix(), 10)
	myData.State = 1
	myData.Insert()
}

// Insert ...
func (thi *MyData) Insert() {
	var bufname, bufvalue bytes.Buffer

	bufname.WriteString("INSERT INTO example_table (")
	bufvalue.WriteString(" values(")
	if thi.Name != "" {
		bufname.WriteString("name")
		bufvalue.WriteString("'" + thi.Name + "'")
	}
	if thi.State != 0 {
		bufname.WriteString(", state")
		bufvalue.WriteString("," + strconv.Itoa(thi.State))
	}
	bufname.WriteString(")")
	bufvalue.WriteString(")")
	bufname.WriteString(bufvalue.String())

	fmt.Println()
	fmt.Println(bufname.String())
	fmt.Println()

	AllOperat <- bufname.String()
}

// SQLInsert ...
func SQLInsert(db *sql.DB) {
	for query := range AllOperat {
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}

func TestMysqlSql(t *testing.T) {

	host := "112.74.172.81"
	port := "3306"
	user := "root"
	password := "123456"
	database := "mytestdb"

	dbHander := OpenSQL(user, password, host, port, database)
	dbHander.SetMaxOpenConns(100)
	dbHander.SetMaxIdleConns(50)
	err := dbHander.Ping()
	if err != nil {
		dbHander.Close()
		fmt.Println("Ping err=" + err.Error())
		time.Sleep(time.Second * 1)
		return
	}

	defer func() {
		if dbHander != nil {
			dbHander.Close()
		}
	}()

	// Example 1
	AllOperat = make(chan string, 20000)
	go SQLInsert(dbHander)
	saveToChan()

	// Example 2
	names := GetNames(dbHander)
	fmt.Printf("%+v", names)
	fmt.Println()

	now := time.Now()
	next := now.Add(time.Second * 5)
	ti := time.NewTimer(next.Sub(now))
	<-ti.C
}
