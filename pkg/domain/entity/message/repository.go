package message

import "github.com/dwethmar/atami/pkg/domain/entity"

// Reader allows read operations
type Reader interface {
	GetByUID(UID entity.UID) (*Message, error)
	Get(ID entity.ID) (*Message, error)
	List(limit, offset uint) ([]*Message, error)
}

// Writer allows write operations
type Writer interface {
	Create(m Create) (entity.ID, error)
	Update(ID entity.ID, u Update) error
	Delete(ID entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
