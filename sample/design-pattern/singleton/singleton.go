package main

import (
	"fmt"
	"sync"

	"github.com/ono5/money-boy/sample/design-pattern/utils"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once init() -- thread safety
// laziness

var once sync.Once
var instance *singletonDatabase

// GetSingletonDatabase - gets single instance
func GetSingletonDatabase(filePath string) *singletonDatabase {
	once.Do(func() {
		fmt.Println("Do at once!")
		caps, err := utils.ReadData(filePath)
		db := singletonDatabase{}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

// GetTotalPopulation - returns total population num
func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase("/capitals.txt").GetPopulation(city)
	}
	return result
}

// GetTotalPopulationEx - returns total population num
func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

// DummyDatabase - difines dummy database
type DummyDatabase struct {
	dummyData map[string]int
}

// GetPopulation - retuns dummy data
func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	// db := GetSingletonDatabase("/capitals.txt")
	// pop := db.GetPopulation("Tokyo")
	// fmt.Println("Pop of Tokyo = ", pop)

	// // No get new instance
	// db2 := GetSingletonDatabase("test.txt")
	// pop2 := db2.GetPopulation("Tokyo")
	// fmt.Println("Pop of Tokyo = ", pop2)

	// cities := []string{"Tokyo", "New York"}
	// tp := GetTotalPopulation(cities)
	// ok := tp == (33200000 + 17800000)
	// fmt.Println(ok)

	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
