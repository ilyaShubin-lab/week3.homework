package v1

import (
	"context"

	"boilerplates/order/internal/client/converter"
	"boilerplates/order/internal/model"
	paymentV1 "boilerplates/shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, orderUUID, userUUID string, method model.PaymentMethod) (string, error) {
	resp, err := c.generatedClient.PayOrder(ctx, &paymentV1.PayOrderRequest{
		OrderUuid:     orderUUID,
		UserUuid:      userUUID,
		PaymentMethod: converter.PaymentMethodToProto(method),
	})
	if err != nil {
		return "", model.ErrPaymentUnavailable
	}

	return resp.GetTransactionUuid(), nil
}
