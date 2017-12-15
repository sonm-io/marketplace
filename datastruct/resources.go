package datastruct

// Resources are a set of computational units.
type Resources struct {
	// CPU core count.
	CPUCores uint64
	// RAM, in bytes.
	RAMBytes uint64
	// GPU devices count.
	GPUCount GPUCount
	// storage volume, in Megabytes.
	Storage uint64
	// Inbound network traffic (the higher value), in bytes.
	NetTrafficIn uint64
	// Outbound network traffic (the higher value), in bytes.
	NetTrafficOut uint64
	// Allowed network connections.
	NetworkType NetworkType
	// Other properties/benchmarks. The higher, the better.
	Properties map[string]float64
}

// GPUCount defines GPU computation options.
type GPUCount int32

// List of possible GPU computation options.
const (
	NoGPU       GPUCount = 0
	SingleGPU            = 1
	MultipleGPU          = 2
)

// NetworkType defines available network type.
type NetworkType int32

// List of possible network types.
const (
	NoNetwork NetworkType = 0
	Outbound              = 1
	Inbound               = 2
)
