package xco

// Header contains the common fields for every XEP-0114 message
type Header struct {
	ID   string   `xml:"id,attr,omitempty"`
	From *Address `xml:"from,attr,omitempty"`
	To   *Address `xml:"to,attr,omitempty"`
}
