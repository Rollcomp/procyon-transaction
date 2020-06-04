package tx

import "github.com/google/uuid"

type TransactionalContext interface {
	TransactionalBlock
	GetContextId() uuid.UUID
	GetTransactionManager() TransactionManager
	GetTransactionResources() TransactionResources
}
