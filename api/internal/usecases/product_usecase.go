package usecases

import (
	"context"
	"errors"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"github.com/frtasoniero/subsmanager/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductUseCase handles product business logic
type ProductUseCase struct {
	productRepo repositories.ProductRepository
}

// NewProductUseCase creates a new product use case
func NewProductUseCase(productRepo repositories.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product
func (uc *ProductUseCase) CreateProduct(ctx context.Context, product *entities.Product) error {
	// Validate price
	if !product.ValidatePrice() {
		return errors.New("invalid product price")
	}

	// Check if product already exists
	existingProduct, _ := uc.productRepo.GetByName(ctx, product.Name)
	if existingProduct != nil {
		return errors.New("product with this name already exists")
	}

	return uc.productRepo.Create(ctx, product)
}

// GetProductByID retrieves a product by ID
func (uc *ProductUseCase) GetProductByID(ctx context.Context, id primitive.ObjectID) (*entities.Product, error) {
	return uc.productRepo.GetByID(ctx, id)
}

// GetAllProducts retrieves all products
func (uc *ProductUseCase) GetAllProducts(ctx context.Context) ([]*entities.Product, error) {
	return uc.productRepo.GetAll(ctx)
}

// GetActiveProducts retrieves all active products
func (uc *ProductUseCase) GetActiveProducts(ctx context.Context) ([]*entities.Product, error) {
	return uc.productRepo.GetActive(ctx)
}

// GetProductsByCategory retrieves products by category
func (uc *ProductUseCase) GetProductsByCategory(ctx context.Context, category string) ([]*entities.Product, error) {
	return uc.productRepo.GetByCategory(ctx, category)
}

// UpdateProduct updates a product
func (uc *ProductUseCase) UpdateProduct(ctx context.Context, product *entities.Product) error {
	// Validate price
	if !product.ValidatePrice() {
		return errors.New("invalid product price")
	}

	return uc.productRepo.Update(ctx, product)
}

// DeleteProduct deletes a product
func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	return uc.productRepo.Delete(ctx, id)
}

// DeactivateProduct deactivates a product
func (uc *ProductUseCase) DeactivateProduct(ctx context.Context, id primitive.ObjectID) error {
	product, err := uc.productRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	product.Status = "inactive"
	return uc.productRepo.Update(ctx, product)
}
