package xco

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

// MessageType defines the constants for the types of messages within XEP-0114
type MessageType string

const (

	// CHAT defines the chat message type
	CHAT = MessageType("chat")

	// ERROR defines the error message type
	ERROR = MessageType("error")

	// GROUPCHAT defines the group chat message type
	GROUPCHAT = MessageType("groupchat")

	// HEADLINE defines the headline message type
	HEADLINE = MessageType("headline")

	// NORMAL defines the normal message type
	NORMAL = MessageType("normal")
)

// A Message is an incoming or outgoing Component message
type Message struct {
	Header
	Type MessageType `xml:"type,attr,omitempty"`

	Subject string `xml:"subject,omitempty"`
	Body    string `xml:"body"`
	Error   *Error `xml:"error"`
	Thread  string `xml:"thread,omitempty"`

	XMLName xml.Name
}

// A MessageHandler handles an incoming message
type MessageHandler func(*Component, *Message) error

func noOpMessageHandler(c *Component, m *Message) error {
	return nil
}

// BodyResponseHandler builds a simple request-response style function which returns the body
// of the response message
func BodyResponseHandler(fn func(*Message) (string, error)) MessageHandler {
	return func(c *Component, m *Message) error {

		body, err := fn(m)
		if err != nil {
			return err
		}

		resp := &Message{
			Header: Header{
				From: m.To,
				To:   m.From,
				ID:   m.ID,
			},
			Subject: m.Subject,
			Thread:  m.Thread,
			Type:    m.Type,
			Body:    body,
			XMLName: m.XMLName,
		}

		return errors.Wrap(c.Send(resp), "Error sending message response")
	}
}
