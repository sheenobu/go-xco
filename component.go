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
	UnknownHandler  UnknownElementHandler

	ctx      context.Context
	cancelFn context.CancelFunc

	conn net.Conn
	dec  *xml.Decoder
	enc  *xml.Encoder

	stateFn stateFn

	sharedSecret string
	name         string
}

func (c *Component) init(o *Options) error {
	conn, err := net.Dial("tcp", o.Address)
	if err != nil {
		return err
	}

	c.MessageHandler = noOpMessageHandler
	c.PresenceHandler = noOpPresenceHandler
	c.IqHandler = noOpIqHandler
	c.UnknownHandler = noOpUnknownHandler

	c.conn = conn
	c.name = o.Name
	c.sharedSecret = o.SharedSecret
	c.dec = xml.NewDecoder(conn)
	c.enc = xml.NewEncoder(conn)
	c.stateFn = c.handshakeState

	return nil
}

// Close closes the Component
func (c *Component) Close() {
	if c == nil {
		return
	}
	c.cancelFn()
}

// Run runs the component handlers loop and waits for it to finish
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
			fmt.Fprintf(os.Stderr, "Error: %s\n", err) //TODO: implement error logging via a handler
		}
	}
}

// Send sends the given pointer struct by serializing it to XML.
func (c *Component) Send(i interface{}) error {
	return c.enc.Encode(i)
}

// Write implements the io.Writer interface to allow direct writing to the XMPP connection
func (c *Component) Write(b []byte) (int, error) {
	return c.conn.Write(b)
}
