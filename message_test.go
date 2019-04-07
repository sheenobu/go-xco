package xco

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestWriteMessage(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))
	enc := xml.NewEncoder(b)

	var h Message

	err := enc.Encode(&h)
	if err != nil {
		t.Errorf("Unexpected error encoding message header: %s", err)
		return
	}

	h.From = &Address{}
	h.From.DomainPart = "example.com"
	h.From.ResourcePart = "home"

	h.To = &Address{}
	h.To.LocalPart = "goodbye"
	h.To.DomainPart = "example.com"
	h.To.ResourcePart = "home"

	h.Content = "<hello/>"
	h.Body = "hello"

	err = enc.Encode(&h)
	if err != nil {
		t.Errorf("Unexpected error encoding message header: %s", err)
		return
	}

	fmt.Println(b.String())

}
