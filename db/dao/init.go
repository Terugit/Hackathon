package dao

import (
	"db/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	dsn = fmt.Sprintf("uttc:uttc@term2-teruki-nishino:us-central1:uttc1/hackathon")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	//if err := db.AutoMigrate(&model.Account{}); err != nil {
	//	log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	//}
	if err := db.AutoMigrate(&model.Thank{}); err != nil {
		log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("fail: db.AutoMigrate, %v\n", err)
	}

}
