package push

const (
	// Success sent
	Success Status = "80000000"
	// IllegalToken - Tokens identified by illegal_token are those failed to be sent.
	IllegalToken Status = "80100000"
	// NotCorrectToken - Some token parameters are incorrect.
	NotCorrectToken Status = "80100001"
	// SyncCountToken - The number of tokens must be 1 when a synchronization message is sent.
	SyncCountToken Status = "80100002"
	// IncorrectMessage - Incorrect message structure.
	IncorrectMessage Status = "80100003"
	// TTL - The message expiration time is earlier   than the current time.
	TTL Status = "80100004"
	// ColapseKey - The collapse_key message field is invalid.
	ColapseKey Status = "80100013"
	// SensitiveInformation - The message contains sensitive information.
	SensitiveInformation Status = "80100016"
	// OAuth - OAuth authentication error.
	OAuth Status = "80200001"
	// OAuthExpired - OAuth token expired.
	OAuthExpired Status = "80200003"
	// AppPermission - The current app does not have the permission to send push messages.
	AppPermission Status = "80300002"
	// InvalidTokens - All tokens are invalid.
	InvalidTokens Status = "80300007"
	// MessageSize - The message body size exceeds the default value.
	MessageSize Status = "80300008"
	// NumberTokens - The number of tokens in the message body exceeds the default value.
	NumberTokens Status = "80300010"
	// Priority - You are not authorized to send high-priority notification messages.
	Priority Status = "80300011"
	// Internal - System internal error.
	Internal Status = "81000001"
)

// Status response
type Status string
