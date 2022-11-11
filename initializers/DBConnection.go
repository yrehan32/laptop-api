package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	var err error
	var db_user = os.Getenv("DB_USER")
	var db_password = os.Getenv("DB_PASSWORD")
	var db_host = os.Getenv("DB_HOST")
	var db_port = os.Getenv("DB_PORT")
	var db_name = os.Getenv("DB_NAME")

	dsn := db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}