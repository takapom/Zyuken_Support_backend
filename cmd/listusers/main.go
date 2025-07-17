package main

import (
	"fmt"
	"log"

	"github.com/takagiyuuki/zyuken-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DSN: dbname を zyukendb or testdb に合わせる
	dsn := "host=localhost user=takagiyuuki dbname=zyukendb port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続に失敗しました:", err)
	}

	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatal("ユーザー取得に失敗しました:", err)
	}

	if len(users) == 0 {
		fmt.Println("登録されているユーザーはいません。")
		return
	}

	fmt.Println("=== ユーザー一覧 ===")
	fmt.Printf("%-40s %-30s %-20s %-10s %-5s\n", "ID", "Email", "名前", "学部", "卒業年")
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	
	for _, user := range users {
		fmt.Printf("%-40s %-30s %-20s %-10s %-5d\n", 
			user.ID, 
			user.Email, 
			user.Name, 
			user.Department, 
			user.GraduationYear,
		)
	}
	
	fmt.Printf("\n合計: %d 人\n", len(users))
}