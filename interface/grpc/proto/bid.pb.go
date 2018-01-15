// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bid.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	bid.proto
	bigint.proto
	capabilities.proto
	deal.proto
	hub.proto
	insonmnia.proto
	locator.proto
	marketplace.proto
	miner.proto
	node.proto

It has these top-level messages:
	Geo
	Resources
	Slot
	Order
	BigInt
	Capabilities
	CPUDevice
	RAMDevice
	GPUDevice
	Deal
	ListReply
	HubStartTaskRequest
	HubStartTaskReply
	HubStatusReply
	DealRequest
	GetDevicePropertiesReply
	SetDevicePropertiesRequest
	SlotsReply
	GetAllSlotsReply
	AddSlotRequest
	RemoveSlotRequest
	GetRegisteredWorkersReply
	TaskListReply
	CPUDeviceInfo
	GPUDeviceInfo
	DevicesReply
	InsertSlotRequest
	PullTaskRequest
	DealInfoReply
	CompletedTask
	Empty
	ID
	TaskID
	PingReply
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	InfoReply
	TaskStatusReply
	AvailableResources
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	DiscoverHubRequest
	TaskResourceRequirements
	Timestamp
	Chunk
	Progress
	AnnounceRequest
	ResolveRequest
	ResolveReply
	GetOrdersRequest
	GetOrdersReply
	GetProcessingReply
	TouchOrdersRequest
	MinerHandshakeRequest
	MinerHandshakeReply
	MinerStartRequest
	SocketAddr
	MinerStartReply
	TaskInfo
	Route
	MinerStatusMapRequest
	SaveRequest
	TaskListRequest
	DealListRequest
	DealListReply
*/
package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OrderType int32

const (
	OrderType_ANY OrderType = 0
	OrderType_BID OrderType = 1
	OrderType_ASK OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "ANY",
	1: "BID",
	2: "ASK",
}
var OrderType_value = map[string]int32{
	"ANY": 0,
	"BID": 1,
	"ASK": 2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Geo represent GeoIP results for node
type Geo struct {
	Country string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	City    string  `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	Lat     float32 `protobuf:"fixed32,3,opt,name=lat" json:"lat,omitempty"`
	Lon     float32 `protobuf:"fixed32,4,opt,name=lon" json:"lon,omitempty"`
}

func (m *Geo) Reset()                    { *m = Geo{} }
func (m *Geo) String() string            { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()               {}
func (*Geo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Geo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Geo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Geo) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Geo) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Resources struct {
	// CPU core count
	CpuCores uint64 `protobuf:"varint,1,opt,name=cpuCores" json:"cpuCores,omitempty"`
	// RAM, in bytes
	RamBytes uint64 `protobuf:"varint,2,opt,name=ramBytes" json:"ramBytes,omitempty"`
	// GPU devices count
	GpuCount GPUCount `protobuf:"varint,3,opt,name=gpuCount,enum=sonm.GPUCount" json:"gpuCount,omitempty"`
	// todo: discuss
	// storage volume, in Megabytes
	Storage uint64 `protobuf:"varint,4,opt,name=storage" json:"storage,omitempty"`
	// Inbound network traffic (the higher value), in bytes
	NetTrafficIn uint64 `protobuf:"varint,5,opt,name=netTrafficIn" json:"netTrafficIn,omitempty"`
	// Outbound network traffic (the higher value), in bytes
	NetTrafficOut uint64 `protobuf:"varint,6,opt,name=netTrafficOut" json:"netTrafficOut,omitempty"`
	// Allowed network connections
	NetworkType NetworkType `protobuf:"varint,7,opt,name=networkType,enum=sonm.NetworkType" json:"networkType,omitempty"`
	// Other properties/benchmarks. The higher means better.
	Properties map[string]float64 `protobuf:"bytes,8,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resources) GetCpuCores() uint64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Resources) GetRamBytes() uint64 {
	if m != nil {
		return m.RamBytes
	}
	return 0
}

func (m *Resources) GetGpuCount() GPUCount {
	if m != nil {
		return m.GpuCount
	}
	return GPUCount_NO_GPU
}

func (m *Resources) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Resources) GetNetTrafficIn() uint64 {
	if m != nil {
		return m.NetTrafficIn
	}
	return 0
}

func (m *Resources) GetNetTrafficOut() uint64 {
	if m != nil {
		return m.NetTrafficOut
	}
	return 0
}

func (m *Resources) GetNetworkType() NetworkType {
	if m != nil {
		return m.NetworkType
	}
	return NetworkType_NO_NETWORK
}

