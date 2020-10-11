package tx

import (
	"errors"
	"fmt"
	context "github.com/procyon-projects/procyon-context"
	"runtime/debug"
)

type InvokeCallback func() (interface{}, error)

func invokeWithinTransaction(context context.Context,
	logger context.Logger,
	txDef TransactionDefinition,
	txManager TransactionManager,
	invokeCallback InvokeCallback) (interface{}, error) {
	if invokeCallback == nil {
		panic("invoke Callback function must not be null")
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Error(context, fmt.Sprintf("%s\n%s", r, string(debug.Stack())))
		}
	}()
	/* create a transaction if necessary */
	transactionStatus, err := createTransactionIfNecessary(txDef, txManager)
	if err != nil {
		panic(err)
	}
	defer func() {
		if r := recover(); r != nil {
			/* rollback transaction */
			err = errors.New("transaction couldn't be completed successfully")
			logger.Error(context, err.Error())
			if txDef != nil && transactionStatus != nil {
				err = txManager.Rollback(transactionStatus)
				if err != nil {
					logger.Error(context, err.Error())
				}
			}
			panic(r)
		}
	}()
	/* invoke function */
	var result interface{}
	result, err = invokeCallback()
	if err != nil {
		if txDef != nil && transactionStatus != nil {
			err = txManager.Rollback(transactionStatus)
			if err != nil {
				logger.Error(context, err.Error())
			}
		}
		return result, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("transaction couldn't be committed")
			logger.Error(context, err.Error())
			panic(r)
		}
	}()
	/* complete transaction */
	if txDef != nil && transactionStatus != nil {
		err = txManager.Commit(transactionStatus)
	}
	return result, nil
}

func createTransactionIfNecessary(txDef TransactionDefinition, txManager TransactionManager) (TransactionStatus, error) {
	return txManager.GetTransaction(txDef)
}
