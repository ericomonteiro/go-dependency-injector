# Go Dependency Injector

This project is a proposal to solve dependency injection in Go. It provides a simple and extensible approach to manage dependencies using a `DependencyManager` that supports singleton registration, initialization, and retrieval.

## Overview

Dependency injection is a design pattern that allows you to decouple the creation of dependencies from their usage. In this project, we implement a lightweight dependency injection framework for Go, enabling you to register and initialize singletons and retrieve them when needed.

### Dependency Registration Order

The order in which dependencies are registered does not matter. The `DependencyManager` creates "empty" instances of each singleton before resolving dependencies. After registration, the `Autowire` method is called to wire dependencies and complete the initialization process.

### Key Features

- **Singleton Management**: Register and manage singleton instances.
- **Auto-Wire**: Automatically wire dependencies using struct tags.
- **Type-Safe Retrieval**: Retrieve dependencies safely using generics.

## How to Use

### 1. Define Your Dependencies

Each dependency should implement the `Singleton` interface. For example:

```go
type Cache struct {
    dummy string
}

func (c *Cache) Key() string {
    return "CacheSingletonKey"
}

func NewCache() *Cache {
    return &Cache{}
}
```

### 2. Register Your Dependencies
```go
dm := dependencies.NewDependencyManager()

dm.Register(infra.NewCache())
dm.Register(infra.NewDatabase())
dm.Register(services.NewServiceA())
dm.Register(services.NewServiceB())
```

### 3. Auto-Wire Dependencies

Call `Autowire` to automatically wire all registered dependencies:
```go
if err := dm.Autowire(); err != nil {
    panic(err)
}
```

### 5. Generate Dependency Graph (Optional)
You can generate a dependency graph to visualize the relationships between your singletons:

```go
dm.GenerateDependencyGraph()
```

### Example project strucre
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

### 6. Run the Application

Finally, run your application:

```go
go run main.go
```
