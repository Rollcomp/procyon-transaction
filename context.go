package tx

type TransactionContext interface {
	GetTransactionManager() TransactionManager
}
