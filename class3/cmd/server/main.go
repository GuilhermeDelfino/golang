package main

import (
	"github.com/GuilhermeDelfino/golang/class3/internal/repository"
	"github.com/GuilhermeDelfino/golang/class3/internal/service"
	"github.com/GuilhermeDelfino/golang/class3/pkg/config"
	"github.com/GuilhermeDelfino/golang/class3/pkg/database/mysql"
	"github.com/GuilhermeDelfino/golang/class3/pkg/rest"
)

func main() {
	config.ConfigDotEnv()
	conn := mysql.GetDatabaseConnection()

	productRepository := repository.CreateProductRepository(conn)
	productService := service.CreateProductService(productRepository)

	rest.CreateServer(productService)
}
