package infra

import (
	"fmt"
	"go-dependency-injector/pkg/dependencies"
	"math/rand/v2"
)

type Cache struct {
	dummy string
}

type cacheSingletonKey struct{}

var CacheSingletonKey = cacheSingletonKey{}

func NewCache() *Cache {

	return &Cache{
		dummy: fmt.Sprint(rand.IntN(100)),
	}
}

func (d *Cache) Initialize(_ *dependencies.DependencyManager) error {
	return nil
}

func (d *Cache) Print() {
	fmt.Println("Cache singleton")
	fmt.Println("    My address: ", d)
	return
}
