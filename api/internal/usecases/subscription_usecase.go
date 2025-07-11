package usecases

import (
	"context"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"github.com/frtasoniero/subsmanager/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionUseCase struct {
	subscriptionRepo repositories.SubscriptionRepository
}

func NewSubscriptionUseCase(subscriptionRepo repositories.SubscriptionRepository) *SubscriptionUseCase {
	return &SubscriptionUseCase{
		subscriptionRepo: subscriptionRepo,
	}
}

func (uc *SubscriptionUseCase) GetAllSubscriptions(ctx context.Context) ([]*entities.SubscriptionWithProduct, error) {
	return uc.subscriptionRepo.GetAll(ctx)
}

func (uc *SubscriptionUseCase) GetSubscriptionsByUser(ctx context.Context, userID primitive.ObjectID) ([]*entities.SubscriptionWithProduct, error) {
	return uc.subscriptionRepo.GetByUserID(ctx, userID)
}

func (uc *SubscriptionUseCase) CreateSubscription(ctx context.Context, subscription *entities.Subscription) error {
	return uc.subscriptionRepo.Create(ctx, subscription)
}
