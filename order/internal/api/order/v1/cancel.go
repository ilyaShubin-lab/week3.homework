package v1

import (
	"context"
	"errors"

	"boilerplates/order/internal/model"
	orderV1 "boilerplates/shared/pkg/openapi/order/v1"
)

func (a *api) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	err := a.orderService.Cancel(ctx, params.OrderUUID.String())
	if err != nil {
		switch {
		case errors.Is(err, model.ErrOrderNotFound):
			return &orderV1.NotFoundError{Code: 404, Message: "order not found"}, nil
		case errors.Is(err, model.ErrOrderAlreadyPaid):
			return &orderV1.ConflictError{Code: 409, Message: "order already paid"}, nil
		case errors.Is(err, model.ErrOrderCancelled):
			return &orderV1.ConflictError{Code: 409, Message: "order already cancelled"}, nil
		default:
			return &orderV1.InternalServerError{Code: 500, Message: "internal error"}, nil
		}
	}

	return &orderV1.CancelOrderNoContent{}, nil
}
