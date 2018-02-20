package filter

import (
	"fmt"
	"github.com/gocraft/dbr"

	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/proto"
)

// Operator is used to indicate how to filter different values.
type Operator int

func (op Operator) String() string {
	var res string
	switch op {
	case LessThan:
		res = "<"
	case LessEq:
		res = "<="
	case Equal:
		res = "="
	case NotEqual:
		res = "!="
	case GreaterEq:
		res = ">="
	case GreaterThan:
		res = ">"
	default:
		res = "unknown"
	}
	return res
}

func (op Operator) Condition(column string, value interface{}) dbr.Condition {
	switch op {
	case LessThan:
		return dbr.Lt(column, value)
	case LessEq:
		return dbr.Lte(column, value)
	case Equal:
		return dbr.Eq(column, value)
	case NotEqual:
		return dbr.Neq(column, value)
	case GreaterEq:
		return dbr.Gte(column, value)
	case GreaterThan:
		return dbr.Gt(column, value)
	default:
		panic("unsupported operator given: " + op.String())
	}
}

const (
	// LessThan shows that the field being filtered must be less than the provided value
	LessThan Operator = iota
	// LessEq shows that the the field being filtered must be less than or equal to the provided value
	LessEq
	// Equal shows that the field being filtered must be equal to the provided value
	Equal
	// NotEqual shows that the field being filtered must be not equal to the provided value
	NotEqual
	// GreaterEq shows that the field being filtered must be greater than or equal to the provided value
	GreaterEq
	// GreaterThan shows that the field being filtered must be greater than the provided value
	GreaterThan
)

func MatchOrder(order ds.Order) (dbr.Condition, error) {
	switch order.OrderType {
	case pb.OrderType_ASK:
		return MatchAsk(order), nil
	case pb.OrderType_BID:
		return MatchBid(order), nil
	default:
		return nil, fmt.Errorf("searching by any type is not supported")
	}
}

func MatchBid(order ds.Order) dbr.Condition {
	cond := IsBidOrder()
	if order.ByuerID != "" {
		cond = dbr.And(cond, BuyerID(order.ByuerID))
	}

	if order.SupplierID != "" {
		cond = dbr.And(cond, SupplierID(order.SupplierID))
	}

	if order.Slot == nil {
		return cond
	}

	slot := order.Slot
	cond = dbr.And(cond,
		GPUCount(LessEq, slot.Resources.GpuCount),
		NetType(LessEq, slot.Resources.NetworkType),
	)

	if slot.Resources.CpuCores > 0 {
		cond = dbr.And(cond, CPUCores(LessEq, slot.Resources.CpuCores))
	}

	if slot.Resources.RamBytes > 0 {
		cond = dbr.And(cond, RamBytes(LessEq, slot.Resources.RamBytes))
	}

	if slot.Resources.Storage > 0 {
		cond = dbr.And(cond, Storage(LessEq, slot.Resources.Storage))
	}

	if slot.Resources.NetTrafficIn > 0 {
		cond = dbr.And(cond, NetTrafficIn(LessEq, slot.Resources.NetTrafficIn))
	}

	if slot.Resources.NetTrafficOut > 0 {
		cond = dbr.And(cond, NetTrafficOut(LessEq, slot.Resources.NetTrafficOut))
	}

	return cond
}

func MatchAsk(order ds.Order) dbr.Condition {
	cond := IsAskOrder()
	if order.ByuerID != "" {
		cond = dbr.And(cond, BuyerID(order.ByuerID))
	}

	if order.SupplierID != "" {
		cond = dbr.And(cond, SupplierID(order.SupplierID))
	}

	if order.Slot == nil {
		return cond
	}

	slot := order.Slot
	cond = dbr.And(cond,
		GPUCount(GreaterEq, slot.Resources.GpuCount),
		NetType(GreaterEq, slot.Resources.NetworkType),
		CPUCores(GreaterEq, slot.Resources.CpuCores),
		RamBytes(GreaterEq, slot.Resources.RamBytes),
		Storage(GreaterEq, slot.Resources.Storage),
		NetTrafficIn(GreaterEq, slot.Resources.NetTrafficIn),
		NetTrafficOut(GreaterEq, slot.Resources.NetTrafficOut),
	)

	return cond
}

func IsAskOrder() dbr.Condition {
	return dbr.Eq("type", pb.OrderType_ASK)
}

func IsBidOrder() dbr.Condition {
	return dbr.Eq("type", pb.OrderType_BID)
}

func BuyerID(ID string) dbr.Condition {
	return dbr.Eq("buyer_id", ID)
}

func SupplierID(ID string) dbr.Condition {
	return dbr.Eq("supplier_id", ID)
}

func CPUCores(op Operator, value uint64) dbr.Condition {
	return op.Condition("resources_cpu_cores", value)
}

func GPUCount(op Operator, value pb.GPUCount) dbr.Condition {
	return op.Condition("resources_gpu_count", value)
}

func RamBytes(op Operator, value uint64) dbr.Condition {
	return op.Condition("resources_ram_bytes", value)
}

func Storage(op Operator, value uint64) dbr.Condition {
	return op.Condition("resources_storage", value)
}

func NetType(op Operator, value pb.NetworkType) dbr.Condition {
	return op.Condition("resources_net_type", value)
}

func NetTrafficIn(op Operator, value uint64) dbr.Condition {
	return op.Condition("resources_net_inbound", value)
}

func NetTrafficOut(op Operator, value uint64) dbr.Condition {
	return op.Condition("resources_net_outbound", value)
}
