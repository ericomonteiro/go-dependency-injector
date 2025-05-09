package infra

import (
	"fmt"
	"math/rand/v2"
)

const DataBaseSingletonKey = "DataBaseSingletonKey"

type Database struct {
	dummy float64
}

func (d *Database) Key() string {
	return DataBaseSingletonKey
}

func NewDatabase() *Database {
	return &Database{
		dummy: rand.Float64(),
	}
}

func (d *Database) Print() {
	fmt.Println("Database singleton")
	fmt.Println("    My address: ", d)
	return
}
