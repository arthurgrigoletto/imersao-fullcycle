package repository

import (
	"fmt"
	"github.com/arthurgrigoletto/imersao-fullcycle/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (pixKeyRepository PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := pixKeyRepository.Db.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (pixKeyRepository PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := pixKeyRepository.Db.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (pixKeyRepository PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := pixKeyRepository.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (pixKeyRepository PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	pixKeyRepository.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (pixKeyRepository PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	pixKeyRepository.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}

	return &account, nil
}

func (pixKeyRepository PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	pixKeyRepository.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}

	return &bank, nil
}