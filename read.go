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

			// recognize XEP-0030 service discovery info queries
			if iq.IsDiscoInfo() {
				return c.discoInfo(&iq)
			}

			// handle all other iq stanzas
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

func (c *Component) discoInfo(iq *Iq) (stateFn, error) {
	ids, features, err := c.DiscoInfoHandler(c, iq)
	if err != nil {
		return nil, err
	}
	if len(ids) < 1 {
		return c.readLoopState, nil
	}

	features = append(features, DiscoFeature{
		Var: discoInfoSpace,
	})
	query := DiscoInfoQuery{
		Identities: ids,
		Features:   features,
	}
	queryContent, err := xml.Marshal(query)
	if err != nil {
		return nil, err
	}
	resp := &Iq{
		Header: Header{
			From: iq.To,
			To:   iq.From,
			ID:   iq.ID,
		},
		Type:    "result",
		Content: string(queryContent),
		XMLName: iq.XMLName,
	}
	if err := c.Send(resp); err != nil {
		return nil, err
	}

	return c.readLoopState, nil
}
