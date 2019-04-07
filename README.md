# go-xco

[![GoDoc](https://godoc.org/github.com/sheenobu/go-xco?status.svg)](https://godoc.org/github.com/sheenobu/go-xco)

Library for building XMPP/Jabber ([XEP-0114](http://xmpp.org/extensions/xep-0114.html)) components in golang.

## XEPs

Additional XEPs have been added through forks and cherry-picked into
a downstream branch. If you've got additional XEPs to add, create a pull
request or issue to list it:

 * XEP-0114 - XMPP Components
 * XEP-0172 - User Nicknames [mndrix/go-xco](https://github.com/mndrix/go-xco)
 * XEP-0184 - Message Delivery Receipt [mndrix/go-xco](https://github.com/mndrix/go-xco)
 * XEP-0030 - Service Discovery [mndrix/go-xco](https://github.com/mndrix/go-xco)
 * XEP-0054 - vCard [mndrix/go-xco](https://github.com/mndrix/go-xco)

## Usage:

	import (
		"github.com/sheenobu/go-xco"
	)

	func main(){

		opts := xco.Options{
			Name:         Name,
			SharedSecret: SharedSecret,
			Address:      Address,
		}

		c, err := xco.NewComponent(opts)
		if err != nil {
			panic(err)
		}

		// Uppercase Echo Component
		c.MessageHandler = xco.BodyResponseHandler(func(msg *xco.Message) (string, error) {
			return strings.ToUpper(msg.Body), nil
		})
		
		c.Run()
	}

## Raw Usage

The various handlers take the arguments of Component and either Message, Iq, Presence, etc.

You can work with the messages directly without a helper function:

	// Uppercase Echo Component
	c.MessageHandler = func(c *xco.Component, msg *xco.Message) error {
		resp := xco.Message{
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

		return c.Send(&resp)
	}

