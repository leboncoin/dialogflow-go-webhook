package dialogflow

// Platform is a simple type intended to be used with responses
type Platform string

// Platform constants, used in the webhook responses
const (
	Unspecified     Platform = "PLATFORM_UNSPECIFIED"
	Facebook        Platform = "FACEBOOK"
	Slack           Platform = "SLACK"
	Telegram        Platform = "TELEGRAM"
	Kik             Platform = "KIK"
	Skype           Platform = "SKYPE"
	Line            Platform = "LINE"
	Viber           Platform = "VIBER"
	ActionsOnGoogle Platform = "ACTIONS_ON_GOOGLE"
)
