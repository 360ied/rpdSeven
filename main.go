package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"

	"rpdSeven/keepalive"
)

func main() {
	// so that repl.it won't exit after the page is closed
	go keepalive.KeepAlive()

	var bot, err = discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		panic(err)
	}

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

func channelCreateEventHandler(session *discordgo.Session, event *discordgo.ChannelCreate) {}

func channelDeleteEventHandler(session *discordgo.Session, event *discordgo.ChannelCreate) {}

func channelPinsUpdateEventHandler(session *discordgo.Session, event *discordgo.ChannelPinsUpdate) {}

func channelUpdateEventHandler(session *discordgo.Session, event *discordgo.ChannelUpdate) {}

func connectEventHandler(session *discordgo.Session, event *discordgo.Connect) {}

func disconnectEventHandler(session *discordgo.Session, event *discordgo.Disconnect) {}

func eventEventHandler(session *discordgo.Session, event *discordgo.Event) {}

func guildBanAddEventHandler(session *discordgo.Session, event *discordgo.GuildBanAdd) {}

func guildBanRemoveEventHandler(session *discordgo.Session, event *discordgo.GuildBanRemove) {}

func guildCreateEventHandler(session *discordgo.Session, event *discordgo.GuildCreate) {}

func guildDeleteEventHandler(session *discordgo.Session, event *discordgo.GuildDelete) {}

func guildEmojisUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildEmojisUpdate) {}

func guildIntegrationsUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildIntegrationsUpdate) {
}

func guildMemberAddEventHandler(session *discordgo.Session, event *discordgo.GuildMemberAdd) {}

func guildMemberRemoveEventHandler(session *discordgo.Session, event *discordgo.GuildMemberRemove) {}

func guildMemberUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildMemberUpdate) {}

func guildMembersChunkEventHandler(session *discordgo.Session, event *discordgo.GuildMembersChunk) {}

func guildRolesCreateEventHandler(session *discordgo.Session, event *discordgo.GuildRoleCreate) {}

func guildRoleDeleteEventHandler(session *discordgo.Session, event *discordgo.GuildRoleDelete) {}

func guildRoleUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildRoleUpdate) {}

func guildUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildUpdate) {}

func messageAckEventHandler(session *discordgo.Session, event *discordgo.MessageAck) {}

func messageCreateEventHandler(session *discordgo.Session, event *discordgo.MessageCreate) {}

func messageDeleteEventHandler(session *discordgo.Session, event *discordgo.MessageDelete) {}

func messageDeleteBulkEventHandler(session *discordgo.Session, event *discordgo.MessageDeleteBulk) {}

func messageReactionAddEventHandler(session *discordgo.Session, event *discordgo.MessageReactionAdd) {}

func messageReactionRemoveEventHandler(session *discordgo.Session, event *discordgo.MessageReactionRemove) {
}

func messageReactionRemoveAllEventHandler(session *discordgo.Session, event *discordgo.MessageReactionRemoveAll) {
}

func messageUpdateEventHandler(session *discordgo.Session, event *discordgo.MessageUpdate) {}

func presenceUpdateEventHandler(session *discordgo.Session, event *discordgo.PresenceUpdate) {}

func presencesReplaceEventHandler(session *discordgo.Session, event *discordgo.PresencesReplace) {}

func rateLimitEventHandler(session *discordgo.Session, event *discordgo.RateLimit) {}

func readyEventHandler(session *discordgo.Session, event *discordgo.Ready) {}

func relationshipAddEventHandler(session *discordgo.Session, event *discordgo.RelationshipAdd) {}

func relationshipRemoveEventHandler(session *discordgo.Session, event *discordgo.RelationshipRemove) {}

func resumedEventHandler(session *discordgo.Session, event *discordgo.Resumed) {}

func typingStartEventHandler(session *discordgo.Session, event *discordgo.TypingStart) {}

func userGuildSettingsUpdateEventHandler(session *discordgo.Session, event *discordgo.UserGuildSettingsUpdate) {
}

func userNoteUpdateEventHandler(session *discordgo.Session, event *discordgo.UserNoteUpdate) {}

func userSettingsUpdateEventHandler(session *discordgo.Session, event *discordgo.UserSettingsUpdate) {}

func userUpdateEventHandler(session *discordgo.Session, event *discordgo.UserUpdate) {}

func voiceServerUpdateEventHandler(session *discordgo.Session, event *discordgo.VoiceServerUpdate) {}

func voiceStateUpdateEventHandler(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {}

func webhooksUpdateEventHandler(session *discordgo.Session, event *discordgo.WebhooksUpdate) {}
