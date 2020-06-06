package tx

import "log"

type InvokeCallback func()

func invokeWithinTransaction(txDef TransactionDefinition, txManager TransactionManager, invokeCallback InvokeCallback) {
	if invokeCallback == nil {
		panic("Invoke Callback function must not be null")
	}
	/* create a transaction if necessary */
	transactionStatus := createTransactionIfNecessary(txDef, txManager)
	defer func() {
		if r := recover(); r != nil {
			/* rollback transaction */
			log.Printf("Transaction couldn't be completed successfully")
			if txDef != nil && transactionStatus != nil {
				txManager.Rollback(transactionStatus)
			}
		}
	}()
	/* invoke function */
	invokeCallback()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Transaction couldn't be committed")
		}
	}()
	/* complete transaction */
	if txDef != nil && transactionStatus != nil {
		log.Printf("Transaction has just been completed")
		txManager.Commit(transactionStatus)
	}
}

func createTransactionIfNecessary(txDef TransactionDefinition, txManager TransactionManager) TransactionStatus {
	return txManager.GetTransaction(txDef)
}
