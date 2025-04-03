package payments

import (
	"context"
	merchant "github.com/bd878/merchant_bot"
)

type ClientsRepository interface {
	Find(ctx context.Context, id int64) (*merchant.Chat, error)
}

type ClientsDomain struct {
	repo ClientsRepository
}

func NewClientsDomain(repo ClientsRepository) *ClientsDomain {
	return &ClientsDomain{repo}
}

func (d ClientsDomain) FindClient(ctx context.Context, id int64) (*merchant.Chat, error) {
	return d.repo.Find(ctx, id)
}

func (d ClientsDomain) GetLocale(ctx context.Context, id int64) merchant.LangCode {
	c, _ := d.FindClient(ctx, id)
	return c.Lang
}