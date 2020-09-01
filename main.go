package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"

	"rpdSeven/discordextended"
	"rpdSeven/keepalive"
)

func main() {
	// so that repl.it won't exit after the page is closed
	go keepalive.KeepAlive()

	var bot, err = discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		panic(err)
	}

	discordextended.RegisterEventHandlers(bot)

	err = bot.Open()

	if err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}

	fmt.Println("Bot is now running.")

	// wait forever
	select {}

	// unreachable code
	// bot.Close()
}
