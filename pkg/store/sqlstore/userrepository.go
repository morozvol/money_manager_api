package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}
type user struct {
	Id                int           `db:"id"`
	Name              string        `db:"name"`
	DefaultCurrencyId sql.NullInt64 `db:"id_default_currency"`
}

func (u user) toModel() model.User {
	return model.User{
		Id:                u.Id,
		Name:              u.Name,
		DefaultCurrencyId: int(u.DefaultCurrencyId.Int64),
	}
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {

	var id sql.NullInt64
	if u.DefaultCurrencyId != 0 {
		id.Int64 = int64(u.DefaultCurrencyId)
		id.Valid = true
	}
	return r.store.db.QueryRowx(
		"INSERT INTO \"user\" (id, name, id_default_currency) VALUES ($1,$2,$3)",
		u.Id,
		u.Name,
		id,
	).Err()
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &user{}
	if err := r.store.db.QueryRowx(
		"SELECT id ,name, id_default_currency FROM \"user\" WHERE id = $1",
		id,
	).StructScan(u); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	res := u.toModel()
	return &res, nil
}
