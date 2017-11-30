package datastruct

// Resources are a set of computational units.
type Resources struct {
	// CPU core count
	CPUCores uint64
	// RAM, in bytes
	RAMBytes uint64
	// GPU devices count
	////GpuCount GPUCount
	// todo: discuss
	// storage volume, in Megabytes
	Storage uint64
	// Inbound network traffic (the higher value), in bytes
	NetTrafficIn uint64
	// Outbound network traffic (the higher value), in bytes
	NetTrafficOut uint64
	// Allowed network connections
	//NetworkType NetworkType `protobuf:"varint,7,opt,name=networkType,enum=sonm.NetworkType" json:"networkType,omitempty"`
	// Other properties/benchmarks. The higher means better.
	Properties map[string]float64
}
