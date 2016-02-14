package xco

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
)

func (c *Component) handshakeState() (stateFn, error) {

	if _, err := c.conn.Write([]byte(fmt.Sprintf(`
		<stream:stream
			xmlns='jabber:component:accept'
			xmlns:stream='http://etherx.jabber.org/streams'
			to='%s'>`, c.name))); err != nil {
		return nil, err
	}

	for {
		t, err := c.dec.Token()
		if err != nil {
			return nil, err
		}

		stream, ok := t.(xml.StartElement)
		if !ok {
			continue
		}

		var id string

		for _, a := range stream.Attr {
			if a.Name.Local == "id" {
				id = a.Value
			}
		}

		if id == "" {
			return nil, errors.New("Unable to find ID in stream response")
		}

		handshakeInput := id + c.sharedSecret
		handshake := sha1.Sum([]byte(handshakeInput))
		hexHandshake := hex.EncodeToString(handshake[:])
		if _, err := c.conn.Write([]byte(fmt.Sprintf("<handshake>%s</handshake>", hexHandshake))); err != nil {
			return nil, err
		}

		//TODO: separate each step into a state

		// get handshake response
		t, err = c.dec.Token()
		if err != nil {
			return nil, err
		}

		t, err = c.dec.Token()
		if err != nil {
			return nil, err
		}

		return c.readLoopState, nil
	}
}
