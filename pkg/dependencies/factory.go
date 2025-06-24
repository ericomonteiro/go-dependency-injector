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
	singletonsByKey  map[string]Singleton
	singletonsByType map[string]Singleton
}

func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		singletonsByKey:  make(map[string]Singleton),
		singletonsByType: make(map[string]Singleton),
	}
}

func (dm *DependencyManager) Autowire() error {
	for _, singleton := range dm.singletonsByKey {
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
			// Check if field is a pointer
			field := objType.Field(i)

			injectTag := field.Tag.Get("inject")
			if injectTag == "" {
				continue
			}

			var instance Singleton
			var isOk bool
			if injectTag == "auto" {
				// Check if field type is an interface
				if field.Type.Kind() == reflect.Interface {
					// find singleton implements the interface
					for _, s := range dm.singletonsByType {
						if reflect.TypeOf(s).Implements(field.Type) {
							instance = s
							isOk = true
							break
						}
					}
				} else {
					// Get Singleton by type
					fieldType := field.Type.String()
					instance, isOk = dm.singletonsByType[fieldType]
					if !isOk {
						return errors.New(fmt.Sprintf("singleton %s dependency %s not found", objType.String(), field.Type.String()))
					}
				}
			} else {
				// Get Singleton by key
				instance, isOk = dm.singletonsByKey[injectTag]
				if !isOk {
					return errors.New(fmt.Sprintf("singleton %s dependency %s not found", objType.String(), field.Type.String()))
				}
			}

			singletonValue.Field(i).Set(reflect.ValueOf(instance))
		}
	}

	return nil
}

func (dm *DependencyManager) Register(singleton Singleton) {
	objType := reflect.TypeOf(singleton)
	// Check if the object is a pointer and get the underlying type
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	dm.singletonsByKey[singleton.Key()] = singleton
	dm.singletonsByType[objType.String()] = singleton
}

func (dm *DependencyManager) GenerateDependencyGraph() {
	for _, singleton := range dm.singletonsByKey {
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
			if field.Type.Kind() != reflect.Ptr && field.Type.Kind() != reflect.Interface {
				continue
			}

			fmt.Printf("  - %s\n", field.Type.String())
		}
	}
}

func GetSingletonByKey[T any](dm *DependencyManager, key string) (T, error) {
	var nilReturn T

	singleton, isOk := dm.singletonsByKey[key]
	if !isOk {
		return nilReturn, errDependencyNotFound
	}

	instance, ok := singleton.(T)
	if !ok {
		return nilReturn, errDependencyTypeNotMatch
	}

	return instance, nil
}

func GetSingletonByType[T any](dm *DependencyManager, typ string) (T, error) {
	var nilReturn T
	objString := getStringType(new(T))

	singleton, isOk := dm.singletonsByType[objString]
	if !isOk {
		return nilReturn, errDependencyNotFound
	}

	instance, ok := singleton.(T)
	if !ok {
		return nilReturn, errDependencyTypeNotMatch
	}

	return instance, nil
}

func getStringType(obj any) string {
	objType := reflect.TypeOf(obj)
	// Check if the object is a pointer and get the underlying type
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.String()
}
