package user

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain"
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
	u, err := queryRowSelectUserByUID(r.db, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *postgresRepo) GetByEmail(email string) (*User, error) {
	u, err := queryRowSelectUserByEmail(r.db, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *postgresRepo) GetByUsername(username string) (*User, error) {
	u, err := queryRowSelectUserByUsername(r.db, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *postgresRepo) GetCredentials(email string) (*UserCredentials, error) {
	u, err := queryRowSelectUserCredentials(r.db, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *postgresRepo) Get(ID entity.ID) (*User, error) {
	u, err := queryRowSelectUserByID(r.db, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *postgresRepo) List(limit, offset uint) ([]*User, error) {
	return querySelectUsers(r.db, limit, offset)
}

func (r *postgresRepo) Create(e *User) (entity.ID, error){
	user, err := queryRowInsertUser(r.db, e.UID, e.Username, e.Biography, e.Email, e.Password, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *postgresRepo) Update(e *User) error{
	_, err := queryRowUpdateUser(r.db, e.ID, e.Biography, e.UpdatedAt)
	return err
}

func (r *postgresRepo) Delete(ID entity.ID) error{
	result, err := execDeleteUser(r.db, ID) 
	if err != nil{
		return err
	}
	if e, err := result.RowsAffected(); err == nil {
		if e != 0 {
			return domain.ErrCannotBeDeleted
		} 
	} else {
		return err
	}

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

	e.CreatedAt = entity.SetDefaultTimePrecision(e.CreatedAt)
	e.UpdatedAt = entity.SetDefaultTimePrecision(e.UpdatedAt)

	return e, nil
}

func mapIsUniqueCheck(row Row) (bool, error) {
	i := 0
	row.Scan(&i)
	return i == 0, row.Err()
}

// mapCredentials creates a UserCredentials
func mapCredentials(row Row) (*UserCredentials, error) {
	e := &UserCredentials{}

	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Username,
		&e.Email,
		&e.Password,
	); err != nil {
		return nil, err
	}
	
	return e, nil
}
