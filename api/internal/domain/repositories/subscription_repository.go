package repositories

import (
	"context"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *entities.Subscription) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Subscription, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*entities.SubscriptionWithProduct, error)
	GetAll(ctx context.Context) ([]*entities.SubscriptionWithProduct, error)
	Update(ctx context.Context, subscription *entities.Subscription) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
