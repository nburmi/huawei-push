package push

import "encoding/json"

const (
	LOW    Importance = "LOW"
	NORMAL Importance = "NORMAL"
	HIGH   Importance = "HIGH"

	VisibilityUnspecified Visibilty = "VISIBILITY_UNSPECIFIED"
	PRIVATE               Visibilty = "PRIVATE"
	PUBLIC                Visibilty = "PUBLIC"
	SECRET                Visibilty = "SECRET"
)

type Importance string

type Visibilty string

type dataPush struct {
	M            *Message `json:"message"`
	ValidateOnly bool     `json:"validate_only,omitempty"`
}

// Message to be sent via HMS.
type Message struct {
	Data         json.RawMessage `json:"data,omitempty"`
	Notification *Notification   `json:"notification,omitempty"`
	Android      *AndroidConfig  `json:"android,omitempty"`
	Webpush      *WebpushConfig  `json:"webpush,omitempty"`
	APNS         *APNSConfig     `json:"apns,omitempty"`
	Tokens       []string        `json:"token,omitempty"`
	Topic        string          `json:"-"`
	Condition    string          `json:"condition,omitempty"`
}

// Notification is the basic notification template to use across all platforms.
type Notification struct {
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	ImageURL string `json:"image,omitempty"`
}

// AndroidConfig contains messaging options specific to the Android platform.
type AndroidConfig struct {
	CollapseKey   int                  `json:"collapse_key,omitempty"`
	Urgency       string               `json:"urgency,omitempty"`      // one of "normal" or "high"
	Category      string               `json:"category,omitempty"`     // Scenario where a high-priority data message is sent.
	TTL           string               `json:"ttl,omitempty"`          // Message cache time, in seconds.
	BiTag         string               `json:"bi_tag,omitempty"`       // Tag of a message in a batch delivery task.
	FastAppTarget int                  `json:"fast_app_target"`        // State of a mini program when a quick app sends a data message. 1 - dev, 2 - prod(default).
	Data          json.RawMessage      `json:"data,omitempty"`         // if specified, overrides the Data field on Message type
	Notification  *AndroidNotification `json:"notification,omitempty"` // Android notification message structure.
}

// AndroidNotification is a notification to send to Android devices.
type AndroidNotification struct {
	ForegroundShow    bool                   `json:"foreground_show,omitempty"`
	DefaultSound      bool                   `json:"default_sound,omitempty"`
	AutoCancel        bool                   `json:"auto_cancel,omitempty"` //Indicates whether an Android notification message is not still displayed in the notification bar after a user taps the message.
	UseDefaultVibrate bool                   `json:"use_default_vibrate,omitempty"`
	UseDefaultLight   bool                   `json:"use_default_light,omitempty"`
	Title             string                 `json:"title,omitempty"` // if specified, overrides the Title field of the Notification type
	Body              string                 `json:"body,omitempty"`  // if specified, overrides the Body field of the Notification type
	Icon              string                 `json:"icon,omitempty"`
	Color             string                 `json:"color,omitempty"` // notification color in #RRGGBB format
	Sound             string                 `json:"sound,omitempty"`
	Tag               string                 `json:"tag,omitempty"`
	ClickAction       ClickAction            `json:"click_action,omitempty"`
	BodyLocKey        string                 `json:"body_loc_key,omitempty"`
	BodyLocArgs       []string               `json:"body_loc_args,omitempty"`
	TitleLocKey       string                 `json:"title_loc_key,omitempty"`
	TitleLocArgs      []string               `json:"title_loc_args,omitempty"`
	MultiLangKey      map[string]interface{} `json:"multi_lang_key,omitempty"`
	ChannelID         string                 `json:"channel_id,omitempty"`
	NotifySummary     string                 `json:"notify_summary"`
	ImageURL          string                 `json:"image,omitempty"`
	Style             int                    `json:"style,omitempty"`
	BigTitle          string                 `json:"big_title,omitempty"`
	BigBody           string                 `json:"big_body,omitempty"`
	AutoClear         int                    `json:"auto_clear,omitempty"`
	NotifyID          int                    `json:"notify_id,omitempty"` // Unique notification ID of a message.
	Group             string                 `json:"group,omitempty"`
	Badge             *BadgeNotification     `json:"badge,omitempty"`
	Ticker            string                 `json:"ticker,omitempty"`
	When              string                 `json:"when,omitempty"`
	Importance        Importance             `json:"importance,omitempty"`
	VibrateConfig     []string               `json:"vibrate_config,omitempty"`
	Visibility        Visibilty              `json:"visibility,omitempty"`
	LightSettings     *LightSettings         `json:"light_settings,omitempty"`
}

