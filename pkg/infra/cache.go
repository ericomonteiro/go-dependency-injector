package infra

import (
	"fmt"
	"go-dependency-injector/pkg/dependencies"
)

type Cache struct {
	dummy string
}

type cacheSingletonKey struct{}

var CacheSingletonKey = cacheSingletonKey{}

func NewCache() *Cache {
	return &Cache{}
}

func (d *Cache) Initialize(_ *dependencies.DependencyManager) error {
	return nil
}

func (d *Cache) Print() {
	fmt.Println("Cache singleton")
	fmt.Println("    My address: ", d)
	return
}
