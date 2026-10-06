package payment

import (
	"context"
	"log"

	"boilerplates/payment/internal/model"
	"github.com/google/uuid"
)

func (s *service) Pay(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error) {
	if paymentMethod == "" || paymentMethod == "PAYMENT_METHOD_UNSPECIFIED" {
		return "", model.ErrInvalidPaymentMethod
	}
	transactionUUID := uuid.NewString()
	log.Printf("payment: orderUUID: %s, userUUID: %s, paymentMethod: %s, transactionUUID: %s", orderUUID, userUUID, paymentMethod, transactionUUID)
	return transactionUUID, nil
}
