package repository

import "career_focus_hw/app/api"

type EthRepo interface {
	ById(address string) ([]api.Transaction, error)
	Save(address string, transaction api.Transaction) (bool, error)
}

func NewEthRepo(repo EthRepo) *EthRepo {
	return &repo
}
