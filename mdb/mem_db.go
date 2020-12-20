package mdb

import (
	"sync"
)

// DB is a simple key-value pair in memory.
//
// WARNING: DO NOT USE THIS IN PRODUCTION.
type DB struct {
	lock *sync.RWMutex // mutex to lock our data.
	data map[string]interface{}
}

// New returns a new in memory DB.
func New() *DB {
	return &DB{
		lock: new(sync.RWMutex),
		data: make(map[string]interface{}),
	}
}

// TX represents a structure that holds the current transaction data.
type TX struct {
	db       *DB  // Reference to the underlying database.
	writable bool // Whether the transaction is for reading or writing.
}

func (tx *TX) Set(key, value string) {
	if !tx.writable {
		panic("Called tx.Set() on an unwritable transaction.")
	}
	tx.db.data[key] = value
}

func (tx *TX) Get(key string) (interface{}, bool) {
	v, ok := tx.db.data[key]
	return v, ok
}

// Begin opens a TX.
func (db *DB) Begin(writable bool) (*TX, error) {
	tx := &TX{
		db:       db,
		writable: writable,
	}
	tx.lock()
	return tx, nil
}

// View opens and manages a read transaction.
func (db *DB) View(fn func(tx *TX) error) error {
	return db.managed(false, fn)
}

// Update opens and manages a write transaction.
func (db *DB) Update(fn func(tx *TX) error) error {
	return db.managed(true, fn)
}

func (db *DB) managed(writable bool, fn func(tx *TX) error) error {
	tx, err := db.Begin(writable)
	if err != nil {
		return err
	}
	defer tx.unlock()
	return fn(tx)
}

func (tx *TX) lock() {
	if tx.writable {
		tx.db.lock.Lock()
	} else {
		tx.db.lock.RLock()
	}
}

func (tx *TX) unlock() {
	if tx.writable {
		tx.db.lock.Unlock()
	} else {
		tx.db.lock.RUnlock()
	}
}
