package infra

import (
	"fmt"
	"math/rand/v2"
)

const CacheSingletonKey = "CacheSingletonKey"

type Cache struct {
	dummy string
}

func (c *Cache) Key() string {
	return CacheSingletonKey
}

func NewCache() *Cache {

	return &Cache{
		dummy: fmt.Sprint(rand.IntN(100)),
	}
}

func (c *Cache) Print() {
	fmt.Println("Cache singleton")
	fmt.Println("    My address: ", c)
	return
}
