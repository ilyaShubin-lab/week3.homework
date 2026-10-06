package converter

import (
	"boilerplates/inventory/internal/model"
	repoModel "boilerplates/inventory/internal/repository/model"
)

func PartToModel(part repoModel.Part) model.Part {
	return model.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model.Category(part.Category),
		Dimensions:    DimensionsToModel(part.Dimensions),
		Manufacturer:  ManufacturerToModel(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      MetadataToModel(part.Metadata),
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func PartsToModel(parts []repoModel.Part) []model.Part {
	partsSlice := make([]model.Part, 0, len(parts))

	for _, v := range parts {
		partsSlice = append(partsSlice, PartToModel(v))
	}
	return partsSlice
}

func DimensionsToModel(dimensions *repoModel.Dimensions) *model.Dimensions {
	if dimensions == nil {
		return nil
	}

	return &model.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func ManufacturerToModel(manufacturer *repoModel.Manufacturer) *model.Manufacturer {
	if manufacturer == nil {
		return nil
	}

	return &model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func MetadataToModel(metadata map[string]repoModel.Value) map[string]model.Value {
	if metadata == nil {
		return nil
	}

	result := make(map[string]model.Value, len(metadata))
	for k, v := range metadata {
		result[k] = model.Value{
			StringValue: v.StringValue,
			Int64Value:  v.Int64Value,
			DoubleValue: v.DoubleValue,
			BoolValue:   v.BoolValue,
		}
	}

	return result
}

func PartsFilterToRepoModel(filter model.PartsFilter) repoModel.PartsFilter {
	categories := make([]repoModel.Category, 0, len(filter.Categories))

	for _, v := range filter.Categories {
		categories = append(categories, repoModel.Category(v))
	}

	return repoModel.PartsFilter{
		UUIDs:                 filter.UUIDs,
		Names:                 filter.Names,
		Categories:            categories,
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}
