package registerEvent

import "encoding/json"

//	ResponseValue contains the root of the value field when a register request was made
type ResponseValue struct {
	Backend Backend `json:"backend"`
	ID      string  `json:"id"`
}

//	Backend contains information of the parameters for pull or sse mode
type Backend struct {
	SSE  SSEInfo  `json:"sse"`
	Pull PullInfo `json:"pull"`
}

//	PullInfo contains information of how much entries can be pulled at once
type PullInfo struct {
	MaxEntries int    `json:"maxEntries"`
	Store      string `json:"store"`
}

//	SSEInfo contains information of contenttypes and encoding of sse events
type SSEInfo struct {
	ContentType string `json:"backChannel.contentType"`
	Encoding    string `json:"backChannel.encoding"`
}

func DecodeResponseValue(value json.RawMessage) (reg ResponseValue, err error) {
	err = json.Unmarshal(value, reg)
	return
}
