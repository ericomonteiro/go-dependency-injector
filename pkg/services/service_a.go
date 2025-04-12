package services

import (
	"go-dependency-injector/pkg/dependencies"
)

type ServiceA struct {
	serviceB *ServiceB
}

type serviceASingletonKey struct{}

var ServiceASingletonKey = serviceASingletonKey{}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (d *ServiceA) Initialize(dm *dependencies.DependencyManager) error {
	serviceB, err := dependencies.GetSingleton[*ServiceB](dm, ServiceBSingletonKey)
	if err != nil {
		return nil
	}

	d.serviceB = serviceB
	return nil
}

func (d *ServiceA) Print() {
	println("Service A singleton")
	println("    My address: ", d)
	println("    Service B address: ", d.serviceB)
	println("    Service B database: ", d.serviceB.Database)
	println("    Service B cache: ", d.serviceB.Cache)
	return
}
