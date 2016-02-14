package xco

type Header struct {
	ID   string `xml:"id,attr"`
	From string `xml:"from,attr"` //TODO: make address type
	To   string `xml:"to,attr"`   //TODO: make address type
}
