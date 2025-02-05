package database

import(
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil{
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	fmt.Println("✅ Database connection established successfully")
}