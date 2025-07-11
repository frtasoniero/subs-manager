package repositories

import (
	"context"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoSubscriptionRepository struct {
	collection *mongo.Collection
	db         *mongo.Database
}

func NewMongoSubscriptionRepository(db *mongo.Database) *MongoSubscriptionRepository {
	return &MongoSubscriptionRepository{
		collection: db.Collection("subscriptions"),
		db:         db,
	}
}

func (r *MongoSubscriptionRepository) Create(ctx context.Context, subscription *entities.Subscription) error {
	_, err := r.collection.InsertOne(ctx, subscription)
	return err
}

func (r *MongoSubscriptionRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Subscription, error) {
	var subscription entities.Subscription
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *MongoSubscriptionRepository) GetAll(ctx context.Context) ([]*entities.SubscriptionWithProduct, error) {
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "product_id",
				"foreignField": "_id",
				"as":           "product",
			},
		},
		{
			"$unwind": "$product",
		},
		{
			"$project": bson.M{
				"id":           "$_id",
				"user_id":      1,
				"product_id":   1,
				"product_name": "$product.name",
				"description":  "$product.description",
				"price":        "$product.price",
				"status":       1,
				"start_date":   1,
				"end_date":     1,
				"next_billing": 1,
				"created_at":   1,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var subscriptions []*entities.SubscriptionWithProduct
	if err = cursor.All(ctx, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *MongoSubscriptionRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*entities.SubscriptionWithProduct, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"user_id": userID},
		},
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "product_id",
				"foreignField": "_id",
				"as":           "product",
			},
		},
		{
			"$unwind": "$product",
		},
		{
			"$project": bson.M{
				"id":           "$_id",
				"user_id":      1,
				"product_id":   1,
				"product_name": "$product.name",
				"description":  "$product.description",
				"price":        "$product.price",
				"status":       1,
				"start_date":   1,
				"end_date":     1,
				"next_billing": 1,
				"created_at":   1,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var subscriptions []*entities.SubscriptionWithProduct
	if err = cursor.All(ctx, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *MongoSubscriptionRepository) Update(ctx context.Context, subscription *entities.Subscription) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": subscription.ID}, bson.M{"$set": subscription})
	return err
}

func (r *MongoSubscriptionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
