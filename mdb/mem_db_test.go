package mdb_test

import (
	"fmt"
	"github.com/scim2/example-server/mdb"
)

func ExampleNew() {
	db := mdb.New()
	_ = db.Update(func(tx *mdb.TX) error {
		tx.Set("key1", mdb.Instance{
			Value: "hello",
		})
		tx.Set("key2", mdb.Instance{
			Value: "world",
		})
		tx.Set("key3", mdb.Instance{
			Value: "!",
		})
		return nil
	})
	_ = db.View(func(tx *mdb.TX) error {
		fmt.Println(tx.Get("key1"))
		fmt.Println(tx.Get("key2"))
		fmt.Println(tx.Get("key3"))
		return nil
	})

	// Output:
	// {hello {<nil> <nil> }} true
	// {world {<nil> <nil> }} true
	// {! {<nil> <nil> }} true
}
