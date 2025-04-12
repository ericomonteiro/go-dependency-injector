package infra

import (
	"fmt"
	"go-dependency-injector/pkg/dependencies"
)

type Database struct {
	dummy float64
}

type dataBaseSingletonKey struct{}

var DataBaseSingletonKey = dataBaseSingletonKey{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Initialize(_ *dependencies.DependencyManager) error {
	return nil
}

func (d *Database) Print() {
	fmt.Println("Database singleton")
	fmt.Println("    My address: ", d)
	return
}
