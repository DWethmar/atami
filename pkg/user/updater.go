package user

import "time"

// UpdaterRepository declares a storage repository
type UpdaterRepository interface {
	Update(ID int, action UpdateAction) (*User, error)
}

// Updater struct declaration
type Updater struct {
	validator   *Validator
	updaterRepo UpdaterRepository
}

// Update updates a new user
func (m *Updater) Update(ID int, request UpdateRequest) (*User, error) {
	if err := m.validator.ValidateUpdateUser(request); err != nil {
		return nil, err
	}
	return m.updaterRepo.Update(ID, UpdateAction{
		Biography: request.Biography,
		UpdatedAt: time.Now().UTC(),
	})
}

// NewUpdater returns a new searcher
func NewUpdater(r UpdaterRepository) *Updater {
	return &Updater{
		updaterRepo: r,
		validator:   NewDefaultValidator(),
	}
}
