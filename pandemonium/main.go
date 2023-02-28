package pandemonium

import (
	"github.com/toolifelesstocode/eludris-api-types.go/oprish"
)


type Ping struct {
	// Always `"PING"`
	Op string `json:"op"`
}

type Pong struct {
	// Always `"PONG"`
	Op string `json:"op"`
}

type MessageCreate struct {
	// Always `"MESSAGE_CREATE"`
	Op string `json:"op"`
	// The message object that triggered the event.
	D oprish.Message `json:"d"`
}
