package tx

import "errors"

type TransactionManagerAdapter interface {
	DoGetTransaction() TransactionObject
	DoBeginTransaction(txObj TransactionObject, txDef TransactionDefinition)
	DoSuspendTransaction(txObj TransactionObject) TransactionSuspendedResources
	DoResumeTransaction(txObj TransactionObject)
	DoCommitTransaction(txStatus TransactionStatus)
	DoRollback(txStatus TransactionStatus)
	IsExistingTransaction(txObj TransactionObject) bool
	SupportsPropagation(propagation TransactionPropagation) bool
}

type TransactionManager interface {
	GetTransaction(txDef TransactionDefinition) TransactionStatus
	Commit(txStatus TransactionStatus) error
	Rollback(txStatus TransactionStatus) error
}

type AbstractTransactionManager struct {
	TransactionManagerAdapter
}

func NewAbstractTransactionManager(txManagerAdapter TransactionManagerAdapter) *AbstractTransactionManager {
	if txManagerAdapter == nil {
		panic("This is an abstract. That's why transaction manager adapter must not be null")
	}
	return &AbstractTransactionManager{
		txManagerAdapter,
	}
}

func (txManager *AbstractTransactionManager) GetTransaction(txDef TransactionDefinition) TransactionStatus {
	if txDef == nil {
		txDef = NewDefaultTransactionDefinition()
	}
	if txManager.SupportsPropagation(txDef.GetPropagation()) {
		panic("Propagation is not supported by current transaction manager.")
	}
	if txDef.GetPropagation() == PropagationNested {
		panic("unfortunately it is not supported yet by procyon transaction")
	}
	tx := txManager.DoGetTransaction()
	if !txManager.IsExistingTransaction(tx) {

	} else if txDef.GetPropagation() == PropagationMandatory {
		panic("There must be an existing transaction for Propagation Mandatory")
	} else if txDef.GetPropagation() == PropagationRequired || txDef.GetPropagation() == PropagationRequiredNew {
		/* start transaction */
	} else {

	}
	return nil
}

func (txManager *AbstractTransactionManager) Commit(txStatus TransactionStatus) error {
	if !txStatus.IsCompleted() {

	} else {
		return errors.New("transaction is already completed")
	}
	defer txStatus.SetCompleted()
	txManager.DoCommitTransaction(txStatus)
	return nil
}

func (txManager *AbstractTransactionManager) Rollback(txStatus TransactionStatus) error {
	if !txStatus.IsCompleted() {

	} else {
		return errors.New("transaction is already completed")
	}
	defer txStatus.SetCompleted()
	txManager.DoRollback(txStatus)
	return nil
}

func (txManager *AbstractTransactionManager) handleExistingTransaction(tx interface{}, txDef TransactionDefinition) {
	if txDef.GetPropagation() == PropagationNever {
		panic("Propagation never does not support an existing transaction which was created before")
	} else {
		if txDef.GetPropagation() == PropagationNotSupported {

		} else if txDef.GetPropagation() == PropagationRequiredNew {

		} else {
			/* participate */
		}
	}
}

func (txManager *AbstractTransactionManager) startTransaction(txObj TransactionObject, txDef TransactionDefinition) TransactionStatus {
	txManager.DoBeginTransaction(txObj, txDef)
	return nil
}

func (txManager *AbstractTransactionManager) suspendTransaction(txObj TransactionObject) TransactionSuspendedResources {
	return txManager.DoSuspendTransaction(txObj)
}
