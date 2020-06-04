package tx

type TransactionalFunc interface{}
type TransactionBlockOption func(txBlockObj *TransactionBlockObject)

type TransactionBlockObject struct {
	fun         TransactionalFunc
	propagation TransactionPropagation
}

func NewTransactionBlockObject(fun TransactionalFunc, options ...TransactionBlockOption) *TransactionBlockObject {
	obj := &TransactionBlockObject{
		fun,
		PropagationRequired,
	}
	for _, opt := range options {
		opt(obj)
	}
	return obj
}

func (txBlockObj *TransactionBlockObject) GetTransactionFunc() TransactionalFunc {
	return txBlockObj.fun
}

func (txBlockObj *TransactionBlockObject) GetPropagation() TransactionPropagation {
	return txBlockObj.propagation
}

type TransactionalBlock interface {
	Block(fun TransactionalFunc, options ...TransactionBlockOption) error
}

func WithPropagation(propagation TransactionPropagation) TransactionBlockOption {
	return func(txBlockObj *TransactionBlockObject) {
		txBlockObj.propagation = propagation
	}
}
