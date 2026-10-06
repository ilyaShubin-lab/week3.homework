package v1

import (
	"context"
	"net/http"

	orderV1 "boilerplates/shared/pkg/openapi/order/v1"
)

func (a *api) NewError(ctx context.Context, err error) *orderV1.GenericErrorStatusCode {
	return &orderV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: orderV1.GenericError{
			Code:    orderV1.NewOptInt32(http.StatusInternalServerError),
			Message: orderV1.NewOptString(err.Error()),
		},
	}
}
