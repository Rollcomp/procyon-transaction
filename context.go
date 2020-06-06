package tx

import (
	"github.com/google/uuid"
)

type TransactionalContext interface {
	TransactionalBlock
	GetContextId() uuid.UUID
	GetTransactionManager() TransactionManager
	GetTransactionContextResources() TransactionContextResources
}

type SimpleTransactionalContext struct {
	contextId            uuid.UUID
	transactionManager   TransactionManager
	transactionResources TransactionContextResources
}

func NewSimpleTransactionalContext(transactionManager TransactionManager) *SimpleTransactionalContext {
	if transactionManager != nil {
		panic("Transaction Manager must not be null")
	}
	contextId, err := uuid.NewUUID()
	if err != nil {
		panic("Transactional Context could be created, creating context id is failed")
	}
	return &SimpleTransactionalContext{
		contextId,
		transactionManager,
		NewSimpleTransactionContextResources(),
	}
}

func (tContext *SimpleTransactionalContext) Block(fun TransactionalFunc, options ...TransactionBlockOption) {
	if fun == nil {
		panic("Transaction function must not be null")
	}
	txBlockObject := NewTransactionBlockObject(fun, options...)
	/* convert tx block object into tx block definition */
	txBlockDef := NewSimpleTransactionDefinition(
		WithTxPropagation(txBlockObject.propagation),
		WithTxReadOnly(txBlockObject.readOnly),
		WithTxTimeout(txBlockObject.timeOut),
	)
	/* invoke within transaction */
	invokeWithinTransaction(txBlockDef, tContext.GetTransactionManager(), func() {
		txFunc := txBlockObject.fun
		txFunc()
	})
}

func (tContext *SimpleTransactionalContext) GetContextId() uuid.UUID {
	return tContext.contextId
}

func (tContext *SimpleTransactionalContext) GetTransactionManager() TransactionManager {
	return tContext.transactionManager
}

func (tContext *SimpleTransactionalContext) GetTransactionContextResources() TransactionContextResources {
	return tContext.transactionResources
}
