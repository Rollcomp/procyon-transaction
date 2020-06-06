package tx

const TransactionMinTimeout = 1000

type TransactionObject interface {
	GetTransaction() interface{}
}

type TransactionSuspendedResources interface {
	GetSuspendedResources() interface{}
}

type TransactionDefinitionOption func(definition *TransactionDefinition)

type TransactionDefinition interface {
	GetTimeout() int
	GetPropagation() TransactionPropagation
	IsReadOnly() bool
}

type SimpleTransactionDefinitionOption func(txDef *SimpleTransactionDefinition)

func WithTxPropagation(propagation TransactionPropagation) SimpleTransactionDefinitionOption {
	return func(txDef *SimpleTransactionDefinition) {
		txDef.propagation = propagation
	}
}

func WithTxTimeout(timeOut int) SimpleTransactionDefinitionOption {
	return func(txDef *SimpleTransactionDefinition) {
		txDef.timeOut = timeOut
	}
}

func WithTxReadOnly(readOnly bool) SimpleTransactionDefinitionOption {
	return func(txDef *SimpleTransactionDefinition) {
		txDef.readOnly = readOnly
	}
}

type SimpleTransactionDefinition struct {
	propagation TransactionPropagation
	timeOut     int
	readOnly    bool
}

func NewSimpleTransactionDefinition(options ...SimpleTransactionDefinitionOption) *SimpleTransactionDefinition {
	def := &SimpleTransactionDefinition{
		PropagationRequired,
		-1,
		false,
	}
	for _, option := range options {
		option(def)
	}
	return def
}

func (txDef *SimpleTransactionDefinition) GetPropagation() TransactionPropagation {
	return txDef.propagation
}

func (txDef *SimpleTransactionDefinition) GetTimeout() int {
	return txDef.timeOut
}

func (txDef *SimpleTransactionDefinition) IsReadOnly() bool {
	return txDef.readOnly
}

type TransactionStatus interface {
	GetTransactionObject() TransactionObject
	GetTransactionDefinition() TransactionDefinition
	IsCompleted() bool
	SetCompleted()
	GetSuspendedResources() interface{}
}

type defaultTransactionStatus struct {
	txObj              TransactionObject
	txDef              TransactionDefinition
	isCompleted        bool
	suspendedResources interface{}
}

func newDefaultTransactionStatus(txObj TransactionObject, txDef TransactionDefinition, suspendedResources interface{}) *defaultTransactionStatus {
	return &defaultTransactionStatus{
		txObj,
		txDef,
		false,
		suspendedResources,
	}
}

func (txStatus *defaultTransactionStatus) SetCompleted() {
	txStatus.isCompleted = true
}

func (txStatus *defaultTransactionStatus) GetTransactionObject() TransactionObject {
	return txStatus.txObj
}

func (txStatus *defaultTransactionStatus) GetTransactionDefinition() TransactionDefinition {
	return txStatus.txDef
}

func (txStatus *defaultTransactionStatus) IsCompleted() bool {
	return txStatus.isCompleted
}

func (txStatus *defaultTransactionStatus) GetSuspendedResources() interface{} {
	return txStatus.suspendedResources
}
