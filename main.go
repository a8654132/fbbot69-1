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
	content := Redis_IDtoMAC(opts.Sender.ID)
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你好，現在是被動的回復訊息。\n你的ID為%s\n你剛剛說的話為：%s\n\n%s", opts.Sender.ID ,msg.Text,content))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
