package v1

import (
	"context"

	"boilerplates/order/internal/client/converter"
	"boilerplates/order/internal/model"
	inventoryV1 "boilerplates/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	resp, err := c.generatedClient.ListParts(ctx, &inventoryV1.ListPartsRequest{
		Filter: &inventoryV1.PartsFilter{Uuids: filter.UUIDs},
	})
	if err != nil {
		return nil, model.ErrInventoryUnavailable
	}
	return converter.PartListToModel(resp.GetParts()), nil
}
