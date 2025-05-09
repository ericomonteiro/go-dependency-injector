package services

import (
	"fmt"
	"go-dependency-injector/pkg/infra"
)

type ServiceB struct {
	Database *infra.Database `inject:"DataBaseSingletonKey"`
	Cache    *infra.Cache    `inject:"CacheSingletonKey"`
}

func (d *ServiceB) Key() string {
	return "ServiceBKey"
}

type serviceBSingletonKey struct{}

var ServiceBSingletonKey = serviceBSingletonKey{}

func NewServiceB() *ServiceB {
	return &ServiceB{}
}

func (d *ServiceB) Print() {
	fmt.Println("Service B singleton")
	fmt.Println("    My address: ", d)
	fmt.Println("    Database address: ", d.Database)
	fmt.Println("    Cache address: ", d.Cache)
	return
}
