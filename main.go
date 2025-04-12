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
	dm.Register(infra.CacheSingletonKey, infra.NewCache())
	dm.Register(services.ServiceBSingletonKey, services.NewServiceB())
	dm.Register(services.ServiceASingletonKey, services.NewServiceA())
	dm.Register(infra.DataBaseSingletonKey, infra.NewDatabase())

	// Wire dependencies by initialize method of each singleton
	//if err := dm.InitializeAll(); err != nil {
	//	panic(err)
	//}

	// Auto-wire dependencies using types (does not support multiple implementations of the same type)
	err := dm.AutoWire()
	if err != nil {
		panic(err)
	}

	dm.GenerateDependencyGraph()
	fmt.Println("")

	// Now you can use the initialized services
	// You can get singleton by type
	serviceA := dependencies.GetSingletonByType[*services.ServiceA](dm)

	// You can get singleton by key (in case you have multiple implementations)
	serviceB, _ := dependencies.GetSingletonByKey[*services.ServiceB](dm, services.ServiceBSingletonKey)

	serviceA.Print()
	serviceB.Print()

}
