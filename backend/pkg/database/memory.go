package database

import (
	"ecolink-core/internal/models"
	"errors"
	"sync"
)

type MemoryDB struct {
	links map[string]*models.Link
	users map[string]*models.User
	mutex sync.RWMutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		links: make(map[string]*models.Link),
		users: make(map[string]*models.User),
	}
}

func (db *MemoryDB) SaveLink(link *models.Link) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.links[link.Code] = link
	return nil
}

func (db *MemoryDB) GetLink(code string) (*models.Link, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	link, exists := db.links[code]
	if !exists {
		return nil, errors.New("link not found")
	}
	return link, nil
}

func (db *MemoryDB) GetUserLinks(userID string) ([]*models.Link, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var userLinks []*models.Link
	for _, link := range db.links {
		if link.UserID == userID {
			userLinks = append(userLinks, link)
		}
	}
	return userLinks, nil
}

func (db *MemoryDB) IncrementClicks(code string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if link, exists := db.links[code]; exists {
		link.Clicks++
		return nil
	}
	return errors.New("link not found")
}

func (db *MemoryDB) DeleteLink(code string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.links[code]; exists {
		delete(db.links, code)
		return nil
	}
	return errors.New("link not found")
}

func (db *MemoryDB) SaveUser(user *models.User) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.users[user.ID] = user
	return nil
}

func (db *MemoryDB) GetUser(id string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (db *MemoryDB) GetUserByGoogleID(googleID string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	for _, user := range db.users {
		if user.GoogleID == googleID {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
