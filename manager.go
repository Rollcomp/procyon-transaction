package tx

type TransactionManager interface {
	GetTransaction() interface{}
	BeginTransaction(tx interface{}, txDefinition TransactionDefinition)
	SuspendTransaction(tx interface{})
	ResumeTransaction(tx interface{})
	CommitTransaction(txInfo TransactionStatusInfo)
	Rollback(txInfo TransactionStatusInfo)
	IsExistingTransaction(tx interface{}) bool
}
