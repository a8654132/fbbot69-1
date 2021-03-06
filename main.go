package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "encoding/json"

	"github.com/maciekmm/messenger-platform-go-sdk"
	// "github.com/maciekmm/messenger-platform-go-sdk/template"
)

var mess = &messenger.Messenger{}

func main() {
	port := os.Getenv("PORT")
	log.Println("Server start in port:", port)
	mess.VerifyToken = os.Getenv("TOKEN")
	mess.AccessToken = os.Getenv("TOKEN")
	log.Println("Bot start in token:", mess.VerifyToken)

	http.HandleFunc("/webhook", mess.Handler)
	mess.MessageReceived = MessageReceived

	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	stringid := fmt.Sprintf("%s",opts.Sender.ID)
	content := Redis_IDtoMAC(stringid)
	fmt.Println(content)

	mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你的id為:%s",opts.Sender.ID))

	for i:=0 ; i< len(content) ; i++{
			mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("%s",content[i]))
	}
}
