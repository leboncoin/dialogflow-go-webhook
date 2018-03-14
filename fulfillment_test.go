package dialogflow

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// wrong type that fails on marshal and implements the RichMessage interface
type wrong struct{}

func (wrong) GetKey() string {
	return "wrong"
}
func (wrong) MarshalJSON() ([]byte, error) {
	return []byte(""), errors.New("That's wrong")
}

func TestForGoogle(t *testing.T) {
	type args struct {
		r RichMessage
	}
	tests := []struct {
		name string
		args args
		want Message
	}{
		{"should generate rich message for google", args{PayloadWrapper{"hello"}}, Message{Platform: ActionsOnGoogle, RichMessage: PayloadWrapper{"hello"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForGoogle(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForGoogle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleForGoogle() {
	output := "Hello World !"
	fulfillment := Fulfillment{
		FulfillmentMessages: Messages{
			ForGoogle(SingleSimpleResponse(output, output)),
		},
	}
	fmt.Println(fulfillment)
}

func TestMessage_MarshalJSON(t *testing.T) {
	type fields struct {
		Platform    Platform
		RichMessage RichMessage
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"should marshal", fields{}, []byte(`{}`), false},
		{"should marshal with platform", fields{Platform: ActionsOnGoogle}, []byte(`{"platform": "ACTIONS_ON_GOOGLE"}`), false},
		{
			"should marshal with platform and message",
			fields{Platform: ActionsOnGoogle, RichMessage: SingleSimpleResponse("hi", "hi")},
			[]byte(`{"platform": "ACTIONS_ON_GOOGLE", "simpleResponses": {"simpleResponses":[{"textToSpeech":"hi","displayText":"hi"}]}}`),
			false,
		},
		{
			"should fail because of message",
			fields{RichMessage: wrong{}},
			[]byte(""),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Platform:    tt.fields.Platform,
				RichMessage: tt.fields.RichMessage,
			}
			got, err := m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if err := JSONEqual(got, tt.want); err != nil {
				t.Errorf("Message.MarshalJSON() error =%v", err)
			}
		})
	}
}