type ClickAction struct {
	Type         int    `json:"type,omitempty"`
	Intent       string `json:"intent,omitempty"`
	URL          string `json:"url,omitempty"`
	RichResource string `json:"rich_resource,omitempty"`
	Action       string `json:"action,omitempty"`
}

type LightSettings struct {
	Color            Color  `json:"color,omitempty"`
	LightOnDuration  string `json:"light_on_duration,omitempty"`
	LightOffDuration string `json:"light_off_duration,omitempty"`
}

type Color struct {
	Alpha float64 `json:"alpha,omitempty"`
	Red   float64 `json:"red,omitempty"`
	Green float64 `json:"green,omitempty"`
	Blue  float64 `json:"blue,omitempty"`
}

type BadgeNotification struct {
	AddNum int    `json:"add_num,omitempty"`
	Class  string `json:"class,omitempty"`
	SetNum int    `json:"set_num,omitempty"`
}

// AndroidFCMOptions contains additional options for features provided by the FCM Android SDK.
type AndroidFCMOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
}

// WebpushConfig contains messaging options specific to the WebPush protocol.
//
// See https://tools.ietf.org/html/rfc8030#section-5 for additional details, and supported
// headers.
type WebpushConfig struct {
	Headers      Headers              `json:"headers,omitempty"`
	Notification *WebpushNotification `json:"notification,omitempty"`
	HmsOptions   *WebpushHmsOptions   `json:"hms_options,omitempty"`
}

type Headers struct {
	TTL     string `json:"ttl,omitempty"`
	Topic   string `json:"topic,omitempty"`
	Urgency string `json:"urgency,omitempty"` //very-low, low, normal, or high.
}

type WebpushHmsOptions struct {
	Link string `json:"link,omitempty"`
}

// WebpushNotificationAction represents an action that can be performed upon receiving a WebPush notification.
type WebpushNotificationAction struct {
	Action string `json:"action,omitempty"`
	Title  string `json:"title,omitempty"`
	Icon   string `json:"icon,omitempty"`
}

// WebpushNotification is a notification to send via WebPush protocol.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/notification/Notification for additional
// details.
type WebpushNotification struct {
	Title              string                       `json:"title,omitempty"` // if specified, overrides the Title field of the Notification type
	Body               string                       `json:"body,omitempty"`  // if specified, overrides the Body field of the Notification type
	Icon               string                       `json:"icon,omitempty"`
	Image              string                       `json:"image,omitempty"`
	Language           string                       `json:"lang,omitempty"`
	Tag                string                       `json:"tag,omitempty"`
	Badge              string                       `json:"badge,omitempty"`
	Direction          string                       `json:"dir,omitempty"` // one of 'ltr' or 'rtl'
	Vibrate            []int                        `json:"vibrate,omitempty"`
	Renotify           bool                         `json:"renotify,omitempty"`
	RequireInteraction bool                         `json:"requireInteraction,omitempty"`
	Silent             bool                         `json:"silent,omitempty"`
	TimestampMillis    *int64                       `json:"timestamp,omitempty"`
	Actions            []*WebpushNotificationAction `json:"actions,omitempty"`
}

// WebpushFcmOptions contains additional options for features provided by the FCM web SDK.
type WebpushFcmOptions struct {
	Link string `json:"link,omitempty"`
}

// APNSConfig contains messaging options specific to the Apple Push Notification Service (APNS).
//
// See https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html
// for more details on supported headers and payload keys.
type APNSConfig struct {
	Headers    map[string]string `json:"headers,omitempty"`
	Payload    map[string]string `json:"payload,omitempty"`
	HMSOptions *APNSHMSOptions   `json:"hms_options,omitempty"`
}

// APNSHMSOptions contains additional options for features provided by the FCM Aps SDK.
type APNSHMSOptions struct {
	TargetUserType int `json:"target_user_type,omitempty"`
}
