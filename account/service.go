package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Repository
}

func newService(r Repository) Service {
	return &accountService{r}
}

func (a *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	account := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}

	if err := a.repository.PutAccount(ctx, *account); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return a.repository.GetAccountByID(ctx, id)
}

func (a *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return a.repository.ListAccounts(ctx, skip, take)
}
