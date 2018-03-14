package dialogflow

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Fulfillment is the response sent back to dialogflow in case of a successful
// webhook call
type Fulfillment struct {
	FulfillmentText     string      `json:"fulfillmentText,omitempty"`
	FulfillmentMessages Messages    `json:"fulfillmentMessages,omitempty"`
	Source              string      `json:"source,omitempty"`
	Payload             interface{} `json:"payload,omitempty"`
	OutputContexts      Contexts    `json:"outputContexts,omitempty"`
	FollowupEventInput  interface{} `json:"followupEventInput,omitempty"`
}

// Messages is a simple slice of Message
type Messages []Message

// RichMessage is an interface used in the Message type.
// It is used to send back payloads to dialogflow
type RichMessage interface {
	GetKey() string
}

// Message is a struct holding a platform and a RichMessage.
// Used in the FulfillmentMessages of the response sent back to dialogflow
type Message struct {
	Platform
	RichMessage RichMessage
}

// MarshalJSON implements the Marshaller interface for the JSON type.
// Custom marshalling is necessary since there can only be one rich message
// per Message and the key associated to each type is dynamic
func (m *Message) MarshalJSON() ([]byte, error) {
	var err error
	var b []byte
	buffer := bytes.NewBufferString("{")
	if m.Platform != "" {
		buffer.WriteString(fmt.Sprintf(`"platform": "%s"`, m.Platform))
	}
	if m.Platform != "" && m.RichMessage != nil {
		buffer.WriteString(", ")
	}
	if m.RichMessage != nil {
		if b, err = json.Marshal(m.RichMessage); err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf(`"%s": %s`, m.RichMessage.GetKey(), string(b)))
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

// ForGoogle takes a rich message wraps it in a message with the appropriate
// platform set
func ForGoogle(r RichMessage) Message {
	return Message{
		Platform:    ActionsOnGoogle,
		RichMessage: r,
	}
}
