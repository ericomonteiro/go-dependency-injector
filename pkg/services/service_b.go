package services

import (
	"fmt"
	"go-dependency-injector/pkg/dependencies"
	"go-dependency-injector/pkg/infra"
)

type ServiceB struct {
	Database *infra.Database
	Cache    *infra.Cache
}

type serviceBSingletonKey struct{}

var ServiceBSingletonKey = serviceBSingletonKey{}

func NewServiceB() *ServiceB {
	return &ServiceB{}
}

func (d *ServiceB) Initialize(dm *dependencies.DependencyManager) error {
	database, err := dependencies.GetSingletonByKey[*infra.Database](dm, infra.DataBaseSingletonKey)
	if err != nil {
		return err
	}

	cache, err := dependencies.GetSingletonByKey[*infra.Cache](dm, infra.CacheSingletonKey)
	if err != nil {
		return err
	}

	d.Database = database
	d.Cache = cache

	return nil
}

func (d *ServiceB) Print() {
	fmt.Println("Service B singleton")
	fmt.Println("    My address: ", d)
	fmt.Println("    Database address: ", d.Database)
	fmt.Println("    Cache address: ", d.Cache)
	return
}
