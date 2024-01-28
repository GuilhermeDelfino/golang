package route

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/GuilhermeDelfino/golang/class3/internal/entity"
	"github.com/GuilhermeDelfino/golang/class3/internal/service"
	"github.com/GuilhermeDelfino/golang/class3/pkg/rest/route/dto"
	"github.com/gin-gonic/gin"
)

type productRoute struct {
	service service.ProductService
}

func CreateProductRoute(server *gin.Engine, service service.ProductService) {
	group := server.Group("api/v1/products")
	route := productRoute{service}

	group.GET("/", route.findAll)
	group.GET("/:id", route.findById)
	group.POST("/", route.save)

}

func (r productRoute) findAll(ctx *gin.Context) {
	data, err := r.service.FindAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    data,
	})
}

func (r productRoute) findById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else if id <= 0 {
		ctx.AbortWithError(http.StatusNotFound, errors.New("invalid id"))
	}

	product, err := r.service.FindById(&id)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    product,
	})
}

func (r productRoute) save(ctx *gin.Context) {
	var productDto dto.ProductSaveDto
	err := ctx.ShouldBindJSON(&productDto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	product := entity.Product{Title: productDto.Title, Description: productDto.Description, Price: productDto.Price}
	err = r.service.Save(&product)

	if err != nil {
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    &product,
	})
}
