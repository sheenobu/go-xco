package xco

const (
	SUBSCRIBE    = "subscribe"
	SUBSCRIBED   = "subscribed"
	UNSUBSCRIBE  = "unsubscribe"
	UNSUBSCRIBED = "unsubscribed"
	UNAVAILABLE  = "unavailable"
	PROBE        = "probe"
)

type Presence struct {
	Header

	Show     string `xml:"show"`
	Status   string `xml:"status"`
	Priority byte   `xml:"priority"`

	Type string `xml:"type"`

	XMLName string `xml:"presence"`
}

type PresenceHandler func(c *Component, p *Presence) error

func noOpPresenceHandler(c *Component, p *Presence) error {
	return nil
}
