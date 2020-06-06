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

type DefaultTransactionDefinitionOption func(txDef *DefaultTransactionDefinition)

func WithTxPropagation(propagation TransactionPropagation) DefaultTransactionDefinitionOption {
	return func(txDef *DefaultTransactionDefinition) {
		txDef.propagation = propagation
	}
}

func WithTxTimeout(timeOut int) DefaultTransactionDefinitionOption {
	return func(txDef *DefaultTransactionDefinition) {
		txDef.timeOut = timeOut
	}
}

func WithTxReadOnly(readOnly bool) DefaultTransactionDefinitionOption {
	return func(txDef *DefaultTransactionDefinition) {
		txDef.readOnly = readOnly
	}
}

type DefaultTransactionDefinition struct {
	propagation TransactionPropagation
	timeOut     int
	readOnly    bool
}

func NewDefaultTransactionDefinition(options ...DefaultTransactionDefinitionOption) *DefaultTransactionDefinition {
	def := &DefaultTransactionDefinition{
		PropagationRequired,
		-1,
		false,
	}
	for _, option := range options {
		option(def)
	}
	return def
}

func (txDef *DefaultTransactionDefinition) GetPropagation() TransactionPropagation {
	return txDef.propagation
}

func (txDef *DefaultTransactionDefinition) GetTimeout() int {
	return txDef.timeOut
}

func (txDef *DefaultTransactionDefinition) IsReadOnly() bool {
	return txDef.readOnly
}

type TransactionStatus interface {
	GetTransaction() interface{}
	IsCompleted() bool
	SetCompleted()
	GetSuspendedResources() interface{}
}

type defaultTransactionStatus struct {
	tx                 interface{}
	isCompleted        bool
	suspendedResources interface{}
}

func newDefaultTransactionStatus(transaction interface{}, suspendedResources interface{}) *defaultTransactionStatus {
	return &defaultTransactionStatus{
		transaction,
		false,
		suspendedResources,
	}
}

func (txStatus *defaultTransactionStatus) SetCompleted() {
	txStatus.isCompleted = true
}

func (txStatus *defaultTransactionStatus) GetTransaction() interface{} {
	return txStatus.tx
}

func (txStatus *defaultTransactionStatus) IsCompleted() bool {
	return txStatus.isCompleted
}

func (txStatus *defaultTransactionStatus) GetSuspendedResources() interface{} {
	return txStatus.suspendedResources
}
