package utils

import (
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	user     = "root"
	password = "root.rb.k201"
	dbname   = "roberto_demo"
	dialect  = "mysql"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *gorm.DB {
	//db, err := gorm.Open(dialect, "root:root.rb.k201@/roberto_demo?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
