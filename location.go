package dialogflow

import "encoding/json"

// Location is a location object sent back by DialogFlow
type Location struct {
	Simple               string
	AdminArea            string          `json:"admin-area,omitempty"`
	AdminAreaOriginal    string          `json:"admin-area.original,omitempty"`
	AdminAreaObject      json.RawMessage `json:"admin-area.object,omitempty"`
	SubAdminArea         string          `json:"subadmin-area,omitempty"`
	SubAdminAreaOriginal string          `json:"subadmin-area.original,omitempty"`
}

// UnmarshalJSON implements the Unmarshaler interface for JSON parsing
// This function will try to unmarshal the incoming data to either a full
// Location object, or a simple string.
func (l *Location) UnmarshalJSON(b []byte) error {
	var err error
	// Aliasing to avoid recursion in the unmarshalling process
	type location Location

	// Possible types to unmarshal to
	ll, s := location{}, ""

	if err = json.Unmarshal(b, &ll); err == nil {
		*l = Location(ll)
		return err
	}
	if err = json.Unmarshal(b, &s); err == nil {
		l.Simple = s
	}
	return err
}
