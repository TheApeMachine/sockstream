package sockstream

import "time"

// Message wraps the data sent over the websocket in a structured format.
type Message struct {
	ts int64
	im []byte
}

// NewMessage returns a referece to a new instance of Message.
func NewMessage(im []byte) *Message {
	return &Message{
		ts: time.Now().UnixNano(),
		im: im,
	}
}
