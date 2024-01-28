package rest

import (
	"fmt"
	"os"

	"github.com/GuilhermeDelfino/golang/class3/internal/service"
	"github.com/GuilhermeDelfino/golang/class3/pkg/rest/middleware"
	"github.com/GuilhermeDelfino/golang/class3/pkg/rest/route"
	"github.com/gin-gonic/gin"
)

func CreateServer(productService service.ProductService) {
	port := os.Getenv("SERVER_PORT")
	server := gin.Default()

	server.Use(middleware.CORSMiddleware())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK, We're working!",
		})
	})

	route.CreateProductRoute(server, productService)

	server.Run(fmt.Sprintf("localhost:%s", port))
}
