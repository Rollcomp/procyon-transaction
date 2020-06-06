package tx

import "log"

type InvokeCallback func()

func invokeWithinTransaction(txDef TransactionDefinition, txManager TransactionManager, invokeCallback InvokeCallback) {
	if invokeCallback == nil {
		panic("Invoke Callback function must not be null")
	}
	transactionStatus := txManager.GetTransaction(txDef)
	defer func() {
		if r := recover(); r != nil {
			err := txManager.Rollback(transactionStatus)
			if err != nil {
				log.Print("Rollback mechanism could not be worked")
			}
		}
	}()
	invokeCallback()
	defer func() {
		if r := recover(); r != nil {
			/* ... */
		}
	}()
	_ = txManager.Commit(transactionStatus)
}
