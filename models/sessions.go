package models

// Represents a session payload on Oprish.
type Session struct {
	// The id of the session.
	Id string `json:"id"`
	// The user ID of the session.
	UserId string `json:"user_id"`
	// The platform the session is on.
	Platform string `json:"platform"`
	// The client the session was created by.
	Client string `json:"client"`
	// The session's creation IP address.
	Ip string `json:"ip"`
}

// Represents a created session payload on Oprish.
type SessionCreate struct {
	// The session user's identifier.
	Identifier string `json:"identifier"`
	// The session user's password.
	Password string `json:"password"`
	// The session's platform.
	Platform string `json:"platform"`
	// The client the session was created by.
	Client string `json:"client"`
}

// Represents a response to SessionCreate
type SessionCreated struct {
	// The session's token.
	Token string `json:"token"`
	// The session object that was created.
	Session Session `json:"session"`
}
