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

// IsActive returns true if the subscription is active
func (s *Subscription) IsActive() bool {
	return s.Status == "active"
}

// IsCancelled returns true if the subscription is cancelled
func (s *Subscription) IsCancelled() bool {
	return s.Status == "cancelled"
}

// IsExpired returns true if the subscription is expired
func (s *Subscription) IsExpired() bool {
	return s.Status == "expired" || (s.EndDate != nil && s.EndDate.Before(time.Now()))
}

// Cancel cancels the subscription
func (s *Subscription) Cancel() {
	s.Status = "cancelled"
	now := time.Now()
	s.EndDate = &now
	s.UpdatedAt = now
}

// Renew renews the subscription for another billing cycle
func (s *Subscription) Renew() {
	if s.IsActive() {
		// Add one month to next billing (assuming monthly for simplicity)
		s.NextBilling = s.NextBilling.AddDate(0, 1, 0)
		s.UpdatedAt = time.Now()
	}
}

// DaysUntilNextBilling returns the number of days until the next billing
func (s *Subscription) DaysUntilNextBilling() int {
	duration := time.Until(s.NextBilling)
	return int(duration.Hours() / 24)
}
