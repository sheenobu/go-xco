package xco

import (
	"encoding/xml"
	"testing"

	"github.com/pkg/errors"
)

var nilAddress = Address{}
var withDomain = Address{DomainPart: "example.com"}

var parseAddressTests = []struct {
	Input  string
	Output Address
	Error  string
}{
	{"", nilAddress, ""},
	{"           ", nilAddress, ""},
	{"@example.com", withDomain, "Localpart is empty"},
	{"example.com/", withDomain, "Resourcepart is empty"},
	{"@example.com/", withDomain, "Multiple errors: [Localpart is empty Resourcepart is empty]"},
	{"@/", nilAddress, "Multiple errors: [Localpart is empty Resourcepart is empty]"},

	{"example.com", Address{"", "example.com", ""}, ""},
	{"hello@example.com", Address{"hello", "example.com", ""}, ""},
	{"example.com/home", Address{"", "example.com", "home"}, ""},
	{"hello@example.com/home", Address{"hello", "example.com", "home"}, ""},

	{"goodbye@example.com/home", Address{"goodbye", "example.com", "home"}, ""},
}

func TestParseAddress(t *testing.T) {
	for _, pat := range parseAddressTests {
		addr, err := ParseAddress(pat.Input)
		matches := addr.Equals(&pat.Output) && (err == nil && pat.Error == "" || err.Error() == pat.Error)
		if !matches {
			t.Errorf("ParseAddress(%s) => {%s,%v}, expected {%s,%s}",
				pat.Input, addr, err, pat.Output, pat.Error)
		}
	}
}

var stringAddressTests = []struct {
	Input  *Address
	Output string
}{
	{&Address{"", "example.com", ""}, "example.com"},
	{&Address{"hello", "example.com", ""}, "hello@example.com"},
	{&Address{"", "example.com", "home"}, "example.com/home"},
	{&Address{"hello", "example.com", "home"}, "hello@example.com/home"},
	{&Address{"goodbye", "example.com", "home"}, "goodbye@example.com/home"},
}

func TestStringAddress(t *testing.T) {
	for _, sat := range stringAddressTests {
		out := sat.Input.String()
		matches := out == sat.Output
		if !matches {
			t.Errorf("%v.String() => {%s}, expected {%s}",
				sat.Input, out, sat.Output)
		}
	}
}

func attrEquals(a *xml.Attr, b *xml.Attr) bool {
	if a == b {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	if a.Name.Space != b.Name.Space {
		return false
	}
	if a.Name.Local != b.Name.Local {
		return false
	}
	if a.Value != b.Value {
		return false
	}

	return true
}

var marshallXMLAddressTests = []struct {
	Input  *Address
	Output xml.Attr
	Error  string
}{
	{nil, xml.Attr{}, ""},
	{&Address{"", "example.com", ""},
		xml.Attr{Name: xml.Name{}, Value: "example.com"}, ""},
	{&Address{"asdf", "", ""},
		xml.Attr{}, "[Domain is empty]"},
}

func TestMarshallXMLAddress(t *testing.T) {
	for _, mat := range marshallXMLAddressTests {
		out, err := mat.Input.MarshalXMLAttr(xml.Name{})
		matches := attrEquals(&out, &mat.Output) && (err == nil && mat.Error == "" || errors.Cause(err).Error() == mat.Error)
		if !matches {
			t.Errorf("{%s}.MarshalXMLAttr({}) => {%s,%v}, expected {%v,%s}",
				mat.Input, out, errors.Cause(err), mat.Output, mat.Error)
		}

	}
}
