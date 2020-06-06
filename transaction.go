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
}

type DefaultTransactionDefinitionOption func(txDef *DefaultTransactionDefinition)

func WithTxPropagation(propagation TransactionPropagation) DefaultTransactionDefinitionOption {
	return func(txDef *DefaultTransactionDefinition) {
		txDef.propagation = propagation
	}
}

func WithTxTimeout(timeout int) DefaultTransactionDefinitionOption {
	return func(txDef *DefaultTransactionDefinition) {
		txDef.timeout = timeout
	}
}

type DefaultTransactionDefinition struct {
	propagation TransactionPropagation
	timeout     int
}

func NewDefaultTransactionDefinition(options ...DefaultTransactionDefinitionOption) *DefaultTransactionDefinition {
	def := &DefaultTransactionDefinition{
		PropagationRequired,
		-1,
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
	return txDef.timeout
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
