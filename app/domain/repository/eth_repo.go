package repository

import "career_focus_hw/app/api"

type EthRepo interface {
	Save(address string, transaction api.Transaction) (bool, error)
	ById(address string) ([]api.Transaction, error)
}

func NewEthRepo(repo EthRepo) EthRepo {
	return repo
}
