package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"

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
	SendButton();
	// mess.MessageReceived = handler


	// button := NewWebURLButton("點此看阿卡莉", "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=62861397")
	// mess.SendSimpleMessage("1460870680701162", button)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func SendButton() {

		binary, _ := Redis_Get(mac)
		user := new(USER_MAC)
		json.Unmarshal(binary,&user)

		for i:=0 ; i< len(user.USER) ; i++ {

		// for i:=0 ; i< 2 ; i++ {
			// mq := messenger.MessageQuery{}
			onlyonecontent := user.USER[i].CONTENT
			// mq.RecipientID("1460870680701162")
			// mq.Template(template.GenericTemplate {Title: "請告訴我們您是否滿意這篇文章：",
			// 	Buttons: []template.Button{
			// 		template.Button{
			// 			Type:    template.ButtonTypePostback,
			// 			Payload: "good",
			// 			Title:   "滿意",
			// 		},
			// 		template.Button{
			// 			Type:    template.ButtonTypePostback,
			// 			Payload: "bad",
			// 			Title:   "不滿意",
			// 		},
			// 		template.Button{
			// 			Type:    template.ButtonTypeWebURL,
			// 			Title:   "點此開啟網頁",
			// 			URL:		 user.USER[i].NAME,
			// 		},
			// 	},
			// })
			x := fmt.Sprintf("%d", i+1)
			mess.SendSimpleMessage("1460870680701162", x + ".\n" + onlyonecontent )
			// mess.SendMessage(mq)
		}

	}



//
//   // MessageReceived :Callback to handle when message received.
// func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
// 	// log.Println("event:", event, " opt:", opts, " msg:", msg)
// 	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你好，現在是被動的回復訊息。\n你的ID為%s\n你剛剛說的話為：%s", opts.Sender.ID ,msg.Text))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%+v", resp)
// }
