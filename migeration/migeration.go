package migeration

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"size:100"`
	Email    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
}

func main() {
	dsn := "root2:Zz123456789!@@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}
