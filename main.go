package main

import (
	"fmt"
	dbm "github.com/tendermint/tmlibs/db"
)

var (
	LevelDBDir = "./data"
	Key        = "my_data"
)

func main() {
	db := dbm.NewDB("test", "leveldb", LevelDBDir)
	defer db.Close()

	db.Set([]byte(Key), []byte("This is a test."))

	value := db.Get([]byte(Key))
	if value == nil {
		return
	}
	fmt.Printf("key:%v, value:%v\n", Key, string(value))

	db.Delete([]byte(Key))
}




