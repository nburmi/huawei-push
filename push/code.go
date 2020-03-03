package push

const (
	// Success sent
	Success Code = "80000000"
	// IllegalToken - Tokens identified by illegal_token are those failed to be sent.
	IllegalToken Code = "80100000"
	// NotCorrectToken - Some token parameters are incorrect.
	NotCorrectToken Code = "80100001"
	// SyncCountToken - The number of tokens must be 1 when a synchronization message is sent.
	SyncCountToken Code = "80100002"
	// IncorrectMessage - Incorrect message structure.
	IncorrectMessage Code = "80100003"
	// TTL - The message expiration time is earlier   than the current time.
	TTL Code = "80100004"
	// ColapseKey - The collapse_key message field is invalid.
	ColapseKey Code = "80100013"
	// SensitiveInformation - The message contains sensitive information.
	SensitiveInformation Code = "80100016"
	// OAuth - OAuth authentication error.
	OAuth Code = "80200001"
	// OAuthExpired - OAuth token expired.
	OAuthExpired Code = "80200003"
	// AppPermission - The current app does not have the permission to send push messages.
	AppPermission Code = "80300002"
	// InvalidTokens - All tokens are invalid.
	InvalidTokens Code = "80300007"
	// MessageSize - The message body size exceeds the default value.
	MessageSize Code = "80300008"
	// NumberTokens - The number of tokens in the message body exceeds the default value.
	NumberTokens Code = "80300010"
	// Priority - You are not authorized to send high-priority notification messages.
	Priority Code = "80300011"
	// Internal - System internal error.
	Internal Code = "81000001"
)

// Code response
type Code string
