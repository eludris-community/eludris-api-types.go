package models

// Rate limits in Eludris fully depend on the instance's configuration.
type RateLimitConf struct {
	// The number of seconds the client should wait before making new requests.
	ResetAfter int `json:"reset_after"`
	//The number of requests that can be made within the time frame denoted by `reset_after`.
	Limit int `json:"limit"`
}

// Represents all rate limits that apply to the connected Eludris instance.
// This includes individual rate limit information for Oprish (REST-api), Pandemonium (Gateway), and Effis (CDN).
type InstanceRateLimits struct {
	// The rate limits that apply to the connected Eludris instance's REST api.
	Oprish OprishRateLimits `json:"oprish"`
	// The rate limits that apply to the connected Eludris instance's gateway.
	Pandemonium RateLimitConf `json:"pandemonium"`
	// The rate limits that apply to the connected Eludris instance's CDN.
	Effis EffisRateLimits `json:"effis"`
}

// Represents the rate limits for Oprish (REST-api).
// This denotes the rate limits on each individual route.
type OprishRateLimits struct {
	// The rate limit information for the `/` route.
	Info RateLimitConf `json:"get_instance_info"`
	// The rate limit information for the `/messages` route.
	MessageCreate RateLimitConf `json:"create_message"`
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
	// The rate limit information for the handling of Files.
	FetchFile EffisRateLimitConf `json:"fetch_file"`
}
