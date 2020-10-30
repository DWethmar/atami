package message

// DeleterRepository deletes messsages
type DeleterRepository interface {
	Delete(ID ID) error
}

// Deleter deletes messages.
type Deleter struct {
	deleteRepo DeleterRepository
}

// Delete a message
func (m *Deleter) Delete(ID ID) error {
	return m.deleteRepo.Delete(ID)
}

// NewDeleter returns a new Listing
func NewDeleter(r DeleterRepository) *Deleter {
	return &Deleter{r}
}
