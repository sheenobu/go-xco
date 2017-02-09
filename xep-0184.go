package xco

// ReceiptAck represents an acknowledgement that the message with Id
// has been received.
type ReceiptAck struct {
	Id string `xml:"id,attr"`
}
