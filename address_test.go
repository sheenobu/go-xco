package xco

import "testing"

var parseAddressTests = []struct {
	Input  string
	Output *Address
	Error  string
}{
	{"", nil, "Address is empty"},
	{"           ", nil, "Address is empty"},
	{"@example.com", nil, "Localpart is empty"},
	{"example.com/", nil, "Resourcepart is empty"},
	{"@example.com/", nil, "Multiple errors: [Localpart is empty Resourcepart is empty]"},
	{"@/", nil, "Multiple errors: [Domain is empty Localpart is empty Resourcepart is empty]"},

	{"example.com", &Address{"", "example.com", ""}, ""},
	{"hello@example.com", &Address{"hello", "example.com", ""}, ""},
	{"example.com/home", &Address{"", "example.com", "home"}, ""},
	{"hello@example.com/home", &Address{"hello", "example.com", "home"}, ""},

	{"goodbye@example.com/home", &Address{"goodbye", "example.com", "home"}, ""},
}

func TestParseAddress(t *testing.T) {
	for _, pat := range parseAddressTests {
		addr, err := ParseAddress(pat.Input)
		matches := addr.Equals(pat.Output) && (err == nil && pat.Error == "" || err.Error() == pat.Error)
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
