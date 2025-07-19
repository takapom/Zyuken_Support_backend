package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takagiyuuki/zyuken-backend/handler"
	"github.com/takagiyuuki/zyuken-backend/model"
	"github.com/takagiyuuki/zyuken-backend/repository"
	"github.com/takagiyuuki/zyuken-backend/service"
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
	log.Println("✅ データベースに接続しました")

	// AutoMigrate にモデルを並べる
	if err := db.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.School{},
		&model.Schedule{},
		&model.Report{},
		&model.MockExamScore{},
		&model.MockExam{},
		&model.Cost{},
		&model.Budget{},
		&model.School{},
		&model.Budget{},
		&model.UserSettings{},
	); err != nil {
		log.Fatal("マイグレーションに失敗しました:", err)
	}
	log.Println("✅ マイグレーション完了！")

	// 各層の初期化
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// ユーザーからのアクセスは一度routerに入る
	router := gin.Default()

	// CORS設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, // フロントエンドのURL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ルートの登録
	userHandler.RegisterRoutes(router)

	// サーバー起動
	log.Println("🚀 サーバーを起動します！: http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("サーバーの起動に失敗しました:", err)
	}
}
