package model

import "time"

type Part struct {
	UUID          string           `bson:"_id"`
	Name          string           `bson:"name"`
	Description   string           `bson:"description"`
	Price         float64          `bson:"price"`
	StockQuantity int64            `bson:"stock_quantity"`
	Category      int32            `bson:"category"`
	Dimensions    *Dimensions      `bson:"dimensions,omitempty"`
	Manufacturer  *Manufacturer    `bson:"manufacturer,omitempty"`
	Tags          []string         `bson:"tags,omitempty"`
	Metadata      map[string]Value `bson:"metadata,omitempty"`
	CreatedAt     time.Time        `bson:"created_at"`
	UpdatedAt     *time.Time       `bson:"updated_at,omitempty"`
}

type PartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}

type Dimensions struct {
	Length float64 `bson:"length"`
	Width  float64 `bson:"width"`
	Height float64 `bson:"height"`
	Weight float64 `bson:"weight"`
}

type Manufacturer struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Website string `bson:"website"`
}

type Value struct {
	StringValue *string  `bson:"string_value,omitempty"`
	Int64Value  *int64   `bson:"int64_value,omitempty"`
	DoubleValue *float64 `bson:"double_value,omitempty"`
	BoolValue   *bool    `bson:"bool_value,omitempty"`
}

type Category int

const (
	CategoryUnknown Category = iota
	CategoryEngine
	CategoryFuel
	CategoryPorthole
	CategoryWing
)
