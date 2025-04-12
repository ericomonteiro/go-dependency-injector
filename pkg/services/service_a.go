package services

import (
	"fmt"
	"go-dependency-injector/pkg/dependencies"
)

type ServiceA struct {
	ServiceB *ServiceB
}

type serviceASingletonKey struct{}

var ServiceASingletonKey = serviceASingletonKey{}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (d *ServiceA) Initialize(dm *dependencies.DependencyManager) error {
	serviceB, err := dependencies.GetSingletonByKey[*ServiceB](dm, ServiceBSingletonKey)
	if err != nil {
		return nil
	}

	d.ServiceB = serviceB
	return nil
}

func (d *ServiceA) Print() {
	fmt.Println("Service A singleton")
	fmt.Println("    My address: ", d)
	fmt.Println("    Service B address: ", d.ServiceB)
	fmt.Println("    Service B database: ", d.ServiceB.Database)
	fmt.Println("    Service B cache: ", d.ServiceB.Cache)
	return
}
