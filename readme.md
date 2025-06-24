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
 
Note: You must use the `Key` method to define a unique key for each singleton. This key is used for dependency injection.
Using the `inject` tag.

You can use either explicit key-based injection or automatic type-based injection:

1. **Explicit Key-based Injection**
```go
type ServiceB struct {
    Database *infra.Database `inject:"DataBaseSingletonKey"`
    Cache    *infra.Cache    `inject:"CacheSingletonKey"`
}
```

2. **Automatic Type-based Injection**
```go
type ServiceB struct {
    Database *infra.Database `inject:"auto"`
    Cache    *infra.Cache    `inject:"auto"`
}
```

When using `inject="auto"`, the dependency manager will automatically resolve the dependency based on its type. This makes your code more maintainable and reduces the chance of errors from mismatched keys.

**Note**: When using interfaces with multiple implementations, you must use explicit key-based injection to avoid ambiguity. The dependency manager cannot automatically resolve which implementation to use when multiple options are available.

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

Result with this example:
```
Dependency graph for infra.Cache:
Dependency graph for services.ServiceB:
  - *infra.Database
  - *infra.Cache
Dependency graph for services.ServiceA:
  - *services.ServiceB
Dependency graph for infra.Database:

Service A singleton
    My address:  &{0x14000010070}
    Service B address:  &{0x1400000e0b8 0x14000010060}
    Service B database:  &{0.6024735673187461}
    Service B cache:  &{19}
Service B singleton
    My address:  &{0x1400000e0b8 0x14000010060}
    Database address:  &{0.6024735673187461}
    Cache address:  &{19}
```

### Example project structure
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
