package dialogflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		in      []byte
		want    *Location
		wantErr bool
	}{
		{"should unmarshal to location", []byte(`{"subadmin-area": "Paris"}`), &Location{Department: "Paris"}, false},
		{"should unmarshal to simple string", []byte(`"Paris"`), &Location{Simple: "Paris"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Location{}
			if err := l.UnmarshalJSON(tt.in); (err != nil) != tt.wantErr {
				t.Errorf("Location.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want, l)
		})
	}
}
