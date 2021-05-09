package message

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/domain/message"
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

func (r *postgresRepo) Update(message *Message) error {
	return nil
}

func (r *postgresRepo) Create(message *Message) (entity.ID, error) {
	msg, err := queryRowInsertMessage(
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
	return msg, nil
}

func (r *postgresRepo) Delete(ID entity.ID) error {
	result, err := execDeleteMessage(r.db, ID)
	if err != nil {
		return err
	}

	if a, err := result.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return message.ErrCouldNotDelete
	}

	return nil
}
