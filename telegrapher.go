package main

import (
	"github.com/vistrcm/telegrams"
	"flag"
	"fmt"
	"os"
)

func main() {
	const (
		defaultToken = "defaulttoken"
		defaultChatID = "111111111111"
	)


	var token = flag.String("token", defaultToken, "bot token")
	var chat = flag.String("chat", defaultChatID, "chat ID to send message to")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("message text is empty. Use telegrapher [options] <message>")
		os.Exit(1)
	}

	message := flag.Arg(0)

	bot := telegrams.NewTelegramBot(*token)
	_, err := bot.SendMessage(*chat, message)

	if err != nil{
		fmt.Println(err)
		panic("something goes wrong!" )
	}
}
