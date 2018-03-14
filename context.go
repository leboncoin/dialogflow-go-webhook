package dialogflow

import "encoding/json"

// Context is a context contained in a query
type Context struct {
	Name          string          `json:"name,omitempty"`
	LifespanCount int             `json:"lifespanCount,omitempty"`
	Parameters    json.RawMessage `json:"parameters,omitempty"`
}

// Contexts is a slice of pointer to Context
type Contexts []*Context
