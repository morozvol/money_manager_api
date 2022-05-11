package sqlstore

import (
	"context"
	"database/sql"
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
	"time"
)

type OperationRepository struct {
	store *Store
}

// Create ...
func (r *OperationRepository) Create(o ...*model.Operation) error {

	ctx := context.Background()
	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	var lastId int
	for _, operation := range o {
		operation.Time = time.Now()
		err = tx.QueryRow("SELECT public.apply_operation($1, $2, $3, $4, $5)",
			operation.IdAccount,
			operation.Sum,
			operation.Category.Id,
			operation.Description,
			operation.Time).Scan(&lastId)
		if err != nil {
			errRb := tx.Rollback()
			if errRb != nil {
				return errRb
			}
			return err
		}
		operation.Id = lastId
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Find ...
func (r *OperationRepository) Find(id int) (*model.Operation, error) {
	u := &model.Operation{}
	if err := r.store.db.QueryRowx(
		"SELECT id, id_account, time, sum, description FROM operation WHERE id = $1",
		id,
	).StructScan(u); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil
}
func (r *OperationRepository) Get(dateFrom, dateTo time.Time, idUser int) ([]model.Operation, error) {
	c := model.Operation{}
	res := make([]model.Operation, 0)
	rows, err := r.store.db.Queryx("SELECT id, id_account, time, sum, description FROM operation WHERE time BETWEEN $1 AND $2", dateFrom, dateTo)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, nil
}

func (r *OperationRepository) Update(o *model.Operation) error {
	_, err := r.Find(o.Id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	err = r.store.db.QueryRow(
		"UPDATE operation SET(time, sum, id_category,id_account, description) = (SELECT $2,$3,$4,$5,$6) WHERE id = $1",
		o.Id,
		o.Time,
		o.Sum,
		o.Category.Id,
		o.IdAccount,
		o.Description,
	).Err()
	if err != nil {
		return err
	}

	return nil
}
func (r *OperationRepository) Delete(id int) error {
	_, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	err = r.store.db.QueryRow("DELETE FROM operation WHERE id = $1;", id).Err()
	if err != nil {
		return err
	}
	return nil
}
