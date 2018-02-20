package mapper

import (
	"github.com/sonm-io/marketplace/ds"
	sds "github.com/sonm-io/marketplace/mapper/datastruct"
	pb "github.com/sonm-io/marketplace/proto"
)

func OrderToRow(order *ds.Order, row *sds.OrderRow) {
	if order == nil {
		return
	}

	row.ID = order.Id
	row.Type = int32(order.OrderType)
	row.BuyerID = order.ByuerID
	row.SupplierID = order.SupplierID
	if order.PricePerSecond != nil {
		row.Price = order.PricePerSecond.Unwrap().String()
	}

	if order.Slot == nil {
		return
	}

	row.Duration = order.Slot.Duration
	row.BuyerRating = order.Slot.BuyerRating
	row.SupplierRating = order.Slot.SupplierRating

	row.CPUCores = order.Slot.Resources.CpuCores
	row.RAMBytes = order.Slot.Resources.RamBytes
	row.GPUCount = uint64(order.Slot.Resources.GpuCount)
	row.Storage = order.Slot.Resources.Storage

	row.NetType = uint64(order.Slot.Resources.NetworkType)
	row.NetInbound = order.Slot.Resources.NetTrafficIn
	row.NetOutbound = order.Slot.Resources.NetTrafficOut

	row.Properties = sds.Properties(order.Slot.Resources.Properties)
}

func OrderFromRow(order *ds.Order, row *sds.OrderRow) {
	if row == nil {
		return
	}

	if order == nil {
		return
	}

	order.Id = row.ID
	order.OrderType = pb.OrderType(row.Type)

	pricePerSecond, err := pb.NewBigIntFromString(row.Price)
	if err == nil {
		order.PricePerSecond = pricePerSecond
	}

	order.ByuerID = row.BuyerID
	order.SupplierID = row.SupplierID

	if order.Slot == nil {
		order.Slot = &pb.Slot{}
	}

	slot := order.Slot
	slot.Duration = row.Duration
	slot.BuyerRating = row.BuyerRating
	slot.SupplierRating = row.SupplierRating

	if slot.Resources == nil {
		slot.Resources = &pb.Resources{}
	}

	res := slot.Resources
	res.CpuCores = row.CPUCores
	res.RamBytes = row.RAMBytes
	res.GpuCount = pb.GPUCount(row.GPUCount)
	res.Storage = row.Storage

	res.NetworkType = pb.NetworkType(row.NetType)
	res.NetTrafficIn = row.NetInbound
	res.NetTrafficOut = row.NetOutbound
	res.Properties = row.Properties
}
