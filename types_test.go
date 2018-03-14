package dialogflow

import (
	"reflect"
	"testing"
)

func TestBasicCard_GetKey(t *testing.T) {
	want := "basicCard"
	b := BasicCard{}
	if got := b.GetKey(); got != want {
		t.Errorf("BasicCard.GetKey() = %v, want %v", got, want)
	}
}

func TestCard_GetKey(t *testing.T) {
	want := "card"
	c := Card{}
	if got := c.GetKey(); got != want {
		t.Errorf("Card.GetKey() = %v, want %v", got, want)
	}
}

func TestSimpleResponsesWrapper_GetKey(t *testing.T) {
	want := "simpleResponses"
	s := SimpleResponsesWrapper{}
	if got := s.GetKey(); got != want {
		t.Errorf("SimpleResponsesWrapper.GetKey() = %v, want %v", got, want)
	}
}

func TestText_GetKey(t *testing.T) {
	want := "text"
	txt := Text{}
	if got := txt.GetKey(); got != want {
		t.Errorf("Text.GetKey() = %v, want %v", got, want)
	}
}

func TestQuickReplies_GetKey(t *testing.T) {
	want := "quickReplies"
	q := QuickReplies{}
	if got := q.GetKey(); got != want {
		t.Errorf("QuickReplies.GetKey() = %v, want %v", got, want)
	}
}

func TestImage_GetKey(t *testing.T) {
	want := "image"
	i := Image{}
	if got := i.GetKey(); got != want {
		t.Errorf("Image.GetKey() = %v, want %v", got, want)
	}
}

func TestPayloadWrapper_GetKey(t *testing.T) {
	want := "payload"
	p := PayloadWrapper{}
	if got := p.GetKey(); got != want {
		t.Errorf("Payload.GetKey() = %v, want %v", got, want)
	}
}

func TestSuggestions_GetKey(t *testing.T) {
	want := "suggestions"
	s := Suggestions{}
	if got := s.GetKey(); got != want {
		t.Errorf("Suggestions.GetKey() = %v, want %v", got, want)
	}
}

func TestLinkOutSuggestion_GetKey(t *testing.T) {
	want := "linkOutSuggestion"
	l := LinkOutSuggestion{}
	if got := l.GetKey(); got != want {
		t.Errorf("LinkOutSuggestion.GetKey() = %v, want %v", got, want)
	}
}

func TestListSelect_GetKey(t *testing.T) {
	want := "listSelect"
	l := ListSelect{}
	if got := l.GetKey(); got != want {
		t.Errorf("ListSelect.GetKey() = %v, want %v", got, want)
	}
}

func TestCarouselSelect_GetKey(t *testing.T) {
	want := "carouselSelect"
	c := CarouselSelect{}
	if got := c.GetKey(); got != want {
		t.Errorf("CarouselSelect.GetKey() = %v, want %v", got, want)
	}
}

func TestSingleSimpleResponse(t *testing.T) {
	type args struct {
		display string
		speech  string
	}
	tests := []struct {
		name string
		args args
		want SimpleResponsesWrapper
	}{
		{
			"same text and speech",
			args{"hello world", "hello world"},
			SimpleResponsesWrapper{
				SimpleResponses: []SimpleResponse{
					{DisplayText: "hello world", TextToSpeech: "hello world"},
				},
			},
		},
		{
			"different text and speech",
			args{"hello world", "hi"},
			SimpleResponsesWrapper{
				SimpleResponses: []SimpleResponse{
					{DisplayText: "hello world", TextToSpeech: "hi"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleSimpleResponse(tt.args.display, tt.args.speech); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleSimpleResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayloadWrapper_MarshalJSON(t *testing.T) {
	type fields struct {
		Payload interface{}
	}
	type minimal struct {
		Hello string `json:"hello"`
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"should fail", fields{make(chan int)}, []byte(""), true},
		{"should marshal", fields{minimal{"hi"}}, []byte(`{"hello":"hi"}`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PayloadWrapper{
				Payload: tt.fields.Payload,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("PayloadWrapper.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayloadWrapper.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
