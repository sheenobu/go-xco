package xco

import "encoding/xml"

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
	Thread  string `xml:"thread,omitempty"`

	XMLName xml.Name
}

// A MessageHandler handles a message
type MessageHandler func(*Component, *Message) error

func noOpMessageHandler(c *Component, m *Message) error {
	return nil
}
