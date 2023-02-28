package oprish

// Represents a message sent to or received from Eludris. The message contains its author and content.
type Message struct {
	// The author of the message, between 2-32 characters long.
	Author string `json:"author"`
	// The content of the message. At least one character; the upper bound is set per-instance inside Eludris.toml.
	// This upper bound can be requested as part of the instance's InstanceInfo.
	Content string `json:"content"`
}

// Represents info about an instance.
type InstanceInfo struct {
	// The name of the connected Eludris instance.
	InstanceName string `json:"instance_name"`
	// The description of the connected Eludris instance. If provided, this must be within 1 and 2048 characters long.
	Description string `json:"description"`
	// The version of the connected Eludris instance.
	Version string `json:"version"`
	// The maximum allowed message content length.
	MessageLimit string `json:"message_limit"`
	// 	The url to this instance's REST api.
	OprishUrl string `json:"oprish_url"`
	// The url to this instance's gateway.
	PandemoniumUrl string `json:"pandemonium_url"`
	// The url to this instance's CDN.
	EffisUrl string `json:"effis_url"`
	// The maximum number of bytes for an asset.
	FileSize int `json:"file_size"`
	// The maximum number of bytes for an attachment.
	AttachmentFileSize int `json:"attachment_file_size"`
}

// Rate limits in Eludris fully depend on the instance's configuration.
type RateLimitConf struct {
	// The number of seconds the client should wait before making new requests.
	ResetAfter int `json:"reset_after"`
	//The number of requests that can be made within the time frame denoted by `reset_after`.
	Limit int `json:"limit"`
}

//Represents all rate limits that apply to the connected Eludris instance.
// This includes individual rate limit information for Oprish (REST-api), Pandemonium (Gateway), and Effis (CDN).
type InstanceRateLimits struct {
	// The rate limits that apply to the connected Eludris instance's REST api.
	Oprish OprishRateLimits  `json:"oprish"`
	// The rate limits that apply to the connected Eludris instance's gateway.
	Pandemonium RateLimitConf `json:"pandemonium"`
	// The rate limits that apply to the connected Eludris instance's CDN.
	Effis EffisRateLimits `json:"effis"`
}

// Represents the rate limits for Oprish (REST-api).
// This denotes the rate limits on each individual route.
type OprishRateLimits struct {
	// The rate limit information for the `/` route.
	Info RateLimitConf `json:"info"`
	// The rate limit information for the `/messages` route.
	MessageCreate RateLimitConf `json:"message_create"`
	// The rate limit information for the `/rate_limits` route.
	RateLimits RateLimitConf `json:"rate_limits"`
}

// Represents a singular rate limit for Effis (CDN).
// Unlike normal rate limits, these also include a file size limit.
type EffisRateLimitConf struct {
	// The number of seconds the client should wait before making new requests.
	ResetAfter int `json:"reset_after"`
	// The number of requests that can be made within the time frame denoted by `reset_after`.
	Limit int `json:"limit"`
	// The maximum number of bytes that can be uploaded in the time frame denoted by `reset_after`.
	FileSizeLimit int `json:"file_size_limit"`
}

// Represents the rate limits for Effis.
// These rate limits denote how often files can be requested from the CDN, as well as maximum file size limits.
type EffisRateLimits struct {
	// The rate limit information for the handling of Assets.
	Assets EffisRateLimitConf `json:"assets"`
	// The rate limit information for the handling of Attachments.
	Attachments EffisRateLimitConf `json:"attachments"`
}
