package mock

import "github.com/gochallenge/gochallenge/model"

// Users represents users collection, mocked out as in-memory map
type Users struct {
	index       map[model.UserID]*model.User
	indexAPIKey map[string]*model.User
}

// NewUsers returns a new initialised users collection.
func NewUsers() Users {
	return Users{
		index:       make(map[model.UserID]*model.User),
		indexAPIKey: make(map[string]*model.User),
	}
}

// Add user to the mock users.
func (us *Users) Add(u *model.User) error {
	if _, ok := us.index[u.ID]; ok {
		return model.ErrDuplicateRecord
	}
	us.index[u.ID] = u
	us.indexAPIKey[u.APIKey] = u

	return nil
}

// Find searches for a user in the collection by its id.
func (us *Users) Find(id model.UserID) (*model.User, error) {
	var (
		u  *model.User
		ok bool
	)

	if u, ok = us.index[id]; !ok {
		return nil, model.ErrNotFound
	}
	return u, nil
}

// FindByAPIKey finds a user in the collection by its API Key.
func (us *Users) FindByAPIKey(key string) (*model.User, error) {
	var (
		u  *model.User
		ok bool
	)

	if u, ok = us.indexAPIKey[key]; !ok {
		return nil, model.ErrNotFound
	}
	return u, nil
}

// FindByGithubID finds a user in the collection by its Github ID
func (us *Users) FindByGithubID(id int) (*model.User, error) {
	for _, u := range us.index {
		if u.GithubID == id {
			return u, nil
		}
	}
	return nil, model.ErrNotFound
}
