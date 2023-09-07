package pandemonium

import (
	"github.com/eludris-community/eludris-api-types.go/v2/models"
)

const (
	PongOp          OpcodeType = "PONG"
	RatelimitOp     OpcodeType = "RATE_LIMIT"
	HelloOp         OpcodeType = "HELLO"
	AuthenticatedOp OpcodeType = "AUTHENTICATED"
	UserUpdateOP    OpcodeType = "USER_UPDATE"
	PresenceUpdateOP OpcodeType = "PRESENCE_UPDATE"
	MessageCreateOp OpcodeType = "MESSAGE_CREATE"
)

type Ratelimit struct {
	// The number of milliseconds to wait before making new requests.
	Wait int `json:"wait"`
}

type Hello struct {
	// The heartbeat interval in milliseconds that the client should heartbeat to.
	HeartbeatInterval int `json:"heartbeat_interval"`
	// The information about the connected Eludris instance.
	// Rate limits are not included in this payload.
	InstanceInfo models.InstanceInfo `json:"instance_info"`
	// The pandemonium rate limit information for the connected Eludris instance.
	RateLimit models.RateLimitConf `json:"rate_limit"`
}

type Authenticated struct {
	User models.User `json:"user"`
	// The currently online users who are relavent to the connector.
	Users []models.User `json:"users"`
}

type UserCreate models.UserCreate

type PresenceUpdate struct {
	UserId int `json:"user_id"`
	Status models.Status `json:"status"`
}

type MessageCreate models.Message
