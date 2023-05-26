package pandemonium

import (
	"encoding/json"
)

type OpcodeType string

type Payload struct {
	// The opcode of the payload.
	Op OpcodeType `json:"op"`
	// The data of the payload.
	D json.RawMessage `json:"d,omitempty"`
}
