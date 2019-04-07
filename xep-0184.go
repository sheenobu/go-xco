package xco

// ReceiptAck represents an acknowledgement that the message with ID
// has been received.
type ReceiptAck struct {
	ID string `xml:"id,attr"`
}
