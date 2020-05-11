package main

import (
	"fmt"
	"sync"
)

// パッケージ外から呼び出されないようにlowercase
type singletonDB struct {
	name string
}

func (db *singletonDB) GetDBInfo() string {
	return db.name
}

var once sync.Once
var instance *singletonDB

func GetSingletonDatabase(dbName string) *singletonDB {
	// once.Doで1回のみ実行できる関数を作成できる
	once.Do(func() {
		db := singletonDB{dbName}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase("PostgreSQL")
	dbName := db.GetDBInfo()
	fmt.Println("DB Name: ", dbName)

	db2 := GetSingletonDatabase("MySQL")
	dbName2 := db2.GetDBInfo()
	fmt.Println("DB Name: ", dbName2)
}
