package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/event"
)

// CreateBidOrderStorage adds an order to the storage.
type CreateBidOrderStorage interface {
	Add(o *entity.Order) error
}

// CreateBidOrderHandler creates new bid orders.
type CreateBidOrderHandler struct {
	intf.Observable

	s CreateBidOrderStorage
}

// NewCreateBidOrderHandler creates a new instance of CreateBidOrderHandler.
func NewCreateBidOrderHandler(s CreateBidOrderStorage) *CreateBidOrderHandler {
	return &CreateBidOrderHandler{s: s}
}

// Handle handles the given command.
// Creates the given bid order.
func (h *CreateBidOrderHandler) Handle(cmd intf.Command) error {

	c, ok := cmd.(CreateBidOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	order, err := newBidOrder(c)
	if err != nil {
		return err
	}

	if err := h.s.Add(order); err != nil {
		return err
	}

	h.Publish(newBidOrderCreatedEvent(order))

	return nil
}

func newBidOrder(c CreateBidOrder) (*entity.Order, error) {

	res := entity.Resources(c.Slot.Resources)
	slot := entity.Slot{
		BuyerRating:    c.Slot.BuyerRating,
		SupplierRating: c.Slot.SupplierRating,
		Resources:      &res,
	}

	return entity.NewBidOrder(c.ID, c.BuyerID, c.Price, slot)
}

func newBidOrderCreatedEvent(order *entity.Order) event.BidOrderCreated {

	var eventOrder event.Order
	bindOrder(order, &eventOrder)

	return event.NewBidOrderCreated(eventOrder)
}

func bindOrder(entityOrder *entity.Order, eventOrder *event.Order) {
	if entityOrder == nil {
		return
	}

	var eventSlot event.Slot
	bindSlot(entityOrder.Slot, &eventSlot)

	eventOrder.ID = entityOrder.ID
	eventOrder.BuyerID = entityOrder.BuyerID
	eventOrder.Price = entityOrder.Price
	eventOrder.Slot = &eventSlot
}

func bindSlot(entitySlot *entity.Slot, eventSlot *event.Slot) {
	if entitySlot == nil {
		return
	}

	eventSlot.BuyerRating = entitySlot.BuyerRating
	eventSlot.SupplierRating = entitySlot.SupplierRating

	if entitySlot.Resources == nil {
		return
	}

	res := event.Resources(*entitySlot.Resources)
	eventSlot.Resources = &res
}
