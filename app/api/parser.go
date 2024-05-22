package api

type Parser interface {
	GetCurrentBlock() int64
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}
