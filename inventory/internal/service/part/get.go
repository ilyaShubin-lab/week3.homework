package part

import (
	"context"

	"boilerplates/inventory/internal/model"
)

func (s service) Get(ctx context.Context, uuid string) (model.Part, error) {
	return s.partRepository.Get(ctx, uuid)
}
