package tx

import "sync"

type TransactionResources interface {
	AddResource(key interface{}, resource interface{})
	ContainsResource(key interface{}) bool
	GetResource(key interface{}) interface{}
	RemoveResource(key interface{}) interface{}
}

type SimpleTransactionResources struct {
	resources map[interface{}]interface{}
	mu        sync.RWMutex
}

func NewSimpleTransactionResources() *SimpleTransactionResources {
	return &SimpleTransactionResources{
		resources: make(map[interface{}]interface{}),
		mu:        sync.RWMutex{},
	}
}

func (tr SimpleTransactionResources) AddResource(key interface{}, resource interface{}) {
	tr.mu.Lock()
	tr.resources[key] = resource
	tr.mu.Unlock()
}

func (tr SimpleTransactionResources) ContainsResource(key interface{}) bool {
	tr.mu.Lock()
	if _, ok := tr.resources[key]; ok {
		return true
	}
	tr.mu.Unlock()
	return false
}

func (tr SimpleTransactionResources) GetResource(key interface{}) interface{} {
	var result interface{}
	tr.mu.Lock()
	if resource, ok := tr.resources[key]; ok {
		result = resource
	}
	tr.mu.Unlock()
	return result
}

func (tr SimpleTransactionResources) RemoveResource(key interface{}) interface{} {
	var result interface{}
	tr.mu.Lock()
	if resource, ok := tr.resources[key]; ok {
		result = resource
		delete(tr.resources, key)
	}
	tr.mu.Unlock()
	return result
}

type TransactionResourcesManager interface {
	GetResource(key interface{}) interface{}
	BindResource(key interface{}, resource interface{})
	UnBindResource(key interface{})
}

type SimpleTransactionResourcesManager struct {
	resources TransactionResources
}

func NewSimpleTransactionResourcesManager(resources TransactionResources) SimpleTransactionResourcesManager {
	if resources == nil {
		panic("Transactional context must not be null")
	}
	return SimpleTransactionResourcesManager{
		resources,
	}
}

func (resourceManager SimpleTransactionResourcesManager) GetResource(key interface{}) interface{} {
	resources := resourceManager.resources
	if resources == nil {
		return nil
	}
	return resources.GetResource(key)
}

func (resourceManager SimpleTransactionResourcesManager) BindResource(key interface{}, value interface{}) {
	resources := resourceManager.resources
	if resources == nil {
		panic("Transactional context resources must not be null")
	}
	if resources.ContainsResource(key) {
		panic("There is already added resource with same key in transactional context")
	}
	resources.AddResource(key, value)
}

func (resourceManager SimpleTransactionResourcesManager) UnBindResource(key interface{}) {
	resources := resourceManager.resources
	if resources == nil {
		panic("Transactional context resources must not be null")
	}
	if !resources.ContainsResource(key) {
		panic("There is no resource for given key in transactional context")
	}
	resources.RemoveResource(key)
}
