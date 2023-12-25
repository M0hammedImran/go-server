package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	docs "github.com/m0hammedimran/go-server/api"
	v1 "github.com/m0hammedimran/go-server/cmd/router/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var mode string = os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.GET("/ping", PingHandler)

	v1.Register(router.Group(""))

	docs.SwaggerInfo.BasePath = ""
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method Not Allowed"})
	})
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run()
}

type PingRequest struct {
	Message string              `json:"message"`
	Data    map[string][]string `json:"data"`
}

// PingExample godoc
// @BasePath
// @Summary ping
// @Description do ping
// @Tags PING
// @Accept json
// @Produce json
// @Success 200 {object} PingRequest "{\"message\": \"pong\", \"data\": { \"Accept\": [\"application/json\"]}}"
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, PingRequest{Message: "pong", Data: c.Request.Header})
}