func (m *Resources) GetProperties() map[string]float64 {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Slot struct {
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	BuyerRating int64 `protobuf:"varint,1,opt,name=buyerRating" json:"buyerRating,omitempty"`
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	SupplierRating int64 `protobuf:"varint,2,opt,name=supplierRating" json:"supplierRating,omitempty"`
	// Geo represent Worker's position
	Geo *Geo `protobuf:"bytes,3,opt,name=geo" json:"geo,omitempty"`
	// Hardware resources requirements
	Resources *Resources `protobuf:"bytes,4,opt,name=resources" json:"resources,omitempty"`
	// Duration is resource rent duration in seconds
	Duration uint64 `protobuf:"varint,5,opt,name=duration" json:"duration,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Slot) GetBuyerRating() int64 {
	if m != nil {
		return m.BuyerRating
	}
	return 0
}

func (m *Slot) GetSupplierRating() int64 {
	if m != nil {
		return m.SupplierRating
	}
	return 0
}

func (m *Slot) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Slot) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Slot) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type Order struct {
	// Order ID, UUIDv4
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Buyer's EtherumID
	ByuerID string `protobuf:"bytes,2,opt,name=byuerID" json:"byuerID,omitempty"`
	// Supplier's is EtherumID
	SupplierID string `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	// Order type (Bid or Ask)
	OrderType OrderType `protobuf:"varint,5,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	// Slot describe resource requiements
	Slot *Slot `protobuf:"bytes,6,opt,name=slot" json:"slot,omitempty"`
	// PricePerSecond specifies order price for ordered resources per second.
	PricePerSecond *BigInt `protobuf:"bytes,7,opt,name=pricePerSecond" json:"pricePerSecond,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetByuerID() string {
	if m != nil {
		return m.ByuerID
	}
	return ""
}

func (m *Order) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *Order) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func (m *Order) GetPricePerSecond() *BigInt {
	if m != nil {
		return m.PricePerSecond
	}
	return nil
}

