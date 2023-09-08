package models

// Represents a message received from Eludris. The message contains its author and content.
type Message struct {
	// The author of the message.
	Author User `json:"author"`
	// The content of the message. At least one character; the upper bound is set per-instance inside Eludris.toml.
	// This upper bound can be requested as part of the instance's InstanceInfo.
	Content  string           `json:"content"`
	Disguise *MessageDisguise `json:"_disguise"`
}

// Represents the message creation payload on Oprish.
type MessageCreate struct {
	// The message's content.
	Content  string           `json:"content"`
	Disguise *MessageDisguise `json:"_disguise"`
}

// A temporary way to mask the message’s author’s name and avatar.
type MessageDisguise struct {
	// The name of the message's disguise
	Name *string `json:"name"`
	// The URL of the message's disguise
	Avatar *string `json:"avatar"`
}
