package event

import (
	"fmt"
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

type ReadStorage interface {
	Add(order *report.GetOrderReport) error
}

type BidOrderCreatedHandler struct {
	s ReadStorage
}

func NewBidOrderCreatedHandler(s ReadStorage) *BidOrderCreatedHandler {
	return &BidOrderCreatedHandler{s: s}
}

func (h *BidOrderCreatedHandler) Handle(event intf.Event) error {
	e, ok := event.(BidOrderCreated)
	if !ok {
		return fmt.Errorf("invalid event %v given", event)
	}

	var reportOrder report.GetOrderReport
	bindOrder(&e.Order, &reportOrder)

	h.s.Add(&reportOrder)
	return nil
}

func bindOrder(eventOrder *Order, reportOrder *report.GetOrderReport) {
	if eventOrder == nil {
		return
	}

	var reportSlot report.Slot
	bindSlot(eventOrder.Slot, &reportSlot)

	reportOrder.ID = eventOrder.ID
	reportOrder.BuyerID = eventOrder.BuyerID
	reportOrder.Price = eventOrder.Price
	reportOrder.Slot = &reportSlot
}

func bindSlot(eventSlot *Slot, reportSlot *report.Slot) {
	if eventSlot == nil {
		return
	}

	reportSlot.BuyerRating = eventSlot.BuyerRating
	reportSlot.SupplierRating = eventSlot.SupplierRating

	if eventSlot.Resources == nil {
		return
	}

	res := report.Resources(*eventSlot.Resources)
	reportSlot.Resources = &res
}
