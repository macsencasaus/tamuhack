package messages

import (
	"encoding/json"
)

type ClientMessageType string

const (
	ClientMessageTypeReady      = "ClientMessageReady"
	ClientMessageTypeClientQuit = "ClientMessageClientQuit"
	ClientMessageTypeSubmit     = "ClientMessageSubmit"
)

type ClientMessage interface {
	clientMessage()
}

type ClientMessageWrapper struct {
	Type ClientMessageType `json:"type"`
	Data json.RawMessage   `json:"data"`
}

type ReadyMessage struct {
    Type ClientMessageType `json:"type"`
}

func (m ReadyMessage) clientMessage() {}

type ClientQuitMessage struct {
	PlayerId int `json:"playerId"`
}

func (m ClientQuitMessage) clientMessage() {}

type SubmitMessage struct {
	PlayerId int    `json:"playerId"`
	Code     string `json:"code"`
}

func (m SubmitMessage) clientMessage() {}
