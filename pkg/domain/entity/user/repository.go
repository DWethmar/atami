package user

import "github.com/dwethmar/atami/pkg/domain/entity"

// Reader allows read operations
type Reader interface {
	GetByUID(UID entity.UID) (*User, error)
	Get(ID entity.ID) (*User, error)
	List(limit, offset uint) ([]*User, error)
}

// Writer allows write operations
type Writer interface {
	Create(e *User) (entity.ID, error)
	Update(e *User) error
	Delete(ID entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}