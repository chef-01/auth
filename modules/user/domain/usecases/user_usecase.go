package usecases

import (
	"auth/common/strings"
	"auth/modules/user/data/model"
	"auth/modules/user/domain/repository"
	"context"
	"errors"
	"time"
)


type UserUseCase struct {
	userRepo repository.UserRepository
	otpSvc   OTPService
	jwtSvc   JWTService
}

func NewUserUseCase(userRepo repository.UserRepository, otpSvc OTPService, jwtSvc JWTService) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		otpSvc:   otpSvc,
		jwtSvc:   jwtSvc,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user  *model.User) (*model.User, error) {
if user.Name == "" {
		return nil, errors.New("user " + strings.InvalidDataMsg + ": title and description are required")
	}
	now := time.Now()
	user.CreatedAt = now
	user.ModifiedAt = now

	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.ModifiedAt = time.Now()
	return uc.userRepo.UpdateUser(ctx, user)
}

// GetAllUsers retrieves all users.
func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return uc.userRepo.GetAllUsers(ctx)
}