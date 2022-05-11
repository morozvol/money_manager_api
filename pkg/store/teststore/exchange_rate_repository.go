package teststore

import (
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
	"time"
)

type ExchangeRateRepository struct {
	store *Store
	rates map[int]*model.ExchangeRate
}

func (r ExchangeRateRepository) Get(idCurrencyFrom, idCurrencyTo int, time time.Time) (float32, error) {
	for _, r := range r.rates {
		if r.IdCurrencyTo == idCurrencyTo && r.IdCurrencyFrom == idCurrencyFrom && r.Date == time {
			return r.Rate, nil
		}
	}
	return 0, store.ErrRecordNotFound
}

func (r ExchangeRateRepository) Create(rate *model.ExchangeRate) error {
	rate.Id = len(r.rates) + 1
	r.rates[rate.Id] = rate
	return nil
}
