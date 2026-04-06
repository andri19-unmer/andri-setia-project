package usecase

import (
	"context"

	"app-backend/internal/domain"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{productRepo: repo}
}

func (pu *productUsecase) Fetch(ctx context.Context) ([]domain.Product, error) {
	return pu.productRepo.Fetch(ctx)
}

func (pu *productUsecase) GetByID(ctx context.Context, id uint) (domain.Product, error) {
	return pu.productRepo.GetByID(ctx, id)
}

func (pu *productUsecase) Store(ctx context.Context, product *domain.Product) error {
	return pu.productRepo.Store(ctx, product)
}

func (pu *productUsecase) Update(ctx context.Context, product *domain.Product) error {
	return pu.productRepo.Update(ctx, product)
}

func (pu *productUsecase) Delete(ctx context.Context, id uint) error {
	return pu.productRepo.Delete(ctx, id)
}
