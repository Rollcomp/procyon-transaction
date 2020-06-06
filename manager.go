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
	Commit(txStatus TransactionStatus)
	Rollback(txStatus TransactionStatus)
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
	// if given is nil, create a default one
	if txDef == nil {
		txDef = NewSimpleTransactionDefinition()
	}

	// custom implementations might not support all kind of propagation
	if !txManager.SupportsPropagation(txDef.GetPropagation()) {
		panic("Propagation is not supported by current transaction manager.")
	}

	// get the current transaction object
	txObj := txManager.DoGetTransaction()

	//  if there is an existing transaction, handle it
	//  if necessary, suspend or create new one depend on your cases
	if !txManager.IsExistingTransaction(txObj) {
		return txManager.handleExistingTransaction(txObj, txDef)
	}

	// don't check it for existing transaction
	if txDef.GetTimeout() < TransactionMinTimeout {
		panic("Invalid timeout for transaction")
	}
	if txDef.GetPropagation() == PropagationMandatory {
		panic("There must be an existing transaction for Propagation Mandatory")
	} else if txDef.GetPropagation() == PropagationRequired || txDef.GetPropagation() == PropagationRequiredNew {
		txSuspendedResources := txManager.suspendTransaction(nil)
		status := newDefaultTransactionStatus(txObj, txDef, txSuspendedResources)
		txManager.startTransaction(txObj, txDef)
		return status
	}
	// create a new empty transaction, it is not exactly a transaction
	return newDefaultTransactionStatus(nil, txDef, nil)
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

func (txManager *AbstractTransactionManager) handleExistingTransaction(txObj TransactionObject, txDef TransactionDefinition) TransactionStatus {
	// if there is an existing transaction, throw an error
	if txDef.GetPropagation() == PropagationNever {
		panic("Propagation never does not support an existing transaction which was created before")
	}
	if txDef.GetPropagation() == PropagationNotSupported {
		// if there is an existing transaction, first suspend it
		// don't create new one
		txSuspendedResources := txManager.suspendTransaction(txObj)
		return newDefaultTransactionStatus(txObj, txDef, txSuspendedResources)
	} else if txDef.GetPropagation() == PropagationRequiredNew {
		// suspend current transaction, then new start transaction
		txManager.startTransaction(txObj, txDef)
	}
	// PropagationMandatory, PropagationSupports, PropagationRequired
	// They will use the existing transaction
	return newDefaultTransactionStatus(txObj, txDef, nil)
}

func (txManager *AbstractTransactionManager) startTransaction(txObj TransactionObject, txDef TransactionDefinition) TransactionStatus {
	txSuspendedResources := txManager.suspendTransaction(txObj)
	status := newDefaultTransactionStatus(txObj, txDef, txSuspendedResources)
	txManager.DoBeginTransaction(txObj, txDef)
	return status
}

func (txManager *AbstractTransactionManager) suspendTransaction(txObj TransactionObject) TransactionSuspendedResources {
	return txManager.DoSuspendTransaction(txObj)
}
