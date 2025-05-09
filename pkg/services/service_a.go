package services

import (
	"fmt"
)

const ServiceASingletonKey = "ServiceASingletonKey"

type ServiceA struct {
	ServiceB IAmServiceB `inject:"ServiceBSingletonKey"`
}

func (d *ServiceA) Key() string {
	return ServiceASingletonKey
}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (d *ServiceA) Print() {
	serviceB, _ := d.ServiceB.(*ServiceB)

	fmt.Println("Service A singleton")
	fmt.Println("    My address: ", d)
	fmt.Println("    Service B address: ", d.ServiceB)
	fmt.Println("    Service B database: ", serviceB.Database)
	fmt.Println("    Service B cache: ", serviceB.Cache)
	return
}
