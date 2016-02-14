package main

import (
	"flag"
	"os"
	"strings"

	"github.com/sheenobu/go-xco"
)

var Name string
var SharedSecret string
var Address string

func init() {
	flag.StringVar(&Name, "name", "", "Name of Component")
	flag.StringVar(&SharedSecret, "secret", "", "Shared Secret between server and component")
	flag.StringVar(&Address, "address", "", "Hostname:port address of XMPP component listener")
}

func main() {

	flag.Parse()

	if Name == "" || SharedSecret == "" || Address == "" {
		flag.Usage()
		os.Exit(-1)
		return
	}

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
