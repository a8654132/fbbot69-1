package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"template"
)

var mess = &Messenger{}

func main() {
	port := os.Getenv("PORT")
	log.Println("Server start in port:", port)
	mess.VerifyToken = os.Getenv("TOKEN")
	mess.AccessToken = os.Getenv("TOKEN")
	log.Println("Bot start in token:", mess.VerifyToken)
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)
	mess.SendSimpleMessage("1460870680701162", fmt.Sprintf("如果你看到這個，\n就代表我成功主動傳送訊息囉！"))
	button := template.NewWebURLButton("點此看阿卡莉", "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=62861397")
	mess.SendSimpleMessage("1460870680701162", button)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//MessageReceived :Callback to handle when message received.
func MessageReceived(event Event, opts MessageOpts, msg ReceivedMessage) {
	// log.Println("event:", event, " opt:", opts, " msg:", msg)
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你好，現在是被動的回復訊息。\n你的ID為%s\n你剛剛說的話為：%s", opts.Sender.ID ,msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
