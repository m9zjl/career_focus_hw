package api

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}
