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

	mess.SendSimpleMessage("1460870680701162", "這是主動的，傳給1460870680701162")
	mess.SendSimpleMessage("2122993047726622", "這是主動的，傳給2122993047726622")

	http.HandleFunc("/webhook", mess.Handler)
	mess.MessageReceived = MessageReceived

	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	stringid := fmt.Sprintf("%s",opts.Sender.ID)
	// content := Redis_IDtoMAC(stringid)
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你好，現在是被動的回復訊息。\n你的ID為%s\n你剛剛說的話為：%s\n\n", opts.Sender.ID , msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
