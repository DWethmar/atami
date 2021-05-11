package user

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/entity"
)

type postgresRepo struct {
	db database.Transaction
}

//NewPostgresRepository create new repository
func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepo{
		db: db,
	}
}

func (r *postgresRepo) GetByUID(UID entity.UID) (*User, error) {
	return nil, nil
}

func (r *postgresRepo) Get(ID entity.ID) (*User, error) {
	return nil, nil
}

func (r *postgresRepo) List(limit, offset uint) ([]*User, error) {
	return nil, nil
}

func (r *postgresRepo) Create(e *User) (entity.ID, error){
	return 0, nil
}

func (r *postgresRepo) Update(e *User) error{
	return nil
}

func (r *postgresRepo) Delete(ID entity.ID) error{
	return nil
}

func defaultMap(row Row) (*User, error) {
	e := &User{}

	var biography sql.NullString

	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&biography,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		return nil, err
	}

	if biography.Valid {
		e.Biography = biography.String
	}

	return e, nil
}

func mapIsUniqueCheck(row Row) (bool, error) {
	i := 0
	row.Scan(&i)
	return i == 0, row.Err()
}

func mapWithPassword(row Row) (*User, error) {
	e := &User{}

	var biography sql.NullString

	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&biography,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.Password,
	); err != nil {
		return nil, err
	}

	if biography.Valid {
		e.Biography = biography.String
	}

	return e, nil
}
