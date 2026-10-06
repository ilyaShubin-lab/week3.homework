package v1

import (
	"context"
	"errors"

	"boilerplates/order/internal/converter"
	"boilerplates/order/internal/model"
	orderV1 "boilerplates/shared/pkg/openapi/order/v1"
)

func (a *api) GetOrder(ctx context.Context, params orderV1.GetOrderParams) (orderV1.GetOrderRes, error) {
	order, err := a.orderService.Get(ctx, params.OrderUUID.String())
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderV1.NotFoundError{Code: 404, Message: "order not found"}, nil
		}
		return &orderV1.InternalServerError{Code: 500, Message: "internal error"}, nil
	}

	dto := converter.OrderToDTO(order)
	return &dto, nil
}
