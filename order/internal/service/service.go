package service

import (
	"context"

	"boilerplates/order/internal/model"
)

type OrderService interface {
	Create(ctx context.Context, userUUID string, partUUIDs []string) (model.Order, error)
	Get(ctx context.Context, orderUUID string) (model.Order, error)
	Pay(ctx context.Context, orderUUID string, method model.PaymentMethod) (string, error)
	Cancel(ctx context.Context, orderUUID string) error
}
