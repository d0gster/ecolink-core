package database

import (
	"ecolink-core/internal/models"
	"errors"
	"sync"
)

// MemoryDB implementa um banco em memória para demonstração
type MemoryDB struct {
	links map[string]*models.Link
	codes map[string]*models.Link
	mutex sync.RWMutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		links: make(map[string]*models.Link),
		codes: make(map[string]*models.Link),
	}
}

func (db *MemoryDB) SaveLink(link *models.Link) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	
	db.links[link.ID] = link
	db.codes[link.ShortCode] = link
	return nil
}

func (db *MemoryDB) GetLinkByCode(code string) (*models.Link, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	
	link, exists := db.codes[code]
	if !exists {
		return nil, errors.New("link not found")
	}
	return link, nil
}

func (db *MemoryDB) GetLinksByUser(userID string) ([]*models.Link, error) {
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

func (db *MemoryDB) IncrementClick(code string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	
	if link, exists := db.codes[code]; exists {
		link.ClickCount++
		return nil
	}
	return errors.New("link not found")
}