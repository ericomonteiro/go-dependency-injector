package services

import (
	"fmt"
	"go-dependency-injector/pkg/infra"
)

type IAmServiceB interface {
	PrintServiceB()
	Key() string
}

const ServiceBSingletonKey = "ServiceBSingletonKey"

type ServiceB struct {
	Database *infra.Database `inject:"DataBaseSingletonKey"`
	Cache    *infra.Cache    `inject:"CacheSingletonKey"`
}

func (d *ServiceB) Key() string {
	return ServiceBSingletonKey
}

func NewServiceB() IAmServiceB {
	return &ServiceB{}
}

func (d *ServiceB) PrintServiceB() {
	fmt.Println("I am a service B")
}

func (d *ServiceB) Print() {
	fmt.Println("Service B singleton")
	fmt.Println("    My address: ", d)
	fmt.Println("    Database address: ", d.Database)
	fmt.Println("    Cache address: ", d.Cache)
	return
}
