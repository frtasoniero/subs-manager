package repositories

import (
	"context"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubscriptionRepository defines the interface for subscription data operations
type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *entities.Subscription) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Subscription, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*entities.SubscriptionWithProduct, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]*entities.Subscription, error)
	GetAll(ctx context.Context) ([]*entities.SubscriptionWithProduct, error)
	GetActive(ctx context.Context) ([]*entities.SubscriptionWithProduct, error)
	GetExpiring(ctx context.Context, days int) ([]*entities.SubscriptionWithProduct, error)
	Update(ctx context.Context, subscription *entities.Subscription) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	Count(ctx context.Context) (int64, error)
}
