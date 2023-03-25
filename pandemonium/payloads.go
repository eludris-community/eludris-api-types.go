package pandemonium

import (
	"github.com/eludris-community/eludris-api-types.go/models"
)

type OpcodeType string

const (
	PingOp          OpcodeType = "PING"
	PongOp          OpcodeType = "PONG"
	MessageCreateOp OpcodeType = "MESSAGE_CREATE"
)

type Payload struct {
	// The opcode of the payload.
	Op OpcodeType `json:"op"`
}

type Ping struct {
	Payload
}

type Pong struct {
	Payload
}

type MessageCreate struct {
	Payload
	// The message object that triggered the event.
	D models.Message `json:"d"`
}
