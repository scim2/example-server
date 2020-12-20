package mdb

import (
	"github.com/elimity-com/scim"
	"sort"
	"sync"
)

// DB is a simple key-value pair in memory.
//
// WARNING: DO NOT USE THIS IN PRODUCTION.
type DB struct {
	lock *sync.RWMutex // mutex to lock our data.
	data map[string]Instance
}

// Instance is a wrapper for a db value instance and their corresponding meta data.
type Instance struct {
	Value interface{}
	Meta  scim.Meta
}

// New returns a new in memory DB.
func New() *DB {
	return &DB{
		lock: new(sync.RWMutex),
		data: make(map[string]Instance),
	}
}

// TX represents a structure that holds the current transaction data.
type TX struct {
	db       *DB  // Reference to the underlying database.
	writable bool // Whether the transaction is for reading or writing.
}

func (tx *TX) Set(key string, value Instance) {
	if !tx.writable {
		panic("Called tx.Set() on an unwritable transaction.")
	}
	tx.db.data[key] = value
}

func (tx *TX) Delete(key string) bool {
	if !tx.writable {
		panic("Called tx.Delete() on an unwritable transaction.")
	}
	if _, ok := tx.db.data[key]; !ok {
		return false
	}
	delete(tx.db.data, key)
	return true
}

func (tx *TX) Get(key string) (Instance, bool) {
	v, ok := tx.db.data[key]
	return v, ok
}

func (tx *TX) GetAll() []Instance {
	keys := make([]string, 0)
	for k, _ := range tx.db.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var (
		i         = 0
		instances = make([]Instance, len(tx.db.data))
	)
	for _, v := range keys {
		instances[i] = tx.db.data[v]
		i++
	}
	return instances
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
