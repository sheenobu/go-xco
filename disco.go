package xco

import "encoding/xml"

// DiscoInfoHandler handles an incoming service discovery request.
// The target entity is described by iq.To. This function's first
// return value is a slice of identities for the target entity.  The
// second is a slice of features offered by the target entity.
//
// You don't have to include a feature indicating support for service
// discovery, one is automatically added for you.  If you return 0
// identity elements, service discovery is disabled.
type DiscoInfoHandler func(c *Component, iq *Iq) ([]DiscoIdentity, []DiscoFeature, error)

func noOpDiscoInfoHandler(c *Component, iq *Iq) ([]DiscoIdentity, []DiscoFeature, error) {
	return nil, nil, nil
}

const discoInfoSpace = `http://jabber.org/protocol/disco#info`

// DiscoInfoQuery represents a service discovery info query message.
// See section 3.1 in XEP-0030.
type DiscoInfoQuery struct {
	Identities []DiscoIdentity
	Features   []DiscoFeature

	XMLName xml.Name `xml:"http://jabber.org/protocol/disco#info query"`
}

// returns true if this is a valid service discovery info query that
// meets the requirements of XEP-0030 section 3.1
func (q *DiscoInfoQuery) isValid() bool {
	return len(q.Identities) == 0 &&
		len(q.Features) == 0 &&
		q.XMLName.Local == "query" &&
		q.XMLName.Space == discoInfoSpace
}

// IsDiscoInfo returns true if an iq stanza is a service discovery
// info query.
func (iq *Iq) IsDiscoInfo() bool {
	if iq.Type == "get" {
		var disco DiscoInfoQuery
		err := xml.Unmarshal([]byte(iq.Content), &disco)
		return err == nil && disco.isValid()
	}
	return false
}

// DiscoIdentity represents an identity element in a response to a
// service discovery info query.
type DiscoIdentity struct {
	// Category is a mandatory description of the category of this
	// identity.  Category often contains values like "conference",
	// "directory", "gateway", "server", "client", etc.
	//
	// See the category registry at http://xmpp.org/registrar/disco-categories.html
	Category string `xml:"category,attr"`

	// Type is a mandatory description of the type of this identity.
	// The type goes together with the Category to help requesting
	// entities know which services are offered.
	//
	// For example, if Category is "gateway" then Type might be "msn"
	// or "aim".  See the type registry at http://xmpp.org/registrar/disco-categories.html
	Type string `xml:"type,attr"`

	// Name is an optional natural language name for the entity
	// described by this identity.
	Name string `xml:"name,attr,omitempty"`

	XMLName string `xml:"identity"`
}

// DiscoFeature represents a feature element in a response to a
// service discovery info query.
//
// See the registry of features at http://xmpp.org/registrar/disco-features.html
type DiscoFeature struct {
	// Var is a mandatory protocol namespace offered by the entity.
	Var string `xml:"var,attr"`

	XMLName string `xml:"feature"`
}
