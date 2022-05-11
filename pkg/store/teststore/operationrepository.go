package teststore

import (
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
	"time"
)

type OperationRepository struct {
	store      *Store
	operations map[int]*model.Operation
}

// Create ...
func (r *OperationRepository) Create(o ...*model.Operation) error {
	for _, op := range o {
		op.Id = len(r.operations) + 1
		r.operations[op.Id] = op
	}
	return nil
}

// Find ...
func (r *OperationRepository) Find(id int) (*model.Operation, error) {
	o, ok := r.operations[id]
	if ok {
		return o, nil
	}
	return nil, store.ErrRecordNotFound
}
func (r *OperationRepository) Get(dateFrom, dateTo time.Time, idUser int) ([]model.Operation, error) {
	res := make([]model.Operation, 0)
	for _, o := range r.operations {

		if o.Time.Before(dateTo) && o.Time.After(dateFrom) {
			res = append(res, *o)
		}
	}
	return res, nil
}

func (r *OperationRepository) Update(int) (*model.Operation, error) {
	return nil, nil //TODO: +

}
func (r *OperationRepository) Delete(id int) error {
	_, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	delete(r.operations, id)
	return nil
}
