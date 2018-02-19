package service

// Storage db gateway.
type Storage interface {
	InsertRow(query string, args ...interface{}) error
	UpdateRow(query string, value ...interface{}) error

	FetchRow(row interface{}, query string, value ...interface{}) error
	FetchRows(rows interface{}, query string, value ...interface{}) error
}

// MarketService manages orders.
type MarketService struct {
	s Storage
}

// NewMarketService creates a new instance of MarketService.
func NewMarketService(s Storage) *MarketService {
	return &MarketService{s: s}
}