func init() {
	proto.RegisterType((*Geo)(nil), "sonm.Geo")
	proto.RegisterType((*Resources)(nil), "sonm.Resources")
	proto.RegisterType((*Slot)(nil), "sonm.Slot")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0xcd, 0x47, 0x77, 0x9b, 0x93, 0xda, 0xd6, 0xc1, 0x8b, 0x50, 0x61, 0x2d, 0x45, 0x96,
	0xb2, 0x60, 0x2f, 0xb2, 0x5e, 0x88, 0x20, 0x62, 0xad, 0x2c, 0x45, 0xd8, 0x96, 0xe9, 0x8a, 0x78,
	0x99, 0x26, 0xb3, 0x61, 0xd8, 0xee, 0x4c, 0x98, 0xcc, 0x28, 0x79, 0x03, 0x1f, 0xc9, 0x27, 0xf1,
	0x79, 0x64, 0x4e, 0x9a, 0xb4, 0xdb, 0xbb, 0x39, 0xbf, 0xf3, 0x9f, 0x9c, 0x8f, 0xff, 0x04, 0x82,
	0x2d, 0xcf, 0x66, 0x85, 0x92, 0x5a, 0x12, 0xbf, 0x94, 0xe2, 0x71, 0xd4, 0xdb, 0xf2, 0x9c, 0x0b,
	0x5d, 0xb3, 0xd1, 0x80, 0x0b, 0x4b, 0x05, 0x4f, 0x6a, 0x30, 0xf9, 0x01, 0xde, 0x0d, 0x93, 0x24,
	0x82, 0xf3, 0x54, 0x1a, 0xa1, 0x55, 0x15, 0x39, 0x63, 0x67, 0x1a, 0xd0, 0x26, 0x24, 0x04, 0xfc,
	0x94, 0xeb, 0x2a, 0x72, 0x11, 0xe3, 0x99, 0x0c, 0xc1, 0xdb, 0x25, 0x3a, 0xf2, 0xc6, 0xce, 0xd4,
	0xa5, 0xf6, 0x88, 0x44, 0x8a, 0xc8, 0xdf, 0x13, 0x29, 0x26, 0x7f, 0x3c, 0x08, 0x28, 0x2b, 0xa5,
	0x51, 0x29, 0x2b, 0xc9, 0x08, 0xba, 0x69, 0x61, 0xbe, 0x48, 0xc5, 0x4a, 0x2c, 0xe0, 0xd3, 0x36,
	0xb6, 0x39, 0x95, 0x3c, 0xce, 0x2b, 0xcd, 0x4a, 0xac, 0xe2, 0xd3, 0x36, 0x26, 0x57, 0xd0, 0xcd,
	0xad, 0xce, 0x88, 0xba, 0x5c, 0x3f, 0xee, 0xcf, 0xec, 0x00, 0xb3, 0x9b, 0xf5, 0x77, 0xa4, 0xb4,
	0xcd, 0xdb, 0x19, 0x4a, 0x2d, 0x55, 0x92, 0x33, 0xec, 0xc3, 0xa7, 0x4d, 0x48, 0x26, 0xd0, 0x13,
	0x4c, 0xdf, 0xa9, 0xe4, 0xfe, 0x9e, 0xa7, 0x4b, 0x11, 0x75, 0x30, 0xfd, 0x84, 0x91, 0x37, 0xf0,
	0xfc, 0x10, 0xaf, 0x8c, 0x8e, 0xce, 0x50, 0xf4, 0x14, 0x92, 0x6b, 0x08, 0x05, 0xd3, 0xbf, 0xa5,
	0x7a, 0xb8, 0xab, 0x0a, 0x16, 0x9d, 0x63, 0x4b, 0x2f, 0xea, 0x96, 0x6e, 0x0f, 0x09, 0x7a, 0xac,
	0x22, 0x9f, 0x00, 0x0a, 0x25, 0x0b, 0xa6, 0x34, 0x67, 0x65, 0xd4, 0x1d, 0x7b, 0xd3, 0x30, 0x7e,
	0x5d, 0xdf, 0x69, 0x37, 0x34, 0x5b, 0xb7, 0x8a, 0xaf, 0x76, 0xef, 0xf4, 0xe8, 0xca, 0xe8, 0x23,
	0x0c, 0x4e, 0xd2, 0x76, 0xe1, 0x0f, 0xac, 0x31, 0xcb, 0x1e, 0xc9, 0x4b, 0xe8, 0xfc, 0x4a, 0x76,
	0x86, 0xe1, 0x0e, 0x1d, 0x5a, 0x07, 0x1f, 0xdc, 0xf7, 0xce, 0xe4, 0xaf, 0x03, 0xfe, 0x66, 0x27,
	0x35, 0x19, 0x43, 0xb8, 0x35, 0x15, 0x53, 0x34, 0xd1, 0x5c, 0xe4, 0x78, 0xd9, 0xa3, 0xc7, 0x88,
	0x5c, 0x42, 0xbf, 0x34, 0x45, 0xb1, 0xe3, 0xad, 0xc8, 0x45, 0xd1, 0x09, 0x25, 0xaf, 0xc0, 0xcb,
	0x99, 0x44, 0x4b, 0xc2, 0x38, 0xd8, 0x5b, 0xc2, 0x24, 0xb5, 0x94, 0xbc, 0x85, 0x40, 0x35, 0x73,
	0xa1, 0x15, 0x61, 0x3c, 0x38, 0x19, 0x97, 0x1e, 0x14, 0xd6, 0xff, 0xcc, 0xa8, 0x44, 0x73, 0xd9,
	0x38, 0xd3, 0xc6, 0x93, 0x7f, 0x0e, 0x74, 0x56, 0x2a, 0x63, 0x8a, 0xf4, 0xc1, 0xe5, 0xd9, 0x7e,
	0x5e, 0x97, 0x67, 0xd6, 0xed, 0x6d, 0x65, 0x98, 0x5a, 0x2e, 0xf6, 0x4f, 0xb3, 0x09, 0xc9, 0x05,
	0x40, 0xd3, 0xed, 0x72, 0x81, 0x2d, 0x06, 0xf4, 0x88, 0xd8, 0xf6, 0xa4, 0xfd, 0x24, 0x3a, 0xd8,
	0x41, 0x07, 0xf7, 0xed, 0xad, 0x1a, 0x4c, 0x0f, 0x0a, 0x72, 0x01, 0x7e, 0xb9, 0x93, 0xf5, 0x7b,
	0x08, 0x63, 0xa8, 0x95, 0x76, 0x9d, 0x14, 0x39, 0x79, 0x07, 0xfd, 0x42, 0xf1, 0x94, 0xad, 0x99,
	0xda, 0xb0, 0x54, 0x8a, 0x0c, 0x5f, 0x45, 0x18, 0xf7, 0x6a, 0xe5, 0x9c, 0xe7, 0x4b, 0xa1, 0xe9,
	0x89, 0xe6, 0xea, 0x12, 0x82, 0xb6, 0x1a, 0x39, 0x07, 0xef, 0xf3, 0xed, 0xcf, 0xe1, 0x33, 0x7b,
	0x98, 0x2f, 0x17, 0x43, 0x07, 0xc9, 0xe6, 0xdb, 0xd0, 0xdd, 0x9e, 0xe1, 0x6f, 0x7a, 0xfd, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0xfe, 0x90, 0xe8, 0xd8, 0x03, 0x00, 0x00,
}
