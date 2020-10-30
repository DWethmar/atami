package message

// CreatorRepository defines a messsage listing repository
type DeleterRepository interface {
	Delete(ID ID) error
}

// Creator creates messages.
type Deleter struct {
	deleteRepo DeleterRepository
}

// Create a new message
func (m *Deleter) Delete(ID ID) error {
	return m.deleteRepo.Delete(ID)
}

// NewDeleter returns a new Listing
func NewDeleter(r CreatorRepository) *Creator {
	return &Creator{r}
}
