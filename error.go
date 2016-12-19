package xco

import (
	"encoding/xml"
	"fmt"
)

// Error is the error sent over XMPP
type Error struct {
	XMLName xml.Name

	Code string `xml:"code,omitempty,attr"`
	Type string `xml:"type,omitempty,attr"`
}

func (e *Error) String() string {
	return fmt.Sprintf("Error{code='%s' type='%s'}", e.Code, e.Type)
}
