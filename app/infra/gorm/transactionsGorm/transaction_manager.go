package transactionsGorm

import (
	"gorm.io/gorm"
	"sync/atomic"
)

//TransactionManager struct.
type TransactionManager struct {
	db           *gorm.DB // store db instance
	tx           *gorm.DB // store active transaction
	transCounter int64    // arc counter
}

//GetTx .
func (t *TransactionManager) GetTx() *gorm.DB {
	return t.tx
}

//Transaction .
func (t *TransactionManager) Transaction(callback func() error) error {
	t.begin()
	//fmt.Println("transactionManager.txBegin")

	defer func() {
		if err := recover(); err != nil {
			// @TODO Loggar erro
			//fmt.Println("transactionManager.txRollback E:", fmt.Sprintf("%s", err))
			t.transCounter = 0
			t.rollback()
		}
	}()

	// get the error and

	if err := callback(); err != nil {
		// @TODO Loggar erro
		//fmt.Println("transactionManager.txRollback E:", fmt.Sprintf("%s", err))
		t.transCounter = 0
		t.rollback()
		return err
	}

	//fmt.Println("transactionManager.txCommit")
	t.commit()

	return nil
}

func (t *TransactionManager) begin() {
	// first time no transaction start yet
	if t.transCounter == 0 {
		// create a internal ref tx
		t.tx = t.db.Begin()
		//fmt.Println("transactionManager.begin.real")
	} else if t.transCounter >= 1 && t.supportSavePoint() {
		// after the first time we create a savepoint if the db were supported
		t.createSavePoint()
	}

	// increase arc var
	atomic.AddInt64(&t.transCounter, 1)
	//fmt.Println("transactionManager.begin transCounter->", t.transCounter)

	// @todo maybe fire [beganTransaction] event
}

func (t *TransactionManager) commit() {
	if t.transCounter == 1 {
		t.tx.Commit()
		//fmt.Println("transactionManager.commit.real")
	}

	// trigger this to be maintains the ref counting
	if t.transCounter > 0 {
		atomic.AddInt64(&t.transCounter, -1)
	}
	//fmt.Println("transactionManager.commit transCounter->", t.transCounter)

	// @todo maybe fire [committed] event
}

func (t *TransactionManager) rollback() {
	if t.transCounter == 0 {
		// create a internal ref tx
		t.tx.Rollback()
		//fmt.Println("transactionManager.rollback.real")
	} else if t.transCounter >= 1 && t.supportSavePoint() {
		t.removeSavePoint()
	}
	//fmt.Println("transactionManager.rollback transCounter->", t.transCounter)
}

//Level get the number of active transactions.
func (t *TransactionManager) Level() int64 {
	return t.transCounter
}

//createSavePoint create a save point within the database.
func (t *TransactionManager) createSavePoint() {
	// @todo db.execRaw('added up savepoint')

}

// removeSavePoint rollback.
func (t *TransactionManager) removeSavePoint() {
	// @todo db.exeRaw('rollback savepoint')
}

// supportSavePoint db check.
func (t *TransactionManager) supportSavePoint() bool {
	// @todo db check
	return false
}

//NewTransactionManager cria novo transaction manager.
func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}
