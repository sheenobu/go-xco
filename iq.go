package xco

// Iq represents an info/query message
type Iq struct {
	Header

	Type string `xml:"type,attr"`

	Content string `xml:",innerxml"`

	XMLName string `xml:"iq"`
}

// IqHandler handles Iq messages
type IqHandler func(c *Component, iq *Iq) error

func noOpIqHandler(c *Component, iq *Iq) error {
	return nil
}
