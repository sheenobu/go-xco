package xco

import (
	"encoding/xml"
	"fmt"

	"github.com/pkg/errors"
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

type multiError struct {
	errs      []error
	mainError error
}

func (m *multiError) Error() string {
	return fmt.Sprintf("%s: %s", m.mainError.Error(), m.errs)
}

func (m *multiError) Errors() []error {
	return m.errs
}

func (m *multiError) Cause() error {
	//FIXME: should "cause" be the mainError or the group of errors?
	return errors.Errorf("%s", m.errs)
}
