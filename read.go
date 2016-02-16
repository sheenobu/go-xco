package xco

import "encoding/xml"

// UnknownElementHandler handles unknown XML entities sent through XMPP
type UnknownElementHandler func(*Component, *xml.StartElement) error

func noOpUnknownHandler(c *Component, x *xml.StartElement) error {
	return nil
}

func (c *Component) readLoopState() (stateFn, error) {

	t, err := c.dec.Token()
	if err != nil {
		return nil, err
	}

	if st, ok := t.(xml.StartElement); ok {

		if st.Name.Local == "message" {
			var m Message
			if err := c.dec.DecodeElement(&m, &st); err != nil {
				return nil, err
			}

			if err := c.MessageHandler(c, &m); err != nil {
				return nil, err
			}
		} else if st.Name.Local == "presence" {
			var p Presence
			if err := c.dec.DecodeElement(&p, &st); err != nil {
				return nil, err
			}

			if err := c.PresenceHandler(c, &p); err != nil {
				return nil, err
			}
		} else if st.Name.Local == "iq" {

			var iq Iq
			if err := c.dec.DecodeElement(&iq, &st); err != nil {
				return nil, err
			}

			if err := c.IqHandler(c, &iq); err != nil {
				return nil, err
			}
		} else {
			if err := c.UnknownHandler(c, &st); err != nil {
				return nil, err
			}
		}
	}

	return c.readLoopState, nil
}
