package order

import (
	"context"
	"fmt"

	"boilerplates/order/internal/model"
	"boilerplates/order/internal/repository/converter"
)

func (r *repository) Create(ctx context.Context, order model.Order) error {
	repoOrder := converter.OrderToRepoModel(order)

	const query = `
		INSERT INTO orders (order_uuid, user_uuid, part_uuids, total_price,
		                    transaction_uuid, payment_method, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.pool.Exec(ctx, query,
		repoOrder.OrderUUID,
		repoOrder.UserUUID,
		repoOrder.PartUUIDs,
		repoOrder.TotalPrice,
		repoOrder.TransactionUUID,
		repoOrder.PaymentMethod,
		repoOrder.Status,
	)
	if err != nil {
		return fmt.Errorf("insert order: %w", err)
	}
	return nil
}
