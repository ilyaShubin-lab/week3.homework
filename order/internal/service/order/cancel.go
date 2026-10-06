package order

import (
	"context"

	"boilerplates/order/internal/model"
)

func (s *service) Cancel(ctx context.Context, orderUUID string) error {
	order, err := s.orderRepository.Get(ctx, orderUUID)
	if err != nil {
		return err
	}

	//  платить дважды нельзя, отменённый оплатить нельзя.
	switch order.Status {
	case model.OrderStatusPaid:
		return model.ErrOrderAlreadyPaid
	case model.OrderStatusCancelled:
		return model.ErrOrderCancelled
	}

	order.Status = model.OrderStatusCancelled

	return s.orderRepository.Update(ctx, order)
}
