package dialogflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Response is the top-level struct holding all the information
// Basically links a response ID with a query result.
type Response struct {
	Session                     string          `json:"session,omitempty"`
	ResponseID                  string          `json:"responseId,omitempty"`
	QueryResult                 QueryResult     `json:"queryResult,omitempty"`
	OriginalDetectIntentRequest json.RawMessage `json:"originalDetectIntentRequest,omitempty"`
}

// GetParams simply unmarshals the parameters to the given struct and returns
// an error if it's not possible
func (rw *Response) GetParams(i interface{}) error {
	return json.Unmarshal(rw.QueryResult.Parameters, &i)
}

// GetContext allows to search in the output contexts of the query
func (rw *Response) GetContext(ctx string, i interface{}) error {
	for _, c := range rw.QueryResult.OutputContexts {
		if strings.HasSuffix(c.Name, ctx) {
			return json.Unmarshal(c.Parameters, &i)
		}
	}
	return errors.New("context not found")
}

// NewContext is a helper function to create a new named context with params
// name and a lifespan
func (rw *Response) NewContext(name string, lifespan int, params interface{}) (*Context, error) {
	var err error
	var b []byte

	if b, err = json.Marshal(params); err != nil {
		return nil, err
	}
	ctx := &Context{
		Name:          fmt.Sprintf("%s/contexts/%s", rw.Session, name),
		LifespanCount: lifespan,
		Parameters:    b,
	}
	return ctx, nil
}

// QueryResult is the dataset sent back by DialogFlow
type QueryResult struct {
	QueryText                 string          `json:"queryText,omitempty"`
	Action                    string          `json:"action,omitempty"`
	LanguageCode              string          `json:"languageCode,omitempty"`
	AllRequiredParamsPresent  bool            `json:"allRequiredParamsPresent,omitempty"`
	IntentDetectionConfidence float64         `json:"intentDetectionConfidence,omitempty"`
	Parameters                json.RawMessage `json:"parameters,omitempty"`
	OutputContexts            []*Context      `json:"outputContexts,omitempty"`
	Intent                    Intent          `json:"intent,omitempty"`
}

// Intent describes the matched intent
type Intent struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
