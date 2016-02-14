package xco

import "encoding/xml"

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

			/*
				if iq.Type == "get" {

					iqResp := &Iq{
						Header: Header{
							From: iq.To,
							To:   iq.From,
							ID:   iq.ID,
						},
						Type: "result",
						Content: `
						  <vCard xmlns='vcard-temp'>
								<FN>Sheena Artrip</FN>
						  </vCard>
							`,
					}

					c.enc.Encode(&iqResp)
				}
			*/
		}
	}

	return c.readLoopState, nil
}
