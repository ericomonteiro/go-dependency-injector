package dependencies

import (
	"context"
	"errors"
)

var (
	errDependencyNotFound = errors.New("dependency not found")
)

type Singleton interface {
	Initialize(manager *DependencyManager) error
}

type DependencyManager struct {
	singletons []Singleton
	ctx        context.Context
}

func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		singletons: make([]Singleton, 0),
		ctx:        context.Background(),
	}
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
	dm.singletons = append(dm.singletons, singleton)
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
