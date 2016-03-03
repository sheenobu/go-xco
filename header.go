package xco

// Header contains the common fields for every XEP-0114 message
type Header struct {
	ID   string `xml:"id,attr,omitempty"`
	From string `xml:"from,attr"` //TODO: make address type
	To   string `xml:"to,attr"`   //TODO: make address type
}
