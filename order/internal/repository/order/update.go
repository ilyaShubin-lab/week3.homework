package order

import (
	"context"
	"fmt"

	"boilerplates/order/internal/model"
	"boilerplates/order/internal/repository/converter"
)

func (r *repository) Update(ctx context.Context, order model.Order) error {
	repoOrder := converter.OrderToRepoModel(order)

	const query = `
	UPDATE orders
		SET status           = $2,
		    transaction_uuid = $3,
		    payment_method   = $4,
		    updated_at       = now()
		WHERE order_uuid = $1`

	tag, err := r.pool.Exec(ctx, query,
		repoOrder.OrderUUID,
		repoOrder.Status,
		repoOrder.TransactionUUID,
		repoOrder.PaymentMethod,
	)
	if err != nil {
		return fmt.Errorf("update order: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return model.ErrOrderNotFound
	}

	return nil
}
