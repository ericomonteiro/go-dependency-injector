package services

import (
	"fmt"
)

const ServiceASingletonKey = "ServiceAKey"

type ServiceA struct {
	ServiceB *ServiceB `inject:"ServiceBKey"`
}

func (d *ServiceA) Key() string {
	return ServiceASingletonKey
}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (d *ServiceA) Print() {
	fmt.Println("Service A singleton")
	fmt.Println("    My address: ", d)
	fmt.Println("    Service B address: ", d.ServiceB)
	fmt.Println("    Service B database: ", d.ServiceB.Database)
	fmt.Println("    Service B cache: ", d.ServiceB.Cache)
	return
}
