package sqllite

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"
)

func orderToRow(order *ds.Order, row *sds.OrderRow) {
	row.ID = order.ID
	row.Type = int32(order.OrderType)
	row.BuyerID = order.BuyerID
	row.SupplierID = order.SupplierID
	row.Price = order.Price

	if order.Slot == nil {
		order.Slot = &ds.Slot{}
	}

	row.Duration = order.Slot.Duration
	row.BuyerRating = order.Slot.BuyerRating
	row.SupplierRating = order.Slot.SupplierRating

	row.CPUCores = order.Slot.Resources.CPUCores
	row.RAMBytes = order.Slot.Resources.RAMBytes
	row.GPUCount = uint64(order.Slot.Resources.GPUCount)
	row.Storage = order.Slot.Resources.Storage

	row.NetType = uint64(order.Slot.Resources.NetworkType)
	row.NetInbound = order.Slot.Resources.NetTrafficIn
	row.NetOutbound = order.Slot.Resources.NetTrafficOut

	row.Properties = sds.Properties(order.Slot.Resources.Properties)
}

func orderFromRow(order *ds.Order, row *sds.OrderRow) {
	if order == nil {
		return
	}

	order.ID = row.ID
	order.OrderType = ds.OrderType(row.Type)
	order.Price = row.Price

	order.BuyerID = row.BuyerID
	order.SupplierID = row.SupplierID

	if order.Slot == nil {
		order.Slot = &ds.Slot{}
	}

	slot := &ds.Slot{
		Duration:       row.Duration,
		BuyerRating:    row.BuyerRating,
		SupplierRating: row.SupplierRating,

		Resources: ds.Resources{
			CPUCores: row.CPUCores,
			RAMBytes: row.RAMBytes,
			GPUCount: ds.GPUCount(row.GPUCount),
			Storage:  row.Storage,

			NetworkType:   ds.NetworkType(row.NetType),
			NetTrafficIn:  row.NetInbound,
			NetTrafficOut: row.NetOutbound,

			Properties: row.Properties,
		},
	}

	order.Slot = slot
}
