package service

import (
	"github.com/GuilhermeDelfino/golang/class3/internal/entity"
	"github.com/GuilhermeDelfino/golang/class3/internal/repository"
)

type ProductService interface {
	FindAll() (*[]entity.Product, error)
	FindById(id *int) (*entity.Product, error)
	Save(product *entity.Product) error
}
type productService struct {
	repository repository.ProductRepository
}

func CreateProductService(repository repository.ProductRepository) ProductService {
	return productService{repository}
}

func (r productService) FindAll() (*[]entity.Product, error) {
	return r.repository.FindAll()
}
func (r productService) Save(p *entity.Product) error {
	return r.repository.Save(p)
}
func (r productService) FindById(id *int) (*entity.Product, error) {
	return r.repository.FindById(id)
}
