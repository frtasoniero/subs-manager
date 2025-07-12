package usecases

import (
	"context"
	"errors"

	"github.com/frtasoniero/subsmanager/internal/domain/entities"
	"github.com/frtasoniero/subsmanager/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUseCase handles user business logic
type UserUseCase struct {
	userRepo repositories.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(ctx context.Context, user *entities.User) error {
	// Validate email
	if !user.ValidateEmail() {
		return errors.New("invalid email format")
	}

	// Check if user already exists
	existingUser, _ := uc.userRepo.GetByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	existingUser, _ = uc.userRepo.GetByUsername(ctx, user.Username)
	if existingUser != nil {
		return errors.New("user with this username already exists")
	}

	// Set password (should be hashed in production)
	user.SetPassword(user.Password)

	return uc.userRepo.Create(ctx, user)
}

// GetUserByID retrieves a user by ID
func (uc *UserUseCase) GetUserByID(ctx context.Context, id primitive.ObjectID) (*entities.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

// GetAllUsers retrieves all users
func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]*entities.User, error) {
	return uc.userRepo.GetAll(ctx)
}

// UpdateUser updates a user
func (uc *UserUseCase) UpdateUser(ctx context.Context, user *entities.User) error {
	// Validate email
	if !user.ValidateEmail() {
		return errors.New("invalid email format")
	}

	return uc.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user
func (uc *UserUseCase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return uc.userRepo.Delete(ctx, id)
}

// AuthenticateUser authenticates a user by email and password
func (uc *UserUseCase) AuthenticateUser(ctx context.Context, email, password string) (*entities.User, error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// In production, compare hashed passwords
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive() {
		return nil, errors.New("user account is inactive")
	}

	return user, nil
}
