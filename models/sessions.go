package models

// Represents a session payload on Oprish.
type Session struct {
	// The session’s ID.
	Id string `json:"id"`
	// The session user’s ID.
	UserId string `json:"user_id"`
	// The session’s platform (linux, windows, mac, etc.)
	Platform string `json:"platform"`
	// The client the session was created by.
	Client string `json:"client"`
	// The session's creation IP address.
	Ip string `json:"ip"`
}

// Represents a created session payload on Oprish.
type SessionCreate struct {
	// The session user's identifier.
	// This can be either their email or username.
	Identifier string `json:"identifier"`
	// The session user's password.
	Password string `json:"password"`
	// The session’s platform (linux, windows, mac, etc.)
	Platform string `json:"platform"`
	// The client the session was created by.
	Client string `json:"client"`
}

// Represents a response to SessionCreate
type SessionCreated struct {
	// The session's token.
	// This can be used by the user to properly interface with the API.
	Token string `json:"token"`
	// The session object that was created.
	Session Session `json:"session"`
}
