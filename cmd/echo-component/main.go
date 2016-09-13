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
	c.MessageHandler = xco.BodyResponseHandler(func(msg *xco.Message) (string, error) {
		return strings.ToUpper(msg.Body), nil
	})

	c.PresenceHandler = xco.AlwaysOnlinePresenceHandler

	if err := c.Run(); err != nil {
		panic(err)
	}
}
