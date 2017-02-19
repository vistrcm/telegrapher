package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/vistrcm/telegrams"
	"io/ioutil"
	"os"
	"os/user"
)

const (
	relativeConfigFilePath = "/.config/telegrapher.json"
	noConfigToken          = "tokenNotSet"
	noConfigChat           = "chatNotSet"
)

type Configuration struct {
	Token  string `json:"token"`
	ChatId string `json:"chat"`
}

func ReadConfig(relativeConfigFilePath string) (config Configuration, err error) {
	usr, err := user.Current()
	if err != nil {
		config = Configuration{Token: noConfigToken, ChatId: noConfigChat}
		return

	}
	configFileName := usr.HomeDir + relativeConfigFilePath
	raw, err := ioutil.ReadFile(configFileName)
	if err != nil {
		config = Configuration{Token: noConfigToken, ChatId: noConfigChat}
		return
	}

	config = Configuration{}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		config = Configuration{Token: noConfigToken, ChatId: noConfigChat}
		return
	}

	return config, nil

}

func main() {

	config, _ := ReadConfig(relativeConfigFilePath)

	var token = flag.String("token", config.Token, "bot token")
	var chat = flag.String("chat", config.ChatId, "chat ID to send message to")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("message text is empty. Use telegrapher [options] <message>")
		os.Exit(1)
	}

	message := flag.Arg(0)

	bot := telegrams.NewTelegramBot(*token)
	_, err := bot.SendMessage(*chat, message)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v. token: %v. chat: %v\n", err, *token, *chat)
		os.Exit(2)
	}
}
