package tx

type TransactionResources interface {
	AddResource(key string, resource interface{})
	ContainsResource(key string) bool
	GetResource(key string) interface{}
	GetResources() []interface{}
}
