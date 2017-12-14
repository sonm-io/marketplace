package datastruct

// Slot buying and selling object which includes resources and some additional parameters.
type Slot struct {
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	BuyerRating int64
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	SupplierRating int64
	// Hardware resources requirements.
	Resources Resources
}
