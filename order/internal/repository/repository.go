package repository

import (
	"context"

	"boilerplates/order/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, orderUUID string) (model.Order, error)
	Update(ctx context.Context, order model.Order) error
}
