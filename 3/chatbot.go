package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", callbackHandler)

	port := "8080"
	fmt.Println("Server running at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			http.Error(w, "Invalid Signature", http.StatusBadRequest)
		} else {
			http.Error(w, "Error parsing request", http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				handleTextMessage(event.ReplyToken, message.Text)
			}
		}
	}
}

func handleTextMessage(replyToken, receivedText string) {
	var reply []linebot.SendingMessage

	switch receivedText {
	case "เมนูขนมยอดฮิต":
		reply = append(reply, linebot.NewTemplateMessage("เมนูขนมยอดฮิต",
			&linebot.ButtonsTemplate{
				Title: "📌 เลือกเมนูที่ต้องการ",
				Text:  "เมนูพิเศษของเรา 🎂🍰",
				Actions: []linebot.TemplateAction{
					&linebot.MessageAction{Label: "🍩 โดนัท", Text: "ฉันต้องการสั่งโดนัท"},
					&linebot.MessageAction{Label: "🍪 คุกกี้", Text: "ฉันต้องการสั่งคุกกี้"},
				},
			}))

	case "ขนม":
		reply = append(reply, linebot.NewTextMessage("🧁 คุณต้องการขนมแบบไหน? เลือกด้านล่างเลย!").WithQuickReplies(
			linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton("", &linebot.MessageAction{Label: "🍰 เค้ก", Text: "ฉันต้องการสั่งเค้ก"}),
				linebot.NewQuickReplyButton("", &linebot.MessageAction{Label: "🥧 พาย", Text: "ฉันต้องการสั่งพาย"}),
				linebot.NewQuickReplyButton("", &linebot.MessageAction{Label: "🍩 โดนัท", Text: "ฉันต้องการสั่งโดนัท"}),
			),
		))

	case "โปรโมชั่น":
		reply = append(reply, linebot.NewTemplateMessage("🎀 โปรโมชั่นช่วงชัมเมอร์ 🎀",
			&linebot.CarouselTemplate{
				Columns: []*linebot.CarouselColumn{
					{
						Title: "🍩 โดนัท",
						Text:  "เลือกท็อปปิ้งเองได้! 39 บาท",
						Actions: []linebot.TemplateAction{
							&linebot.MessageAction{Label: "เลือกโดนัท", Text: "ฉันต้องการสั่งโดนัท"},
						},
					},
					{
						Title: "🍪 คุกกี้",
						Text:  "มีให้เลือกหลายรสชาติ! 20 บาท",
						Actions: []linebot.TemplateAction{
							&linebot.MessageAction{Label: "เลือกคุกกี้", Text: "ฉันต้องการสั่งคุกกี้"},
						},
					},
					{
						Title: "🎂 เค้ก",
						Text:  "เค้กโฮมเมดสุดอร่อย 50 บาท",
						Actions: []linebot.TemplateAction{
							&linebot.MessageAction{Label: "เลือกเค้ก", Text: "ฉันต้องการสั่งเค้ก"},
						},
					},
				},
			}))

	case "ฉันต้องการสั่งโดนัท":
		reply = append(reply,
			linebot.NewTextMessage("🍩 คุณเลือกโดนัทแล้ว! ทานให้อร่อยน้าา"),
			linebot.NewImageMessage(
				"https://assets.tastemadecdn.net/images/bed995/ad7f7d4c04914923b152/efb316.png",
				"https://assets.tastemadecdn.net/images/bed995/ad7f7d4c04914923b152/efb316.png",
			),
		)

	case "ฉันต้องการสั่งคุกกี้":
		reply = append(reply,
			linebot.NewTextMessage("🍪 คุณเลือกคุกกี้แล้ว! ทานให้อร่อยน้าา"),
			linebot.NewImageMessage(
				"https://assets.bonappetit.com/photos/5ca534485e96521ff23b382b/16:9/w_4800,h_2700,c_limit/chocolate-chip-cookie.jpg",
				"https://assets.bonappetit.com/photos/5ca534485e96521ff23b382b/16:9/w_4800,h_2700,c_limit/chocolate-chip-cookie.jpg",
			),
		)

	case "ฉันต้องการสั่งเค้ก":
		reply = append(reply,
			linebot.NewTextMessage("🎂 คุณเลือกเค้กแล้ว! ทานให้อร่อยน้าา"),
			linebot.NewImageMessage(
				"https://www.allrecipes.com/thmb/gDJ1S6ETLfWGyyWw_4A_IGhvDYE=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/9295_red-velvet-cake_ddmfs_4x3_1129-a8ab17b825e3464a9a53ceeda54ff461.jpg",
				"https://www.allrecipes.com/thmb/gDJ1S6ETLfWGyyWw_4A_IGhvDYE=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/9295_red-velvet-cake_ddmfs_4x3_1129-a8ab17b825e3464a9a53ceeda54ff461.jpg",
			),
		)
	case "ฉันต้องการสั่งพาย":
		reply = append(reply,
			linebot.NewTextMessage("🥧 คุณเลือกพายแล้ว! ทานให้อร่อยน้าา"),
			linebot.NewImageMessage(
				"https://img.kapook.com/u/2018/surauch/cooking/co1/apple.jpg",
				"https://img.kapook.com/u/2018/surauch/cooking/co1/apple.jpg",
			),
		)

	default:
		reply = append(reply, linebot.NewTextMessage("🍪 ยินดีต้อนรับสู่ BotBakery! ร้านขนม DIY ที่ให้คุณเลือกท็อปปิ้งเองได้ 🧁✨"))
		reply = append(reply, linebot.NewTextMessage("พิมพ์ : เมนูขนมยอดฮิต✨ เพื่อดูเมนูสุดฮิตของร้าน \nพิมพ์ : ขนม✨ เพื่อดูขนมในร้านทั้งหมด \nพิมพ์ : โปรโมชั่น✨ เพื่อดูโปรโมชั่นช่วงนี้  🍪✨"))
	}

	if _, err := bot.ReplyMessage(replyToken, reply...).Do(); err != nil {
		log.Print(err)
	}
}
