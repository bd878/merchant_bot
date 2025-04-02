package payments

import (
	"fmt"
	"context"

	"github.com/go-telegram/bot/models"
	"github.com/jackc/pgx/v5/pgxpool"
	merchant "github.com/bd878/merchant_bot"
)

type Repository struct {
	pool *pgxpool.Pool
	tableName string
}

func NewRepository(tableName string, pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
		tableName: tableName,
	}
}

func (r Repository) SavePayment(ctx context.Context, payment *merchant.Payment) error {
	const query = `
INSERT INTO %s (id, user_id, refunded, telegram_payment_charge_id, provider_payment_charge_id, invoice_payload, currency, total_amount)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`

	_, err := r.pool.Exec(ctx, r.table(query), payment.ID, payment.UserID, payment.Refunded, payment.TelegramPaymentChargeID,
		payment.ProviderPaymentChargeID, payment.InvoicePayload, payment.Currency, payment.TotalAmount)

	return err
}

func (r Repository) RefundPayment(ctx context.Context, id uint32) error {
	const query = `UPDATE %s SET refunded = true WHERE id = $1`

	_, err := r.pool.Exec(ctx, r.table(query), id)

	return err
}

func (r Repository) ListUserTransactions(ctx context.Context, userID int64, limit, offset int) ([]*merchant.Payment, error) {
	const query = `
SELECT id, refunded, telegram_payment_charge_id, provider_payment_charge_id, invoice_payload, currency, total_amount, created_at
FROM %s WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
	`;

	rows, err := r.pool.Query(ctx, r.table(query), userID, limit, offset)
	if err != nil {
		return nil, err
	}

	payments := make([]*merchant.Payment, 0)

	for rows.Next() {
		payment := &merchant.Payment{
			SuccessfulPayment: &models.SuccessfulPayment{},
			UserID: userID,
		}

		err = rows.Scan(&payment.ID, &payment.Refunded, &payment.TelegramPaymentChargeID, &payment.ProviderPaymentChargeID,
			&payment.InvoicePayload, &payment.Currency, &payment.TotalAmount, &payment.CreatedAt)
		if err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return payments, nil
}

func (r Repository) FindPayment(ctx context.Context, id uint32) (*merchant.Payment, error) {
	const query = `
SELECT user_id, refunded, telegram_payment_charge_id, provider_payment_charge_id, invoice_payload, currency, total_amount
FROM %s WHERE id = $1 LIMIT 1
	`

	payment := &merchant.Payment{
		ID: id,
		SuccessfulPayment: &models.SuccessfulPayment{
		},
	}

	err := r.pool.QueryRow(ctx, r.table(query), id).Scan(&payment.UserID, &payment.Refunded, &payment.TelegramPaymentChargeID, &payment.ProviderPaymentChargeID,
		&payment.InvoicePayload, &payment.Currency, &payment.TotalAmount)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r Repository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
