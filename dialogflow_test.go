package dialogflow

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_GetParams(t *testing.T) {
	type out struct {
		In  string `json:"in"`
		Out string `json:"out"`
	}
	tests := []struct {
		name        string
		params      json.RawMessage
		expected    out
		expectError bool
	}{
		{"should unmarshal fine", []byte(`{"in": "in", "out": "out"}`), out{"in", "out"}, false},
		{"should be empty with other data", []byte(`{"hello": "world"}`), out{}, false},
		{"should err", []byte(``), out{}, true},
		{"should be empty", []byte(`{}`), out{}, false},
		{"should match", []byte(`{"in": "helloworld", "out": "helloworld"}`), out{"helloworld", "helloworld"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := &Request{QueryResult: QueryResult{Parameters: tt.params}}

			var output out
			if err := rw.GetParams(&output); (err != nil) != tt.expectError {
				t.Errorf("Request.GetParams() error = %v, wantErr %v", err, tt.expectError)
			}
			assert.Equal(t, output, tt.expected, "should match")
		})
	}
}

func TestRequest_GetContext(t *testing.T) {
	type out struct {
		In  string `json:"in"`
		Out string `json:"out"`
	}
	tests := []struct {
		name     string
		fields   Contexts
		ctx      string
		expected out
		wantErr  bool
	}{
		{
			"should find and unmarshal",
			Contexts{{"hello-ctx", 1, []byte(`{"in": "in", "out": "out"}`)}},
			"hello-ctx",
			out{"in", "out"},
			false,
		},
		{
			"should fail",
			Contexts{{"hello-ctx", 1, []byte(`{"in": "in", "out": "out"}`)}},
			"random-ctx",
			out{},
			true,
		},
		{
			"should work with multiple contexts",
			Contexts{
				{"random-ctx", 1, []byte(`{"in": "rand", "out": "rand"}`)},
				{"hello-ctx", 1, []byte(`{"in": "in", "out": "out"}`)},
			},
			"hello-ctx",
			out{"in", "out"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := &Request{QueryResult: QueryResult{OutputContexts: tt.fields}}

			var output out
			if err := rw.GetContext(tt.ctx, &output); (err != nil) != tt.wantErr {
				t.Errorf("Request.GetContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, output, tt.expected, "should match")
		})
	}
}

func TestRequest_NewContext(t *testing.T) {
	type out struct {
		In  string `json:"in"`
		Out string `json:"out"`
	}
	type fields struct {
		Session                     string
		ResponseID                  string
		QueryResult                 QueryResult
		OriginalDetectIntentRequest json.RawMessage
	}
	std := fields{Session: "session"}
	type args struct {
		name     string
		lifespan int
		params   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Context
		wantErr bool
	}{
		{
			"should generate properly",
			std,
			args{"hello-ctx", 3, out{"hello", "world"}},
			&Context{"session/contexts/hello-ctx", 3, []byte(`{"in": "hello", "out": "world"}`)},
			false,
		},
		{
			"should error",
			std,
			args{"hello-ctx", 3, make(chan int)},
			&Context{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := &Request{
				Session:                     tt.fields.Session,
				ResponseID:                  tt.fields.ResponseID,
				QueryResult:                 tt.fields.QueryResult,
				OriginalDetectIntentRequest: tt.fields.OriginalDetectIntentRequest,
			}
			got, err := rw.NewContext(tt.args.name, tt.args.lifespan, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.NewContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.Equal(t, got.Name, tt.want.Name)
			assert.Equal(t, got.LifespanCount, tt.want.LifespanCount)
		})
	}
}
