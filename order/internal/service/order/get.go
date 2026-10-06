package order

import (
	"context"

	"boilerplates/order/internal/model"
)

func (s *service) Get(ctx context.Context, orderUUID string) (model.Order, error) {
	order, err := s.orderRepository.Get(ctx, orderUUID)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
