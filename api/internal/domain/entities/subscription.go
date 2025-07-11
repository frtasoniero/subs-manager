package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	Status       string             `bson:"status" json:"status"`
	StartDate    time.Time          `bson:"start_date" json:"start_date"`
	EndDate      *time.Time         `bson:"end_date,omitempty" json:"end_date,omitempty"`
	NextBilling  time.Time          `bson:"next_billing" json:"next_billing"`
	PriceAtStart float64            `bson:"price_at_start" json:"price_at_start"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type SubscriptionWithProduct struct {
	ID          primitive.ObjectID `json:"id"`
	UserID      primitive.ObjectID `json:"user_id"`
	ProductID   primitive.ObjectID `json:"product_id"`
	ProductName string             `json:"product_name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Status      string             `json:"status"`
	StartDate   time.Time          `json:"start_date"`
	EndDate     *time.Time         `json:"end_date,omitempty"`
	NextBilling time.Time          `json:"next_billing"`
	CreatedAt   time.Time          `json:"created_at"`
}
