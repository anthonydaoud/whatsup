package server

import (
	"github.com/brown-csci1380/whatsup/whatsup"

	"encoding/gob"
	"fmt"
	"net"
)

var hashy map[string]whatsup.ChatConn
var reverse map[whatsup.ChatConn]string

func handleConnection(conn net.Conn) {
	// TODO: Implement handling messages from a client.
	// You will find whatsup.SendMsg and whatsup.RecvMsg methods useful for
	// serializing and deserializing messages
	chatConn := whatsup.ChatConn{}
	chatConn.Conn = conn
	chatConn.Enc = gob.NewEncoder(conn)
	chatConn.Dec = gob.NewDecoder(conn)
	for {
		chatMsg, err := whatsup.RecvMsg(chatConn)
		if err == nil {
			switch chatMsg.Action {
			case whatsup.CONNECT:
				fmt.Println(chatMsg.Username + " has connected")
				hashy[chatMsg.Username] = chatConn
				reverse[chatConn] = chatMsg.Username
			case whatsup.MSG:
				endConnection := hashy[chatMsg.Username]
				msg := whatsup.WhatsUpMsg{ Username: chatMsg.Username, Body: chatMsg.Body, Action: whatsup.MSG}
				whatsup.SendMsg(endConnection, msg)
			case whatsup.LIST:
				body := "Users connected:\n"
				for k := range hashy {	
					body += string(k) + "\n"
				}
				msg := whatsup.WhatsUpMsg{ Body: body, Action: whatsup.LIST }
				whatsup.SendMsg(chatConn, msg)
			case whatsup.ERROR:
				fmt.Println("Error not yet implemented")
			case whatsup.DISCONNECT:
				usrNam := reverse[chatConn]
				fmt.Println("Disconnecting " + usrNam)
				delete(reverse, chatConn)
				delete(hashy, usrNam)
				conn.Close()
				return
			default:
				fmt.Println("Not a valid Action")
			}
		}
	}

}

func Start() {
	listen, port, err := whatsup.OpenListener()
	fmt.Printf("Listening on port %d\n", port)
	hashy = make(map[string]whatsup.ChatConn)
	reverse = make(map[whatsup.ChatConn]string)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept() // this blocks until connection or error
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
