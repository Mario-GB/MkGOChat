package gocommon_datas

import (
	"encoding/json"
	"log"
	"net"

	Utils "mkGOchat.commont/utils"
)

// {STRUCTURE} Client structure to save connections an clients data
type MkClient struct {
	ClientConnection net.Conn
	ClientDetais     MkClientDetails
	OnMsgRecive      func(MkMessage)
}

// Get client connection
func (clie MkClient) GetConn() net.Conn {
	if clie.ClientConnection == nil {
		log.Fatalln("Connection is not defined")
	}
	return clie.ClientConnection
}

// Send a message like json
func (clie MkClient) Send(msg MkMessage) (bool, string) {
	isOk := false
	errStr := ""

	jsonMsg, err := json.Marshal(msg)
	jsonMsgStr := string(jsonMsg)
	if err != nil {
		errStr = err.Error()
	}

	errStr = Utils.ConnectionSend(clie.GetConn(), jsonMsgStr)

	isOk = errStr != "0"

	return isOk, errStr
}

// Recive a message and work about it
func (clie MkClient) Recive(msg string) {
	msgMk := MkMessage{}
	convertErr := json.Unmarshal([]byte(msg), &msgMk)

	if convertErr != nil {
		log.Fatalln(convertErr)
	}

	switch msgMk.MsgType {
	case ConnSend_MSG:
		clie.OnMsgRecive(msgMk)
		break
	case ConnSend_PING:
		Utils.ConnectionSend(clie.GetConn(), "0")
		break
	}
}

// Start listent to message and redirect to recive
func (clie MkClient) Listen() {
	go Utils.ConnectionListen(clie.GetConn(), clie.Recive)
}

// {STRUCTURE} ClientDetails structure to save diferent datas about the users
type MkClientDetails struct {
	NameFirst   string
	NameLast    string
	Nick        string
	MailAddress string
	ClientId    string
}

// {STRUCTURE} Message structure to easy communitcation
type MkMessage struct {
	MsgType  int    `json:"MsgType"`
	MsgValue string `json:"MsgValue"`
	ClieFrom string `json:"ClieFrom"`
	ClieTo   string `json:"ClieTo"`
}
