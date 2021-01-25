package user

// UpdaterRepository declares a storage repository
type UpdaterRepository interface {
	Update(updateUser UpdateAction) (*User, error)
}

// Updater struct declaration
type Updater struct {
	validator   *Validator
	updaterRepo UpdaterRepository
}

// Update updates a new user
func (m *Updater) Update(updateUser UpdateRequest) (*User, error) {
	if err := m.validator.ValidateUpdateUser(updateUser); err != nil {
		return nil, err
	}
	return m.updaterRepo.Update(UpdateAction{
		Biography: updateUser.Biography,
	})
}

// NewUpdater returns a new searcher
func NewUpdater(r UpdaterRepository) *Updater {
	return &Updater{
		updaterRepo: r,
		validator:   NewDefaultValidator(),
	}
}
