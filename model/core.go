package model

import (
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var DB *gorm.DB

func createAdminAccount() {
	var user User
	DB.Where(User{Role: common.RoleAdminUser}).Attrs(User{
		Username:    "admin",
		Password:    "123456",
		Role:        common.RoleAdminUser,
		Status:      common.UserStatusEnabled,
		DisplayName: "Administrator",
	}).FirstOrCreate(&user)
}

// CountTable 获取某个表长度
func CountTable(tableName string) (num int) {
	DB.Table(tableName).Count(&num)
	return
}

func InitDB() (db *gorm.DB, err error) {
	if *common.SqlDsn == "" {
		db, err = gorm.Open("mysql", os.Getenv("Sql_Dsn"))
	} else {
		db, err = gorm.Open("mysql", *common.SqlDsn)
	}

	if err == nil {
		DB = db
		db.AutoMigrate(&File{})
		db.AutoMigrate(&User{})
		db.AutoMigrate(&Option{})
		createAdminAccount()

		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}
