package xco

import (
	"bytes"
	"encoding/xml"
	"testing"
)

const headerBody = `
<message
id='asdf'
to='hello@example.com'
from='goodbye@example.com/home'>
</message>
`

func TestReadHeader(t *testing.T) {
	input := bytes.NewReader([]byte(headerBody))
	dec := xml.NewDecoder(input)

	var h Header

	err := dec.Decode(&h)
	if err != nil {
		t.Errorf("Unexpected error parsing message header: %s", err)
		return
	}

	if s := h.From.String(); s != "goodbye@example.com/home" {
		t.Errorf("Expected from string to be 'goodbye@example.com/home', is '%s'", s)
	}

	if s := h.To.String(); s != "hello@example.com" {
		t.Errorf("Expected from string to be 'hello@example.com', is '%s'", s)
	}

	if h.From.DomainPart != "example.com" {
		t.Errorf("domain part equals %s, expected %s", "example.com", h.From.DomainPart)
	}

	if h.ID != "asdf" {
		t.Errorf("Expected ID to be 'asdf', is '%s'", h.ID)
	}

}

func TestWriteHeader(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))
	enc := xml.NewEncoder(b)

	var h Header

	err := enc.Encode(&h)
	if err != nil {
		t.Errorf("Unexpected error encoding message header: %s", err)
		return
	}

	//h.From.DomainPart = "example.com"
	//h.From.ResourcePart = "home"

	h.To = &Address{}
	h.To.LocalPart = "goodbye"
	h.To.DomainPart = "example.com"
	h.To.ResourcePart = "home"

	err = enc.Encode(&h)
	if err != nil {
		t.Errorf("Unexpected error encoding message header: %s", err)
		return
	}

}
