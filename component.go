package xco

import (
	"encoding/xml"
	"fmt"
	"net"
	"os"

	"golang.org/x/net/context"
)

type stateFn func() (stateFn, error)

// A Component is an instance of a Jabber Component (XEP-0114)
type Component struct {
	MessageHandler  MessageHandler
	PresenceHandler PresenceHandler
	IqHandler       IqHandler

	ctx      context.Context
	cancelFn context.CancelFunc

	conn net.Conn
	dec  *xml.Decoder
	enc  *xml.Encoder

	stateFn stateFn

	sharedSecret string
	name         string
}

func (c *Component) dial(o *Options) error {
	conn, err := net.Dial("tcp", o.Address)
	if err != nil {
		return err
	}

	c.MessageHandler = noOpMessageHandler
	c.PresenceHandler = noOpPresenceHandler
	c.IqHandler = noOpIqHandler

	c.conn = conn
	c.name = o.Name
	c.sharedSecret = o.SharedSecret
	c.dec = xml.NewDecoder(conn)
	c.enc = xml.NewEncoder(conn)
	c.stateFn = c.handshakeState

	return nil
}

func (c *Component) Close() {
	c.cancelFn()
}

func (c *Component) Run() {

	defer func() {
		c.conn.Close()
	}()

	var err error

	for {
		if c.stateFn == nil {
			return
		}
		c.stateFn, err = c.stateFn()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
	}
}

func (c *Component) Send(i interface{}) error {
	return c.enc.Encode(i)
}
