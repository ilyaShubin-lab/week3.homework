package order

import (
	"context"

	"boilerplates/order/internal/model"
)

func (s *service) Pay(ctx context.Context, orderUUID string, method model.PaymentMethod) (string, error) {
	order, err := s.orderRepository.Get(ctx, orderUUID)
	if err != nil {
		return "", err
	}

	//  платить дважды нельзя, отменённый оплатить нельзя.
	switch order.Status {
	case model.OrderStatusPaid:
		return "", model.ErrOrderAlreadyPaid
	case model.OrderStatusCancelled:
		return "", model.ErrOrderCancelled
	}

	txUUID, err := s.paymentClient.PayOrder(ctx, orderUUID, order.UserUUID, method)
	if err != nil {
		return "", err
	}

	order.Status = model.OrderStatusPaid
	order.TransactionUUID = &txUUID
	order.PaymentMethod = &method

	err = s.orderRepository.Update(ctx, order)
	if err != nil {
		return "", err
	}
	return txUUID, nil
}
