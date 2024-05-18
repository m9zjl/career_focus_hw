package db

import (
	"career_focus_hw/app/api"
	"sync"
)

type LocalMem struct {
	lock sync.RWMutex
	data map[string][]api.Transaction
}

func NewLocalMem() *LocalMem {
	return &LocalMem{
		lock: sync.RWMutex{},
		data: make(map[string][]api.Transaction),
	}
}

func (r *LocalMem) ById(address string) ([]api.Transaction, error) {
	r.lock.RLock()
	defer r.lock.Unlock()
	return r.data[address], nil
}

func (w *LocalMem) Save(address string, transaction api.Transaction) (bool, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.data[address] = append(w.data[address], transaction)
	return true, nil
}
