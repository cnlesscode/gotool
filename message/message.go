package message

import "encoding/json"

type Response struct {
	ErrorCode int
	Data      any
}

func (m Response) ToString() string {
	message, _ := json.Marshal(m)
	return string(message)
}

func (m Response) ToByte() []byte {
	message, _ := json.Marshal(m)
	return message
}
