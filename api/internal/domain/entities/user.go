package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"-"`
	Email     string             `bson:"email" json:"email"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// IsActive returns true if the user is active
func (u *User) IsActive() bool {
	return u.Status == "active"
}

// SetPassword sets the user password (in production, this should be hashed)
func (u *User) SetPassword(password string) {
	u.Password = password // TODO: Hash password in production
}

// ValidateEmail validates the user email format
func (u *User) ValidateEmail() bool {
	// Basic email validation - in production, use proper validation
	return len(u.Email) > 0 && len(u.Email) <= 255
}
