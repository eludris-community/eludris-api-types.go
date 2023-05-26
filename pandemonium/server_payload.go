package pandemonium

import (
	"github.com/eludris-community/eludris-api-types.go/models"
)

const (
	PongOp          OpcodeType = "PONG"
	RatelimitOp     OpcodeType = "RATE_LIMIT"
	HelloOp         OpcodeType = "HELLO"
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

type MessageCreate models.Message
