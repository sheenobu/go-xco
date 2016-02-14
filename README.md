# go-xco

Library for building XMPP/Jabber ( XEP-0114 ) components in golang.

## Usage:

import (
	"github.com/sheenobu/go-xco"
)

func main(){

	opts := &xco.Options{
		Name:         Name,
		SharedSecret: SharedSecret,
		Address:      Address,
	}

	c, err := opts.NewComponent()
	if err != nil {
		panic(err)
	}

	// Uppercase Echo Component
	c.MessageHandler = func(c *xco.Component, msg *xco.Message) error {
		m := xco.Message{
			Header: xco.Header{
				From: msg.To,
				To:   msg.From,
				ID:   msg.ID,
			},
			Subject: msg.Subject,
			Thread:  msg.Thread,
			Type:    msg.Type,
			Body:    strings.ToUpper(msg.Body),
			XMLName: msg.XMLName,
		}

		return c.Send(m)
	}

	c.Run()
}



