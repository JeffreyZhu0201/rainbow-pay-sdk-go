package main

import (
	"fmt"
	"log"

	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/models"
	Payservice "github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/pkg/PayService"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件（开发环境）
	// 生产环境可通过环境变量指定加载其他文件（如 .env.production）
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("加载 .env 文件失败: %v", err)
	}

	// 初始化 Gin 路由
	r := gin.Default()

	// 注册路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			// "env":     os.Getenv("APP_ENV"), // 读取自定义环境变量
			// "db_host": os.Getenv("DB_HOST"),
		})
	})

	// 启动服务（端口从环境变量读取，若未设置则使用默认 8080）
	port := ":" + "8080"
	if port == ":" { // 处理空值（如未在 .env 中定义 APP_PORT）
		port = ":8080"
	}
	log.Printf("服务启动，端口: %s", port)
	log.Fatal(r.Run(port))

	fmt.Println("hello Go!")
	Payservice.CreateOrder(models.Order{})
	// Payservice.Notify()
}
