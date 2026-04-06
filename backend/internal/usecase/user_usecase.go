package usecase

import (
	"context"

	"app-backend/internal/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) Fetch(ctx context.Context) ([]domain.User, error) {
	return u.userRepo.Fetch(ctx)
}

func (u *userUsecase) GetByID(ctx context.Context, id uint) (domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUsecase) Store(ctx context.Context, user *domain.User) error {
	return u.userRepo.Store(ctx, user)
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) error {
	return u.userRepo.Update(ctx, user)
}

func (u *userUsecase) Delete(ctx context.Context, id uint) error {
	return u.userRepo.Delete(ctx, id)
}
