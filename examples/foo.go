package main

import (
	"fmt"
	"log"
)
import "github.com/tidwall/buntdb"

func main() {
	db, err := buntdb.Open("examples/bar.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("hello", "world", nil)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("hello")
		if err != nil {
			return err
		}

		fmt.Printf("value: %s\n", val)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
