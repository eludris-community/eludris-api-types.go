package models

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
