package dependencies

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errDependencyNotFound     = errors.New("dependency not found")
	errDependencyTypeNotMatch = errors.New("dependency type not match")
)

type Singleton interface {
	Key() string
}

type DependencyManager struct {
	singletons map[string]Singleton
}

func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		singletons: make(map[string]Singleton),
	}
}

func (dm *DependencyManager) Autowire() error {
	for _, singleton := range dm.singletons {
		singletonValue := reflect.ValueOf(singleton)
		if singletonValue.Kind() == reflect.Ptr {
			singletonValue = singletonValue.Elem()
		}

		objType := reflect.TypeOf(singleton)
		// Check if the object is a pointer and get the underlying type
		if objType.Kind() == reflect.Ptr {
			objType = objType.Elem()
		}

		// Iterate through the fields of the struct
		for i := 0; i < objType.NumField(); i++ {
			//Check if field is a pointer
			field := objType.Field(i)

			injectTag := field.Tag.Get("inject")
			if injectTag == "" {
				continue
			}

			instance, isOk := dm.singletons[injectTag]
			if !isOk {
				return errors.New(fmt.Sprintf("singleton %s dependency %s not found", objType.String(), field.Type.String()))
			}

			singletonValue.Field(i).Set(reflect.ValueOf(instance))
		}
	}

	return nil
}

func (dm *DependencyManager) Register(singleton Singleton) {
	dm.singletons[singleton.Key()] = singleton
}

func (dm *DependencyManager) GenerateDependencyGraph() {
	for _, singleton := range dm.singletons {
		singletonValue := reflect.ValueOf(singleton)
		if singletonValue.Kind() == reflect.Ptr {
			singletonValue = singletonValue.Elem()
		}

		// Get the type of the object
		objType := reflect.TypeOf(singleton)

		// Check if the object is a pointer and get the underlying type
		if objType.Kind() == reflect.Ptr {
			objType = objType.Elem()
		}

		fmt.Printf("Dependency graph for %s:\n", objType.String())
		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			if field.Type.Kind() != reflect.Ptr {
				continue
			}

			fmt.Printf("  - %s\n", field.Type.String())
		}
	}
}

func GetSingletonByKey[T any](dm *DependencyManager, key string) (T, error) {
	var nilReturn T

	singleton, isOk := dm.singletons[key]
	if !isOk {
		return nilReturn, errDependencyNotFound
	}

	instance, ok := singleton.(T)
	if !ok {
		return nilReturn, errDependencyTypeNotMatch
	}

	return instance, nil
}
