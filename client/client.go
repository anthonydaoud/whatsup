package client

import (
	"github.com/brown-csci1380/whatsup/whatsup"
	"bufio"
	"os"
	"fmt"
	"strings"
)

var k int

func handleListen(chatConn whatsup.ChatConn) int {
	for {
		chatMsg, err := whatsup.RecvMsg(chatConn)
		if err == nil {
			switch chatMsg.Action {
			case whatsup.MSG:
				fmt.Print("\n" + chatMsg.Username + ": " + chatMsg.Body + "\nEnter Action: ")
			case whatsup.LIST:
				if len(chatMsg.Body) > 0 {
					fmt.Print("\n" + chatMsg.Body + "Enter Action: ")
				}
			case whatsup.ERROR:
				fmt.Println("Error not yet implemented")
				os.Exit(-1)
			default:
				fmt.Println("Not a valid Action")
			}
		} else {
			fmt.Println(err)
			return 0
		}
	}

}
/*
func handleAction(chatConn whatsup.ChatConn) int  {
	for {
		msg := whatsup.WhatsUpMsg{}
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Action: ")
		act, err := buf.ReadBytes('\n')
		if err != nil {
	        fmt.Println(err)
	    } else {
	    	actStr := strings.TrimSpace(string(act))
	    	switch actStr {
	    	case "MSG":
	    		fmt.Print("Enter other Name: ")	
	    		nam, err1 := buf.ReadBytes('\n')
	    		if err1 != nil {
			        fmt.Println(err1)
			        continue
			    }
			    name := strings.TrimSpace(string(nam))
	    		fmt.Print("Enter Message: ")
				sentence, err2 := buf.ReadBytes('\n')
				if err2 != nil {
		        	fmt.Println(err2)
		        	continue
			    }
			    bod := strings.TrimSpace(string(sentence))
			    fmt.Println(whatsup.MSG)
			    msg = whatsup.WhatsUpMsg{ Username: name, Body: bod, Action: whatsup.MSG }
			case "LIST":
				msg = whatsup.WhatsUpMsg{ Action: whatsup.LIST }
			case "DISCONNECT":
				msg = whatsup.WhatsUpMsg{ Action: whatsup.DISCONNECT }
				whatsup.SendMsg(chatConn, msg)
				return 0
			default:
				fmt.Println("Not a valid action")
				continue
	    	}
	    	whatsup.SendMsg(chatConn, msg)
	    }
	}

}

*/

func Start(user string, serverPort string, serverAddr string) {
	// Connect to chat server
	chatConn, err := whatsup.ServerConnect(user, serverAddr, serverPort)
	if err != nil {
		fmt.Printf("unable to connect to server: %s\n", err)
		return
	}
	go handleListen(chatConn)
	for {
		msg := whatsup.WhatsUpMsg{}
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Action: ")
		act, err := buf.ReadBytes('\n')
		if err != nil {
	        fmt.Println(err)
	    } else {
	    	actStr := strings.TrimSpace(string(act))
	    	switch actStr {
	    	case "MSG":
	    		fmt.Print("Enter other Name: ")	
	    		nam, err1 := buf.ReadBytes('\n')
	    		if err1 != nil {
			        fmt.Println(err1)
			        continue
			    }
			    name := strings.TrimSpace(string(nam))
	    		fmt.Print("Enter Message: ")
				sentence, err2 := buf.ReadBytes('\n')
				if err2 != nil {
		        	fmt.Println(err2)
		        	continue
			    }
			    bod := strings.TrimSpace(string(sentence))
			    msg = whatsup.WhatsUpMsg{ Username: name, Body: bod, Action: whatsup.MSG }
			case "LIST":
				msg = whatsup.WhatsUpMsg{ Action: whatsup.LIST }
			case "DISCONNECT":
				msg = whatsup.WhatsUpMsg{ Action: whatsup.DISCONNECT }
				whatsup.SendMsg(chatConn, msg)
				os.Exit(0)
			default:
				fmt.Println("Not a valid action")
				continue
	    	}
	    	whatsup.SendMsg(chatConn, msg)
	    }
	}
	

	// TODO: Receive input from the user and use the first return value of whatsup.ServerConnect
	// (currently ignored so the stencil will compile) to talk to the server.
}
