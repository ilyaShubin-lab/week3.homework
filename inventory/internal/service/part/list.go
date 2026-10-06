package part

import (
	"context"

	"boilerplates/inventory/internal/model"
)

func (s *service) List(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	return s.partRepository.List(ctx, filter)
}
