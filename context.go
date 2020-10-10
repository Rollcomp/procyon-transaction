package tx

import (
	"errors"
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
	transactionResourcesManager TransactionResourcesManager) (*SimpleTransactionalContext, error) {
	if logger == nil {
		return nil, errors.New("logger must not be nil")
	}
	if transactionManager == nil {
		return nil, errors.New("transaction manager must not be nil")
	}
	if transactionResourcesManager == nil {
		return nil, errors.New("transaction resource Manager must not be nil")
	}
	transactionalContext := newSimpleTransactionalContext()
	transactionalContext.logger = logger
	transactionalContext.transactionManager = transactionManager
	transactionalContext.transactionResourcesMgr = transactionResourcesManager
	return transactionalContext, nil
}

func (tContext *SimpleTransactionalContext) Block(ctx context.Context, fun TransactionalFunc, options ...TransactionBlockOption) error {
	if ctx == nil {
		return errors.New("context must not be nil")
	}
	if fun == nil {
		return errors.New("transaction function must not be nil")
	}
	txBlockObject := NewTransactionBlockObject(fun, options...)
	/* convert tx block object into tx block definition */
	txBlockDef := NewSimpleTransactionDefinition(
		WithTxPropagation(txBlockObject.propagation),
		WithTxReadOnly(txBlockObject.readOnly),
		WithTxTimeout(txBlockObject.timeOut),
	)
	/* invoke within transaction */
	return invokeWithinTransaction(tContext.logger, txBlockDef, tContext.GetTransactionManager(), func() {
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
