package goclient

import (
	"fmt"
	"net"
	"strings"

	CData "mkGOchat.commont/datas"
)

const (
	SERVER_HOST = "192.168.0.107"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var thisClie CData.MkClient

func StartClient() {
	thisConn, conErr := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if conErr != nil {
		panic("erro... " + conErr.Error())
	}

	thisClie = CData.MkClient{
		ClientConnection: thisConn,
		ClientDetais:     CData.MkClientDetails{Nick: fmt.Sprint(10)},
	}

	ConsoleInput()
	thisClie.Listen()
}

func ConsoleInput() {
	for {
		fmt.Println("Waitting for message with user:message\n-")
		var toSend string
		l, _ := fmt.Scanln(&toSend)
		if l > 1024 {
			fmt.Println("Demasiado largo el mensaje")
		}

		msgD := strings.Split(toSend, ":")

		msg := CData.MkMessage{
			MsgType:  CData.ConnSend_MSG,
			MsgValue: msgD[1],
			ClieFrom: thisClie.ClientDetais.Nick,
			ClieTo:   msgD[0],
		}

		isOk, errMsg := thisClie.Send(msg)
		if !isOk {
			fmt.Println("Error reading:", errMsg)
		}
	}
}

/*
func StartClient() {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	go listengForServer(connection)
	///send some data
	for {
		var toSend string
		l, err := fmt.Scanln(&toSend)
		if l > 1024 {
			fmt.Println("Demasiado largo el mensaje")
		}
		_, err = connection.Write([]byte(toSend))
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
	}

}*/
