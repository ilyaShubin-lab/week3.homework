-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
order_uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
user_uuid UUID NOT NULL,
part_uuids UUID[] NOT NULL,
total_price NUMERIC(12, 2) NOT NULL,
transaction_uuid UUID,
payment_method TEXT
               CHECK (payment_method IN ('UNKNOWN', 'CARD', 'SBP', 'CREDIT_CARD', 'INVESTOR_MONEY')),
status         TEXT NOT NULL DEFAULT 'PENDING_PAYMENT'
               CHECK (status IN ('PENDING_PAYMENT', 'PAID','CANCELLED')),
created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at     TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS orders;
