package xco

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestWriteMessage(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))
	enc := xml.NewEncoder(b)

	var h Message

	err := enc.Encode(&h)
	if err.Error() != "Malformed Address for Attribute { from}: [Domain is empty]" {
		t.Errorf("Unexpected error encoding message header: %s", err)
		return
	}

	h.From.DomainPart = "example.com"
	h.From.ResourcePart = "home"

	h.To.LocalPart = "goodbye"
	h.To.DomainPart = "example.com"
	h.To.ResourcePart = "home"

	err = enc.Encode(&h)
	if err != nil {
		t.Errorf("Unexpected error encoding message header: %s", err)
		return
	}

}
