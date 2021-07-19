package message

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

func (r *postgresRepo) Get(ID entity.ID) (*Message, error) {
	m, err := queryRowSelectMessageByID(r.db, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return m, nil
}

func (r *postgresRepo) GetByUID(UID entity.UID) (*Message, error) {
	m, err := queryRowSelectMessageByUID(r.db, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return m, nil
}

func (r *postgresRepo) List(limit, offset uint) ([]*Message, error) {
	return querySelectMessages(r.db, limit, offset)
}

func (r *postgresRepo) Update(m *Message) error {
	result, err := execUpdateUser(r.db, m.ID, m.Text, m.UpdatedAt)
	if err != nil {
		return err
	}

	if a, err := result.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return domain.ErrCannotBeUpdated
	}

	return err
}

func (r *postgresRepo) Create(message *Message) (entity.ID, error) {
	ID, err := queryRowInsertMessage(
		r.db,
		message.UID,
		message.Text,
		message.CreatedByUserID,
		message.CreatedAt,
		message.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func (r *postgresRepo) Delete(ID entity.ID) error {
	result, err := execDeleteMessage(r.db, ID)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return domain.ErrCannotBeDeleted
	}

	return nil
}

func insertRowMap(row Row) (entity.ID, error) {
	var ID entity.ID
	if err := row.Scan(
		&ID,
	); err != nil {
		return 0, err
	}
	return ID, nil
}

func messageWithUserRowMap(row Row) (*Message, error) {
	e := &Message{
		CreatedBy: User{},
	}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Text,
		&e.CreatedByUserID,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.CreatedBy.ID,
		&e.CreatedBy.UID,
		&e.CreatedBy.Username,
	); err != nil {
		return nil, err
	}
	e.CreatedAt = entity.SetDefaultTimePrecision(e.CreatedAt)
	e.UpdatedAt = entity.SetDefaultTimePrecision(e.UpdatedAt)
	return e, nil
}
