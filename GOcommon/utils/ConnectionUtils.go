package gocomm_utils

import (
	"log"
	"net"
)

// Sent a message and recive a feedback
func ConnectionSend(conn net.Conn, toSend string) string {

	log.Println("Trying connection send message.\nConnection: " + conn.LocalAddr().String() + "\nMessage: " + toSend)

	_, errWriting := conn.Write([]byte(toSend))
	if errWriting != nil {
		log.Fatalln("Error sending message\nConnection: " + conn.LocalAddr().String() + "\nError: " + errWriting.Error())
	}

	responsedMsg := ConnectionRecive(conn)

	return responsedMsg
}

// Watting for recive a msg by the connection
func ConnectionRecive(conn net.Conn) string {
	bufferRead := make([]byte, 1024)
	messageLeng, errReading := conn.Read(bufferRead)
	responsedMsg := string(bufferRead[:messageLeng])

	if errReading != nil {
		log.Fatalln("Error reciving response...\nConnection: " + conn.LocalAddr().String() +
			"\nError: " + errReading.Error() +
			"\nResponsed message: " + responsedMsg)
	}

	log.Println("Connection message recived.\nConnection: " + conn.LocalAddr().String() + "\nMessage: " + responsedMsg)

	return responsedMsg
}

// Continous watting for a message
func ConnectionListen(conn net.Conn, toExec func(string)) {
	for {
		log.Println("Waitting for message in connection: " + conn.LocalAddr().String())
		toExec(ConnectionRecive(conn))
	}
}
