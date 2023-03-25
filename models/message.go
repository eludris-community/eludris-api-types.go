package models

// Represents a message sent to or received from Eludris. The message contains its author and content.
type Message struct {
	// The author of the message, between 2-32 characters long.
	Author string `json:"author"`
	// The content of the message. At least one character; the upper bound is set per-instance inside Eludris.toml.
	// This upper bound can be requested as part of the instance's InstanceInfo.
	Content string `json:"content"`
}
