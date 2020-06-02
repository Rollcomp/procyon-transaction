package tx

type TransactionInfo interface {
	GetTransactionStatus() string
}

type Transaction interface {
	Begin()
	Rollback()
	Commit()
}

type SimpleTransaction struct {
}

func NewSimpleTransaction() SimpleTransaction {
	return SimpleTransaction{}
}

func (tx SimpleTransaction) Begin() {

}

func (tx SimpleTransaction) Rollback() {

}

func (tx SimpleTransaction) Commit() {

}
