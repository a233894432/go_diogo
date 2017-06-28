package main

import (
	"fmt"

	"github.com/couchbase/moss"
)

func main() {
	c, err := moss.NewCollection(moss.CollectionOptions{})
	defer c.Close()

	batch, err := c.NewBatch(0, 0)
	defer batch.Close()

	batch.Set([]byte("car-0"), []byte("tesla"))
	batch.Set([]byte("car-1"), []byte("honda"))

	err = c.ExecuteBatch(batch, moss.WriteOptions{})

	ss, err := c.Snapshot()
	defer ss.Close()

	ropts := moss.ReadOptions{}

	val0, err := ss.Get([]byte("car-0"), ropts)         // val0 == []byte("tesla").
	valX, err := ss.Get([]byte("car-not-there"), ropts) // valX == nil.

	fmt.Println(val0)
	// A Get can also be issued directly against the collection
	val1, err := c.Get([]byte("car-1"), ropts) // val1 == []byte("honda")
}
