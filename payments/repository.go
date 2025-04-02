package payments

import (
	"fmt"
	"context"

	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot/models"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r Repository) RefundPayment(ctx context.Context, paymentChargeID string) error {
	return nil
}

func (r Repository) FindPayment(ctx context.Context, paymentChargeID string) (*merchant.Payment, error) {
	const query = `
SELECT id, user_id, refunded, provider_payment_charge_id, invoice_payload, currency, total_amount
FROM %s WHERE telegram_payment_charge_id = $1 LIMIT 1
	`

	payment := &merchant.Payment{
		SuccessfulPayment: &models.SuccessfulPayment{
			TelegramPaymentChargeID: paymentChargeID,
		},
	}

	err := r.pool.QueryRow(ctx, r.table(query), paymentChargeID).Scan(&payment.ID, &payment.UserID, &payment.Refunded, &payment.ProviderPaymentChargeID,
		&payment.InvoicePayload, &payment.Currency, &payment.TotalAmount)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r Repository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
