package database

import (
	"context"
	"encoding/json"
	"sync"
)

// Database does ... (TODO)
type Database struct {
	mu   sync.RWMutex
	data map[string]map[int]json.RawMessage
}

// New does ... (TODO)
func New() *Database {
	return &Database{
		data: make(map[string]map[int]json.RawMessage),
	}
}

// NewID does ... (TODO)
func (db *Database) NewID(ctx context.Context, model string) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var newID int
	for {
		newID++
		_, ok := db.data[model][newID]
		if !ok {
			break
		}
	}
	return newID, nil
}

// GetObjects does ... (TODO)
func (db *Database) GetObjects(ctx context.Context, model string) (map[int]json.RawMessage, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return db.data[model], nil
}

// UpdateObject does ... (TODO)
func (db *Database) UpdateObject(ctx context.Context, model string, id int, data json.RawMessage) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if ids := db.data[model]; ids == nil {
		db.data[model] = make(map[int]json.RawMessage)
	}

	db.data[model][id] = data
	return nil
}
