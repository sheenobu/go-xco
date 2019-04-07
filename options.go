package xco

import (
	"log"

	"context"
)

// Options define the series of options required to build a component
type Options struct {

	// Name defines the component name
	Name string

	// SharedSecret is the secret shared between the server and component
	SharedSecret string

	// Address is the address of the XMPP server
	Address string

	// The (optional) parent context
	Context context.Context

	// Logger is an optional logger to which to send raw XML stanzas
	// sent and received.  It's primarily intended for debugging and
	// development.
	Logger *log.Logger
}

// NewComponent creates a new component from the given options
func NewComponent(opts Options) (*Component, error) {

	if opts.Context == nil {
		opts.Context = context.Background()
	}

	var c Component
	c.ctx, c.cancelFn = context.WithCancel(opts.Context)

	if err := c.init(opts); err != nil {
		return nil, err
	}

	return &c, nil
}
