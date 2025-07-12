package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	BillingType string             `bson:"billing_type" json:"billing_type"`
	Category    string             `bson:"category" json:"category"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// IsActive returns true if the product is active
func (p *Product) IsActive() bool {
	return p.Status == "active"
}

// IsMonthly returns true if the product has monthly billing
func (p *Product) IsMonthly() bool {
	return p.BillingType == "monthly"
}

// IsYearly returns true if the product has yearly billing
func (p *Product) IsYearly() bool {
	return p.BillingType == "yearly"
}

// ValidatePrice validates the product price
func (p *Product) ValidatePrice() bool {
	return p.Price > 0
}

// GetMonthlyPrice returns the monthly equivalent price
func (p *Product) GetMonthlyPrice() float64 {
	if p.IsMonthly() {
		return p.Price
	}
	if p.IsYearly() {
		return p.Price / 12
	}
	return p.Price // Default to current price
}
