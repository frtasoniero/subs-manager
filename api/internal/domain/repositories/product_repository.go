package repositories

import (
	"context"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Product, error)
	GetByName(ctx context.Context, name string) (*entities.Product, error)
	GetByCategory(ctx context.Context, category string) ([]*entities.Product, error)
	GetAll(ctx context.Context) ([]*entities.Product, error)
	GetActive(ctx context.Context) ([]*entities.Product, error)
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	Count(ctx context.Context) (int64, error)
}
