package adaptor

import (
	"github.com/sonm-io/marketplace/infra/cqrs"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// ToDomain makes cqrs.Handler compatible with domain.CommandHandler.
func ToDomain(h cqrs.Handler) intf.CommandHandler {
	return domainAdaptor{h: h}
}

type domainAdaptor struct {
	h cqrs.Handler
}

func (a domainAdaptor) Handle(c intf.Command) error {
	return a.h.Handle(c.(cqrs.Command))
}

// FromDomain makes domain.CommandHandler compatible with cqrs.Handler.
func FromDomain(h intf.CommandHandler) cqrs.Handler {
	return infraAdapter{h: h}
}

type infraAdapter struct {
	h intf.CommandHandler
}

func (a infraAdapter) Handle(c cqrs.Command) error {
	return a.h.Handle(c.(intf.Command))
}
