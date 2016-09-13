package xco

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
)

func (c *Component) handshakeState() (st stateFn, err error) {

	if _, err = c.conn.Write([]byte(fmt.Sprintf(`
		<stream:stream
			xmlns='jabber:component:accept'
			xmlns:stream='http://etherx.jabber.org/streams'
			to='%s'>`, c.name))); err != nil {
		return
	}

	for {
		var t xml.Token
		if t, err = c.dec.Token(); err != nil {
			return
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
			err = errors.New("Unable to find ID in stream response")
		}

		handshakeInput := id + c.sharedSecret
		handshake := sha1.Sum([]byte(handshakeInput))
		hexHandshake := hex.EncodeToString(handshake[:])
		if _, err = c.conn.Write([]byte(fmt.Sprintf("<handshake>%s</handshake>", hexHandshake))); err != nil {
			return
		}

		//TODO: separate each step into a state

		// get handshake response
		t, err = c.dec.Token()
		if err != nil {
			return
		}

		t, err = c.dec.Token()
		if err != nil {
			return
		}

		st = c.readLoopState
		return
	}
}
