package xco

type Iq struct {
	Header

	Type string `xml:"type,attr"`

	Content string `xml:",innerxml"`

	XMLName string `xml:"iq"`
}

type IqHandler func(c *Component, iq *Iq) error

func noOpIqHandler(c *Component, iq *Iq) error {
	return nil
}
