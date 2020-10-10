package tx

import (
	context "github.com/procyon-projects/procyon-context"
)

type TransactionalContext interface {
	TransactionalBlock
	GetTransactionManager() TransactionManager
	GetTransactionResourcesManager() TransactionResourcesManager
}

type SimpleTransactionalContext struct {
	logger                  context.Logger
	transactionManager      TransactionManager
	transactionResourcesMgr TransactionResourcesManager
}

func newSimpleTransactionalContext() *SimpleTransactionalContext {
	return &SimpleTransactionalContext{}
}

func NewSimpleTransactionalContext(logger context.Logger,
	transactionManager TransactionManager,
	transactionResourcesManager TransactionResourcesManager) *SimpleTransactionalContext {
	if logger == nil {
		panic("logger must not be nil")
	}
	if transactionManager == nil {
		panic("transaction manager must not be nil")
	}
	if transactionResourcesManager == nil {
		panic("transaction resource Manager must not be nil")
	}
	transactionalContext := newSimpleTransactionalContext()
	transactionalContext.logger = logger
	transactionalContext.transactionManager = transactionManager
	transactionalContext.transactionResourcesMgr = transactionResourcesManager
	return transactionalContext
}

func (tContext *SimpleTransactionalContext) Block(ctx context.Context, fun TransactionalFunc, options ...TransactionBlockOption) {
	if ctx == nil {
		panic("context must not be nil")
	}
	if fun == nil {
		panic("transaction function must not be nil")
	}
	txBlockObject := NewTransactionBlockObject(fun, options...)
	/* convert tx block object into tx block definition */
	txBlockDef := NewSimpleTransactionDefinition(
		WithTxContext(ctx),
		WithTxPropagation(txBlockObject.propagation),
		WithTxReadOnly(txBlockObject.readOnly),
		WithTxTimeout(txBlockObject.timeOut),
	)
	/* invoke within transaction */
	invokeWithinTransaction(ctx, tContext.logger, txBlockDef, tContext.GetTransactionManager(), func() {
		txFunc := txBlockObject.fun
		txFunc()
	})
}

func (tContext *SimpleTransactionalContext) GetTransactionManager() TransactionManager {
	return tContext.transactionManager
}

func (tContext *SimpleTransactionalContext) GetTransactionResourcesManager() TransactionResourcesManager {
	return tContext.transactionResourcesMgr
}
