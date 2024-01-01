//untuk koneksi ke database
package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)
var DB *gorm.DB

func ConnectDatabase(){
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_jwt_mux"))

	if err != nil{
		panic(err)
	}

	db.AutoMigrate(&User{})
	DB = db
}
