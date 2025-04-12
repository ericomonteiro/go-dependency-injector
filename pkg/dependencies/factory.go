package dependencies

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

var (
	errDependencyNotFound = errors.New("dependency not found")
)

type Singleton interface {
	Initialize(manager *DependencyManager) error
}

type DependencyManager struct {
	singletons []Singleton
	types      map[string]Singleton
	ctx        context.Context
}

func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		singletons: make([]Singleton, 0),
		types:      make(map[string]Singleton),
		ctx:        context.Background(),
	}
}

func (dm *DependencyManager) AutoWire() error {
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

		// Iterate through the fields of the struct
		for i := 0; i < objType.NumField(); i++ {
			//Check if field is a pointer
			field := objType.Field(i)
			if field.Type.Kind() != reflect.Ptr {
				continue
			}

			instance := dm.GetSingletonByType(field.Type)
			if instance == nil {
				return errors.New(fmt.Sprintf("singleton %s dependency %s not found", objType.String(), field.Type.String()))
			}

			singletonValue.Field(i).Set(reflect.ValueOf(instance))
		}
	}

	return nil
}

func (dm *DependencyManager) InitializeAll() error {
	for _, singleton := range dm.singletons {
		err := singleton.Initialize(dm)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dm *DependencyManager) Register(key any, singleton Singleton) {
	fullType := reflect.TypeOf(singleton).String()
	dm.singletons = append(dm.singletons, singleton)
	dm.types[fullType] = singleton
	dm.ctx = context.WithValue(dm.ctx, key, singleton)
}

func (dm *DependencyManager) get(key any) (Singleton, error) {
	fromContext := dm.ctx.Value(key)
	if fromContext == nil {
		return nil, errDependencyNotFound
	}

	dep, _ := fromContext.(Singleton)

	return dep, nil
}

func GetSingletonByType[T any](dm *DependencyManager) T {
	var nilReturn T

	for _, singleton := range dm.singletons {
		instance, ok := singleton.(T)
		if ok {
			return instance
		}
	}

	return nilReturn
}

func (dm *DependencyManager) GetSingletonByType(depType reflect.Type) any {
	return dm.types[depType.String()]
}

func GetSingletonByKey[T any](dm *DependencyManager, key any) (T, error) {
	var nilReturn T

	singleton, err := dm.get(key)
	if err != nil {
		return nilReturn, err
	}

	instance, ok := singleton.(T)
	if !ok {
		return nilReturn, errors.New("type assertion failed")
	}

	return instance, nil
}
