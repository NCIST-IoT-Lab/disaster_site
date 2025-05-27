package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"disaster_site_information_management_system/internal/database"
	"disaster_site_information_management_system/internal/handlers"
	"disaster_site_information_management_system/internal/middleware"
	"disaster_site_information_management_system/internal/models"
	"disaster_site_information_management_system/internal/utils"
)

func init() {
    // 加载.env文件到环境变量
    err := godotenv.Load("configs/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	// 开启调试日志
	log.Println("Starting application...")

	// 初始化数据库
	dbConfig := database.Config{
		Username: utils.GetEnv("DB_USERNAME", "root"),
		Password: utils.GetEnv("DB_PASSWORD", "password"),
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnv("DB_PORT", "3306"),
		DBName:   utils.GetEnv("DB_NAME", "disaster_site_db"),
	}
	log.Printf("Database config: %+v", dbConfig)
	db := database.InitDB(dbConfig)

	// 自动迁移数据库模型
	if err := models.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// 初始化Gin路由
	r := gin.Default()

	// 添加全局中间件
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	ctx := context.Background()

	// 创建处理程序
	eventHandler := handlers.NewEventHandler(ctx, db)
	userHandler := handlers.NewUserHandler(db)
	taskHandler := handlers.NewTaskHandler(db)

	// 注册API路由
	apiRoutes := r.Group("/disaster_site/api")
	{
		// 添加健康检查路由
		apiRoutes.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"service": "disaster_site_information_management_system",
			})
		})

		// 事件路由
		events := apiRoutes.Group("/events")
		{
			events.GET("/", eventHandler.GetAllEvents)
			events.GET("/:id", eventHandler.GetEvent)
			events.POST("/", eventHandler.CreateEvent)
			events.PUT("/:id", eventHandler.UpdateEvent)
			events.DELETE("/:id", eventHandler.DeleteEvent)
		}

		// 用户路由
		users := apiRoutes.Group("/users")
		{
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("/", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// 任务路由
		tasks := apiRoutes.Group("/tasks")
		{
			tasks.GET("/", taskHandler.GetAllTasks)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.POST("/", taskHandler.CreateTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

	// 启动服务器
	port := utils.GetEnv("SERVER_PORT", "8080")
	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
