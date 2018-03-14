package dialogflow

import "encoding/json"

// BasicCard is a simple card. Simply adds an extra Image field
type BasicCard struct {
	Title         string       `json:"title,omitempty"`         // Optional. The title of the card.
	Subtitle      string       `json:"subtitle,omitempty"`      // Optional. The subtitle of the card.
	FormattedText string       `json:"formattedText,omitempty"` // Required, unless image is present. The body text of the card.
	Image         *Image       `json:"image,omitempty"`         // Optional. The image for the card.
	Buttons       []CardButton `json:"buttons,omitempty"`       // Optional. The collection of card buttons.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the BasicCard type
func (bc BasicCard) GetKey() string {
	return "basicCard"
}

// Card is a simple card. Different type than the BasicCard, the buttons change.
// This card has buttons with postback field, whereas BasicCard has cards with
// OpenURI action
type Card struct {
	Title    string   `json:"title,omitempty"`    // Optional. The title of the card.
	Subtitle string   `json:"subtitle,omitempty"` // Optional. The subtitle of the card.
	Buttons  []Button `json:"buttons,omitempty"`  // Optional. The collection of card buttons.
	Image
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the Card type
func (c Card) GetKey() string {
	return "card"
}

// SimpleResponsesWrapper wraps SimpleResponses
type SimpleResponsesWrapper struct {
	SimpleResponses []SimpleResponse `json:"simpleResponses,omitempty"` // Required. The list of simple responses.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the SimpleResponsesWrapper type
func (s SimpleResponsesWrapper) GetKey() string {
	return "simpleResponses"
}

// SingleSimpleResponse is a wrapper to create one simple response to display
// and play to the user
func SingleSimpleResponse(display, speech string) SimpleResponsesWrapper {
	return SimpleResponsesWrapper{
		SimpleResponses: []SimpleResponse{
			{TextToSpeech: speech, DisplayText: display},
		},
	}
}

// SimpleResponse is a simple response sent back to dialogflow.
// Composed of two types, TextToSpeech will be converted to speech and
// DisplayText will be displayed if the surface allows to display stuff
type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"` // One of textToSpeech or ssml must be provided. The plain text of the speech output. Mutually exclusive with ssml.
	DisplayText  string `json:"displayText,omitempty"`  // Optional. The text to display.
	SSML         string `json:"ssml,omitempty"`         // One of textToSpeech or ssml must be provided. Structured spoken response to the user in the SSML format. Mutually exclusive with textToSpeech.
}

// Text is a simple wrapper around a list of string
type Text struct {
	Text []string `json:"text,omitempty"` // Optional. The collection of the agent's responses.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the Text type
func (t Text) GetKey() string {
	return "text"
}

// QuickReplies is a structured response sent back to dialogflow
type QuickReplies struct {
	Title   string   `json:"title,omitempty"`        // Optional. The title of the collection of quick replies.
	Replies []string `json:"quickReplies,omitempty"` // Optional. The collection of quick replies.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the QuickReplies type
func (qr QuickReplies) GetKey() string {
	return "quickReplies"
}

// CardButton is a he button object that appears at the bottom of a card
type CardButton struct {
	Title         string         `json:"title,omitempty"`         // Optional. The text to show on the button.
	OpenURIAction *OpenURIAction `json:"openUriAction,omitempty"` // Required. Action to take when a user taps on the button.
}

// Button contains information about a button
type Button struct {
	Text     string `json:"text,omitempty"`     // Optional. The text to show on the button.
	PostBack string `json:"postback,omitempty"` // Optional. The text to send back to the Dialogflow API or a URI to open.
}

// OpenURIAction simply defines the URI associated with a button
type OpenURIAction struct {
	URI string `json:"uri,omitempty"` // Required. The HTTP or HTTPS scheme URI.
}

// Image is a simple type of message sent back to dialogflow
type Image struct {
	ImageURI string `json:"imageUri,omitempty"` // Optional. The public URI to an image file.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the Image type
func (i Image) GetKey() string {
	return "image"
}

// PayloadWrapper acts as a wrapper for the payload type
type PayloadWrapper struct {
	Payload interface{}
}

// MarshalJSON implements the Marshaller interface and will return the
// marshalled payload, ignoring the initial level
func (p PayloadWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Payload)
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated to the Payload type
func (p PayloadWrapper) GetKey() string {
	return "payload"
}

// Suggestions is a rich message that suggests the user quick replies
type Suggestions struct {
	Suggestions []Suggestion `json:"suggestions,omitempty"` // Required. The list of suggested replies.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated to the Suggestions type
func (s Suggestions) GetKey() string {
	return "suggestions"
}

// Suggestion is a simple suggestion
type Suggestion struct {
	Title string `json:"title,omitempty"` // Required. The text shown the in the suggestion chip.
}

// LinkOutSuggestion can be used to suggest the user to click on a link that
// will get him out of the app
type LinkOutSuggestion struct {
	DestinationName string `json:"suggestionName,omitempty"` // Required. The name of the app or site this chip is linking to.
	URI             string `json:"uri,omitempty"`            // Required. The URI of the app or site to open when the user taps the suggestion chip.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the LinkOutSuggestion type
func (l LinkOutSuggestion) GetKey() string {
	return "linkOutSuggestion"
}

// ListSelect can be used to show a list to the user
type ListSelect struct {
	Title string `json:"title,omitempty"` // Required. The overall title of the list.
	Items []Item `json:"items,omitempty"` // Required. List items.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the ListSelect type
func (l ListSelect) GetKey() string {
	return "listSelect"
}

// Item is a single item present in a list
type Item struct {
	Info        SelectItemInfo `json:"info,omitempty"`        // Required. Additional information about this option
	Title       string         `json:"title,omitempty"`       // Required. The title of the list item
	Description string         `json:"description,omitempty"` // Optional. The main text describing the item
	Image       *Image         `json:"image,omitempty"`       // Optional. The image to display.
}

// SelectItemInfo is a struct holding data about a specific item
type SelectItemInfo struct {
	Key      string   `json:"key,omitempty"`      //  Required. A unique key that will be sent back to the agent if this response is given.
	Synonyms []string `json:"synonyms,omitempty"` //  Optional. A list of synonyms that can also be used to trigger this item in dialog.
}

// CarouselSelect shows a carousel to the user
type CarouselSelect struct {
	Items []Item `json:"items,omitempty"` // Required. Carousel items.
}

// GetKey implements the RichMessage interface and returns the JSON key
// associated with the CarouselSelect type
func (c CarouselSelect) GetKey() string {
	return "carouselSelect"
}
