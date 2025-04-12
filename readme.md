# Go Dependency Injector

This project is a proposal to solve dependency injection in Go. It provides a simple and extensible approach to manage dependencies using a `DependencyManager` that supports singleton registration, initialization, and retrieval.

## Overview

Dependency injection is a design pattern that allows you to decouple the creation of dependencies from their usage. In this project, we implement a lightweight dependency injection framework for Go, enabling you to register and initialize singletons and retrieve them when needed.

### Key Features

- **Singleton Management**: Register and manage singleton instances.
- **Initialization**: Automatically initialize all registered dependencies.
- **Type-Safe Retrieval**: Retrieve dependencies with type safety using generics.

## Approach

The project uses a `DependencyManager` to handle the lifecycle of dependencies. Each dependency implements the `Singleton` interface, which requires an `Initialize` method. Dependencies are registered with unique keys, and the `DependencyManager` ensures they are initialized and accessible throughout the application.

### Core Components

1. **`DependencyManager`**: Manages the registration, initialization, and retrieval of dependencies.
2. **`Singleton` Interface**: Defines the contract for dependencies that require initialization.
3. **Type-Safe Retrieval**: Uses Go generics to safely retrieve dependencies by their type.

## How to Use

### 1. Define Your Dependencies

Each dependency should implement the `Singleton` interface. For example:

```go
type Cache struct {
    dummy string
}

func (c *Cache) Initialize(_ *dependencies.DependencyManager) error {
    // Initialization logic
    return nil
}
```

### 2. Register Dependencies

In your `main.go`, register all dependencies with the `DependencyManager`:

```go
dm := dependencies.NewDependencyManager()

dm.Register(infra.CacheSingletonKey, infra.NewCache())
dm.Register(infra.DataBaseSingletonKey, infra.NewDatabase())
dm.Register(services.ServiceASingletonKey, services.NewServiceA())
dm.Register(services.ServiceBSingletonKey, services.NewServiceB())
```

### 3. Initialize All Dependencies

Call `InitializeAll` to initialize all registered dependencies:

```go
if err := dm.InitializeAll(); err != nil {
    panic(err)
}
```

### 4. Retrieve Dependencies

Use the `GetSingleton` function to retrieve dependencies safely:

```go
serviceA, _ := dependencies.GetSingleton[*services.ServiceA](dm, services.ServiceASingletonKey)
serviceB, _ := dependencies.GetSingleton[*services.ServiceB](dm, services.ServiceBSingletonKey)

serviceA.Print()
serviceB.Print()
```

## Example Project Structure

```
go-dependency-injector/
├── main.go
├── pkg/
│   ├── dependencies/
│   │   └── factory.go
│   ├── infra/
│   │   ├── cache.go
│   │   └── database.go
│   └── services/
│       ├── service_a.go
│       └── service_b.go
```

## Running the Project

1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Execute the application:

```bash
go run main.go
```

## Conclusion

This project demonstrates a simple and effective way to implement dependency injection in Go. By using a `DependencyManager`, you can manage the lifecycle of your dependencies, ensuring they are initialized and accessible in a type-safe manner.

Feel free to extend this framework to suit your application's needs!