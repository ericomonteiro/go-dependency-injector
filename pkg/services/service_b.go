package services

import (
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
	database, err := dependencies.GetSingleton[*infra.Database](dm, infra.DataBaseSingletonKey)
	if err != nil {
		return err
	}

	cache, err := dependencies.GetSingleton[*infra.Cache](dm, infra.CacheSingletonKey)
	if err != nil {
		return err
	}

	d.Database = database
	d.Cache = cache

	return nil
}

func (d *ServiceB) Print() {
	println("Service B singleton")
	println("    My address: ", d)
	println("    Database address: ", d.Database)
	println("    Cache address: ", d.Cache)
	return
}
