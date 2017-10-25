package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "encoding/json"

	"github.com/maciekmm/messenger-platform-go-sdk"
	"github.com/maciekmm/messenger-platform-go-sdk/template"
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
	// stringid := fmt.Sprintf("%s",opts.Sender.ID)
	// content := Redis_IDtoMAC(stringid)


	// resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你好，現在是被動的回復訊息。\n你的ID為%s\n你剛剛說的話為：%s\n\n%s", opts.Sender.ID , msg.Text , content))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for i:=0;i<len(content);i++{
	// 	 mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("%s",content[i]))
	//
	// }

	mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你的id為:%s",opts.Sender.ID))

	mq := messenger.MessageQuery{}
	mq.RecipientID(opts.Sender.ID)
	mq.Template(template.GenericTemplate {Title: "請告訴我們您是否滿意這篇文章：",
		Buttons: []template.Button{
			template.Button{
				Type:    template.ButtonTypePostback,
				Payload: "good",
				Title:   "滿意",
			},
			template.Button{
				Type:    template.ButtonTypePostback,
				Payload: "bad",
				Title:   "不滿意",
			},
			template.Button{
				Type:    template.ButtonTypeWebURL,
				Title:   "點此開啟網頁",
				URL:		 "140.115.153.185",
			},
		},
	})

	mess.SendMessage(mq)

	// fmt.Printf("%+v", resp)
}
