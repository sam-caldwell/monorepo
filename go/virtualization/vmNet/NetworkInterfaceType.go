package vmNet

type InterfaceType uint

const (
	Undefined InterfaceType = iota
	Bridge
	NAT
)
