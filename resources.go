package tx

import "sync"

type TransactionContextResources interface {
	AddResource(key string, resource interface{})
	ContainsResource(key string) bool
	GetResource(key string) interface{}
}

type SimpleTransactionContextResources struct {
	resources map[string]interface{}
	mu        sync.RWMutex
}

func NewSimpleTransactionContextResources() *SimpleTransactionContextResources {
	return &SimpleTransactionContextResources{
		resources: make(map[string]interface{}),
		mu:        sync.RWMutex{},
	}
}

func (tr SimpleTransactionContextResources) AddResource(key string, resource interface{}) {
	tr.mu.Lock()
	tr.resources[key] = resource
	tr.mu.Unlock()
}

func (tr SimpleTransactionContextResources) ContainsResource(key string) bool {
	tr.mu.Lock()
	if _, ok := tr.resources[key]; ok {
		return true
	}
	tr.mu.Unlock()
	return false
}

func (tr SimpleTransactionContextResources) GetResource(key string) interface{} {
	var result interface{}
	tr.mu.Lock()
	if resource, ok := tr.resources[key]; ok {
		result = resource
	}
	tr.mu.Unlock()
	return result
}
