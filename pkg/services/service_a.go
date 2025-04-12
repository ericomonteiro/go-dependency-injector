package services

import (
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
	println("Service A singleton")
	println("    My address: ", d)
	println("    Service B address: ", d.ServiceB)
	println("    Service B database: ", d.ServiceB.Database)
	println("    Service B cache: ", d.ServiceB.Cache)
	return
}
