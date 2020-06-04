package tx

type TransactionDefinitionOption func(definition *TransactionDefinition)

type TransactionDefinition interface {
	GetPropagation() TransactionPropagation
	GetTransactionBlockObject() *TransactionBlockObject
}

type DefaultTransactionDefinition struct {
	txBlockObj *TransactionBlockObject
}

func NewTransactionDefinition(txBlockObj *TransactionBlockObject) *DefaultTransactionDefinition {
	def := &DefaultTransactionDefinition{
		txBlockObj,
	}
	return def
}

func (txDef *DefaultTransactionDefinition) GetPropagation() TransactionPropagation {
	return txDef.txBlockObj.GetPropagation()
}

func (txDef *DefaultTransactionDefinition) GetTransactionBlockObject() *TransactionBlockObject {
	return txDef.txBlockObj
}

type TransactionStatusInfo interface {
	GetTransaction() interface{}
	IsCompleted() bool
}

type DefaultTransactionStatusInfo struct {
	tx          interface{}
	isCompleted bool
}

func NewTransactionStatusInfo(transaction interface{}) *DefaultTransactionStatusInfo {
	return &DefaultTransactionStatusInfo{
		transaction,
		false,
	}
}

func (txStatus *DefaultTransactionStatusInfo) SetCompleted(isCompleted bool) {
	txStatus.isCompleted = isCompleted
}

func (txStatus *DefaultTransactionStatusInfo) GetTransaction() interface{} {
	return txStatus.tx
}

func (txStatus *DefaultTransactionStatusInfo) IsCompleted() bool {
	return txStatus.isCompleted
}
