package main

import (
	"go-dependency-injector/pkg/dependencies"
	"go-dependency-injector/pkg/infra"
	"go-dependency-injector/pkg/services"
)

func main() {
	// Initialize the dependency manager
	dm := dependencies.NewDependencyManager()

	// Register dependencies
	dm.Register(infra.CacheSingletonKey, infra.NewCache())
	dm.Register(services.ServiceBSingletonKey, services.NewServiceB())
	dm.Register(services.ServiceASingletonKey, services.NewServiceA())
	dm.Register(infra.DataBaseSingletonKey, infra.NewDatabase())

	// Initialize all dependencies
	if err := dm.InitializeAll(); err != nil {
		panic(err)
	}

	// Now you can use the initialized services
	serviceA, _ := dependencies.GetSingleton[*services.ServiceA](dm, services.ServiceASingletonKey)
	serviceB, _ := dependencies.GetSingleton[*services.ServiceB](dm, services.ServiceBSingletonKey)

	serviceA.Print()
	serviceB.Print()

}
