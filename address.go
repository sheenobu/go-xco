package xco

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Address is an XMPP JID address
type Address struct {
	LocalPart    string
	DomainPart   string
	ResourcePart string
}

// ParseAddress parses the address from the given string
func ParseAddress(s string) (Address, error) {
	var addr Address
	err := addr.parse(s)
	return addr, err
}

// Equals compares the given address
func (a *Address) Equals(o *Address) bool {
	return (a == o) || ((a != nil && o != nil) && (a.LocalPart == o.LocalPart && a.DomainPart == o.DomainPart && a.ResourcePart == o.ResourcePart))
}

// String formats the address as an XMPP JID
func (a *Address) String() string {
	buf := bytes.NewBufferString("")
	if a.LocalPart != "" {
		buf.WriteString(a.LocalPart)
		buf.WriteString("@")
	}

	buf.WriteString(a.DomainPart)

	if a.ResourcePart != "" {
		buf.WriteString("/")
		buf.WriteString(a.ResourcePart)
	}

	return buf.String()
}

// Bare returns a copy of this address with the resource part made
// blank.
func (a *Address) Bare() *Address {
	b := *a
	b.ResourcePart = ""
	return &b
}

// UnmarshalXMLAttr marks the Address struct as being able to be parsed as an XML attribute
func (a *Address) UnmarshalXMLAttr(attr xml.Attr) error {
	return a.parse(attr.Value)
}

// MarshalXMLAttr marks the Address struct as being able to be written as an XML attribute
func (a *Address) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if a == nil {
		return xml.Attr{}, nil
	}
	errs := a.validate()
	if len(errs) != 0 {
		return xml.Attr{}, &multiError{
			errs: errs,
			mainError: errors.Errorf(
				"Malformed Address for Attribute %s", name),
		}
	}

	return xml.Attr{
		Name:  name,
		Value: a.String(),
	}, nil
}

func (a *Address) validate() []error {

	var errs []error
	if a != nil && a.LocalPart != "" && a.DomainPart == "" {
		errs = append(errs, errors.New("Domain is empty"))
	}

	return errs
}

func (a *Address) parse(s string) error {

	// normalization

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return nil //errors.New("Address is empty")
	}

	// parsing

	domainStart := 0
	domainEnd := len(s)

	if idx := strings.IndexAny(s, "@"); idx != -1 {
		a.LocalPart = s[0:idx]
		domainStart = idx + 1
	}

	if idx := strings.IndexAny(s, "/"); idx != -1 {
		a.ResourcePart = s[idx+1:]
		domainEnd = idx
	}

	if domainStart != domainEnd {
		a.DomainPart = s[domainStart:domainEnd]
	}

	// validation

	errs := a.validate()

	if a.LocalPart == "" && domainStart != 0 {
		errs = append(errs, errors.New("Localpart is empty"))
	}

	if a.ResourcePart == "" && domainEnd != len(s) {
		errs = append(errs, errors.New("Resourcepart is empty"))
	}

	if len(errs) == 1 {
		return errs[0]
	} else if len(errs) > 1 {
		return fmt.Errorf("Multiple errors: %v", errs)
	}

	return nil
}
