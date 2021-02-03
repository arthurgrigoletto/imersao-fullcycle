package repository

import (
	"fmt"
	"github.com/arthurgrigoletto/imersao-fullcycle/codepix-go/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (transactionRepository *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := transactionRepository.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (transactionRepository *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := transactionRepository.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (transactionRepository *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	transactionRepository.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}
	return &transaction, nil
}