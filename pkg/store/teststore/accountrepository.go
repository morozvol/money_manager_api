package teststore

import (
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
)

type AccountRepository struct {
	store    *Store
	accounts map[int]*model.Account
}

// Create ...
func (r *AccountRepository) Create(a *model.Account) error {
	a.Id = len(r.accounts) + 1
	r.accounts[a.Id] = a
	return nil
}

// Find ...
func (r *AccountRepository) Find(id int) (*model.Account, error) {
	a, ok := r.accounts[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return a, nil
}

// FindByUserId ...
func (r *AccountRepository) FindByUserId(userId int) ([]model.Account, error) {
	var res []model.Account
	for _, a := range r.accounts {
		if a.IdUser == userId {
			res = append(res, *a)
		}
	}
	if len(res) == 0 {
		return nil, store.ErrRecordNotFound
	}
	return res, nil
}
