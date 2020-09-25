package services

import (
	"github.com/djsxianglei/iris-demo/models"
	"github.com/djsxianglei/iris-demo/repositories"
)

type IProductService interface {
	GetProductById(int64) (*models.Product, error)
	GetAllProduct() ([]*models.Product, error)
	DeleteProductById(int64) bool
	InsertProduct(product *models.Product) (int64, error)
	UpdateProduct(product *models.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

//初始化
func NewProductService(respository repositories.IProduct) IProductService {
	return &ProductService{respository}
}

func (p *ProductService) GetProductById(productId int64) (*models.Product, error) {
	return p.productRepository.SelectByKey(productId)
}

func (p *ProductService) GetAllProduct() ([]*models.Product, error) {
	return p.productRepository.SelectAll()
}

func (p *ProductService) DeleteProductById(productId int64) bool {
	return p.productRepository.Delete(productId)
}
func (p *ProductService) InsertProduct(product *models.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

func (p *ProductService) UpdateProduct(product *models.Product) error {
	return p.productRepository.Update(product)
}
