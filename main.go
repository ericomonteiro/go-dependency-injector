package main

import (
	"fmt"
	"go-dependency-injector/pkg/dependencies"
	"go-dependency-injector/pkg/infra"
	"go-dependency-injector/pkg/services"
)

func main() {
	// Initialize the dependency manager
	dm := dependencies.NewDependencyManager()

	// Register dependencies
	dm.Register(infra.NewCache())
	dm.Register(services.NewServiceB())
	dm.Register(services.NewServiceA())
	dm.Register(infra.NewDatabase())

	// Auto-wire dependencies using types (does not support multiple implementations of the same type)
	err := dm.Autowire()
	if err != nil {
		panic(err)
	}

	dm.GenerateDependencyGraph()
	fmt.Println("")

	// Now you can use the initialized services
	// You can get singleton by key
	serviceA, err := dependencies.GetSingletonByKey[*services.ServiceA](dm, services.ServiceASingletonKey)
	if err != nil {
		panic(err)
	}

	serviceB, err := dependencies.GetSingletonByKey[*services.ServiceB](dm, services.ServiceBSingletonKey)
	if err != nil {
		panic(err)
	}

	serviceA.Print()
	serviceB.Print()

	// You can also use the interface
	serviceA.ServiceB.PrintServiceB()

}
