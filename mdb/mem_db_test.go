package mdb_test

import (
	"fmt"
	"github.com/scim2/example-server/mdb"
)

func ExampleNew() {
	db := mdb.New()
	_ = db.Update(func(tx *mdb.TX) error {
		tx.Set("key1", "hello")
		tx.Set("key2", "world")
		tx.Set("key3", "!")
		return nil
	})
	_ = db.View(func(tx *mdb.TX) error {
		fmt.Println(tx.Get("key1"))
		fmt.Println(tx.Get("key2"))
		fmt.Println(tx.Get("key3"))
		return nil
	})

	// Output:
	// hello true
	// world true
	// ! true
}
