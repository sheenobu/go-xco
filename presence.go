package xco

import (
	"github.com/pkg/errors"
)

const (

	// SUBSCRIBE represents the subscribe Presence message type
	SUBSCRIBE = "subscribe"

	// SUBSCRIBED represents the subscribed Presence message type
	SUBSCRIBED = "subscribed"

	// UNSUBSCRIBE represents the unsubsribe Presence message type
	UNSUBSCRIBE = "unsubscribe"

	// UNSUBSCRIBED represents the unsubsribed Presence message type
	UNSUBSCRIBED = "unsubscribed"

	// UNAVAILABLE represents the unavailable Presence message type
	UNAVAILABLE = "unavailable"

	// PROBE represents the probe Presence message type
	PROBE = "probe"
)

// Presence represents a message identifying whether an entity is available and the subscription requests/responses for the entity
type Presence struct {
	Header

	Show     string `xml:"show"`
	Status   string `xml:"status"`
	Priority byte   `xml:"priority"`

	Type string `xml:"type,attr"`

	XMLName string `xml:"presence"`
}

// PresenceHandler handles incoming presence requests
type PresenceHandler func(c *Component, p *Presence) error

func noOpPresenceHandler(c *Component, p *Presence) error {
	return nil
}

// AlwaysOnlinePresenceHandler always returns "subscribed" to any presence requests
func AlwaysOnlinePresenceHandler(c *Component, p *Presence) error {
	resp := &Presence{
		Header: Header{
			From: p.To,
			To:   p.From,
			ID:   p.ID,
		},
		Type: "subscribed",
	}

	return errors.Wrap(c.Send(resp), "Error sending always online presence")
}

// ToAddressPresenceHandler calls the function with the To address
func ToAddressPresenceHandler(fn func(subject Address) error) PresenceHandler {
	return func(c *Component, p *Presence) error {
		return fn(*p.To)
	}
}
