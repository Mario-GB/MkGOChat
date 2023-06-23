package goserver

import (
	"fmt"
	"net"

	Cdata "mkGOchat.commont/datas"
	Cutils "mkGOchat.commont/utils"
)

var AllClients Cutils.array

func StartSocket() {
	fmt.Println("Server Running...")
	// start server
	server, err := net.Listen(CONFIG_SERVER_TYPE, CONFIG_SERVER_HOST+":"+CONFIG_SERVER_PORT)
	if err != nil {
		fmt.Println("\nError listening:", err.Error(), "\n_______________________")
	}

	// Close server when Start method end
	defer server.Close()

	// Listening for clients
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("\nError accepting: ", err.Error(), "\n_______________________")
		}

		newClie := Cdata.MkClient{
			ClientConnection: connection,
			ClientDetais:     Cdata.MkClientDetails{Nick: fmt.Sprint(len(AllClients))},
			OnMsgRecive:      clientMsgRecived,
		}
		AllClients = append(AllClients, newClie)
		newClie.Listen()
	}
}

func clientMsgRecived(msg Cdata.MkMessage) {
	switch msg.MsgType {
	case Cdata.ConnSend_MSG:
		connectionMsg(msg)
		break
	case Cdata.ConnSend_INFO:
		connectionInfo(msg)
		break
	}
}

func connectionMsg(msg Cdata.MkMessage) {
	for _, eVal := range AllClients {
		if eVal.ClientDetais.Nick != msg.ClieTo {
			eVal.Send(msg)
		}
	}
}

func connectionInfo(msg Cdata.MkMessage) {
	for ()
}
