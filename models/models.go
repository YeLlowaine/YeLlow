package models

import (
	"fmt"
	// "log"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/YeLlowaine/YeLlow/pkg/setting"
	"github.com/YeLlowaine/YeLlow/pkg/logging"
	
	
)

var db *gorm.DB



//Model ...
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`

	CreatedAt time.Time `json:"created_at"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logging.Info(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxOpenConns(2000)
	db.DB().SetMaxIdleConns(1000)
}

//CloseDB ...
func CloseDB() {
	defer db.Close()
}
