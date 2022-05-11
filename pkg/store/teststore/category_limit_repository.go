package teststore

import "github.com/morozvol/money_manager_api/pkg/model"

type CategoryLimitRepository struct {
	store  *Store
	limits map[int]*model.CategoryLimit
}
