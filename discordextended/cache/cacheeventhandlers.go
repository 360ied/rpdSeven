package cache

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterEventHandlers(bot *discordgo.Session) {
	bot.AddHandler(channelCreateEventHandler)
	bot.AddHandler(channelDeleteEventHandler)
	bot.AddHandler(channelPinsUpdateEventHandler)
	bot.AddHandler(channelUpdateEventHandler)
	bot.AddHandler(connectEventHandler)
	bot.AddHandler(disconnectEventHandler)
	bot.AddHandler(eventEventHandler)
	bot.AddHandler(guildBanAddEventHandler)
	bot.AddHandler(guildBanRemoveEventHandler)
	bot.AddHandler(guildCreateEventHandler)
	bot.AddHandler(guildDeleteEventHandler)
	bot.AddHandler(guildEmojisUpdateEventHandler)
	bot.AddHandler(guildIntegrationsUpdateEventHandler)
	bot.AddHandler(guildMemberAddEventHandler)
	bot.AddHandler(guildMemberRemoveEventHandler)
	bot.AddHandler(guildMemberUpdateEventHandler)
	bot.AddHandler(guildMembersChunkEventHandler)
	bot.AddHandler(guildRoleCreateEventHandler)
	bot.AddHandler(guildRoleDeleteEventHandler)
	bot.AddHandler(guildRoleUpdateEventHandler)
	bot.AddHandler(guildUpdateEventHandler)
	bot.AddHandler(messageAckEventHandler)
	bot.AddHandler(messageCreateEventHandler)
	bot.AddHandler(messageDeleteEventHandler)
	bot.AddHandler(messageDeleteBulkEventHandler)
	bot.AddHandler(messageReactionAddEventHandler)
	bot.AddHandler(messageReactionRemoveEventHandler)
	bot.AddHandler(messageReactionRemoveAllEventHandler)
	bot.AddHandler(messageUpdateEventHandler)
	bot.AddHandler(presenceUpdateEventHandler)
	bot.AddHandler(presencesReplaceEventHandler)
	bot.AddHandler(rateLimitEventHandler)
	bot.AddHandler(readyEventHandler)
	bot.AddHandler(relationshipAddEventHandler)
	bot.AddHandler(relationshipRemoveEventHandler)
	bot.AddHandler(resumedEventHandler)
	bot.AddHandler(typingStartEventHandler)
	bot.AddHandler(userGuildSettingsUpdateEventHandler)
	bot.AddHandler(userNoteUpdateEventHandler)
	bot.AddHandler(userSettingsUpdateEventHandler)
	bot.AddHandler(userUpdateEventHandler)
	bot.AddHandler(voiceServerUpdateEventHandler)
	bot.AddHandler(voiceStateUpdateEventHandler)
	bot.AddHandler(webhooksUpdateEventHandler)
}

func channelCreateEventHandler(session *discordgo.Session, event *discordgo.ChannelCreate) {

}

func channelDeleteEventHandler(session *discordgo.Session, event *discordgo.ChannelDelete) {
	for _, function := range eventfuncs.ChannelDeleteEventHandlers {
		go function(session, event)
	}
}

func channelPinsUpdateEventHandler(session *discordgo.Session, event *discordgo.ChannelPinsUpdate) {
	for _, function := range eventfuncs.ChannelPinsUpdateEventHandlers {
		go function(session, event)
	}
}

func channelUpdateEventHandler(session *discordgo.Session, event *discordgo.ChannelUpdate) {
	for _, function := range eventfuncs.ChannelUpdateEventHandlers {
		go function(session, event)
	}
}

func connectEventHandler(session *discordgo.Session, event *discordgo.Connect) {
	for _, function := range eventfuncs.ConnectEventHandlers {
		go function(session, event)
	}
}

func disconnectEventHandler(session *discordgo.Session, event *discordgo.Disconnect) {
	for _, function := range eventfuncs.DisconnectEventHandlers {
		go function(session, event)
	}
}

func eventEventHandler(session *discordgo.Session, event *discordgo.Event) {
	for _, function := range eventfuncs.EventEventHandlers {
		go function(session, event)
	}
}

func guildBanAddEventHandler(session *discordgo.Session, event *discordgo.GuildBanAdd) {
	for _, function := range eventfuncs.GuildBanAddEventHandlers {
		go function(session, event)
	}
}

func guildBanRemoveEventHandler(session *discordgo.Session, event *discordgo.GuildBanRemove) {
	for _, function := range eventfuncs.GuildBanRemoveEventHandlers {
		go function(session, event)
	}
}

func guildCreateEventHandler(session *discordgo.Session, event *discordgo.GuildCreate) {
	for _, function := range eventfuncs.GuildCreateEventHandlers {
		go function(session, event)
	}
}

func guildDeleteEventHandler(session *discordgo.Session, event *discordgo.GuildDelete) {
	for _, function := range eventfuncs.GuildDeleteEventHandlers {
		go function(session, event)
	}
}

func guildEmojisUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildEmojisUpdate) {
	for _, function := range eventfuncs.GuildEmojisUpdateEventHandlers {
		go function(session, event)
	}
}

func guildIntegrationsUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildIntegrationsUpdate) {
	for _, function := range eventfuncs.GuildIntegrationsUpdateEventHandlers {
		go function(session, event)
	}
}

func guildMemberAddEventHandler(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	for _, function := range eventfuncs.GuildMemberAddEventHandlers {
		go function(session, event)
	}
}

func guildMemberRemoveEventHandler(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	for _, function := range eventfuncs.GuildMemberRemoveEventHandlers {
		go function(session, event)
	}
}

func guildMemberUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildMemberUpdate) {
	for _, function := range eventfuncs.GuildMemberUpdateEventHandlers {
		go function(session, event)
	}
}

func guildMembersChunkEventHandler(session *discordgo.Session, event *discordgo.GuildMembersChunk) {
	for _, function := range eventfuncs.GuildMembersChunkEventHandlers {
		go function(session, event)
	}
}

func guildRoleCreateEventHandler(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
	for _, function := range eventfuncs.GuildRoleCreateEventHandlers {
		go function(session, event)
	}
}

func guildRoleDeleteEventHandler(session *discordgo.Session, event *discordgo.GuildRoleDelete) {
	for _, function := range eventfuncs.GuildRoleDeleteEventHandlers {
		go function(session, event)
	}
}

func guildRoleUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildRoleUpdate) {
	for _, function := range eventfuncs.GuildRoleUpdateEventHandlers {
		go function(session, event)
	}
}

func guildUpdateEventHandler(session *discordgo.Session, event *discordgo.GuildUpdate) {
	for _, function := range eventfuncs.GuildUpdateEventHandlers {
		go function(session, event)
	}
}

func messageAckEventHandler(session *discordgo.Session, event *discordgo.MessageAck) {
	for _, function := range eventfuncs.MessageAckEventHandlers {
		go function(session, event)
	}
}

func messageCreateEventHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	for _, function := range eventfuncs.MessageCreateEventHandlers {
		go function(session, event)
	}
}

func messageDeleteEventHandler(session *discordgo.Session, event *discordgo.MessageDelete) {
	for _, function := range eventfuncs.MessageDeleteEventHandlers {
		go function(session, event)
	}
}

func messageDeleteBulkEventHandler(session *discordgo.Session, event *discordgo.MessageDeleteBulk) {
	for _, function := range eventfuncs.MessageDeleteBulkEventHandlers {
		go function(session, event)
	}
}

func messageReactionAddEventHandler(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	for _, function := range eventfuncs.MessageReactionAddEventHandlers {
		go function(session, event)
	}
}

func messageReactionRemoveEventHandler(session *discordgo.Session, event *discordgo.MessageReactionRemove) {
	for _, function := range eventfuncs.MessageReactionRemoveEventHandlers {
		go function(session, event)
	}
}

func messageReactionRemoveAllEventHandler(session *discordgo.Session, event *discordgo.MessageReactionRemoveAll) {
	for _, function := range eventfuncs.MessageReactionRemoveAllEventHandlers {
		go function(session, event)
	}
}

func messageUpdateEventHandler(session *discordgo.Session, event *discordgo.MessageUpdate) {
	for _, function := range eventfuncs.MessageUpdateEventHandlers {
		go function(session, event)
	}
}

func presenceUpdateEventHandler(session *discordgo.Session, event *discordgo.PresenceUpdate) {
	for _, function := range eventfuncs.PresenceUpdateEventHandlers {
		go function(session, event)
	}
}

func presencesReplaceEventHandler(session *discordgo.Session, event *discordgo.PresencesReplace) {
	for _, function := range eventfuncs.PresencesReplaceEventHandlers {
		go function(session, event)
	}
}

func rateLimitEventHandler(session *discordgo.Session, event *discordgo.RateLimit) {
	for _, function := range eventfuncs.RateLimitEventHandlers {
		go function(session, event)
	}
}

func readyEventHandler(session *discordgo.Session, event *discordgo.Ready) {
	for _, function := range eventfuncs.ReadyEventHandlers {
		go function(session, event)
	}
}

func relationshipAddEventHandler(session *discordgo.Session, event *discordgo.RelationshipAdd) {
	for _, function := range eventfuncs.RelationshipAddEventHandlers {
		go function(session, event)
	}
}

func relationshipRemoveEventHandler(session *discordgo.Session, event *discordgo.RelationshipRemove) {
	for _, function := range eventfuncs.RelationshipRemoveEventHandlers {
		go function(session, event)
	}
}

func resumedEventHandler(session *discordgo.Session, event *discordgo.Resumed) {
	for _, function := range eventfuncs.ResumedEventHandlers {
		go function(session, event)
	}
}

func typingStartEventHandler(session *discordgo.Session, event *discordgo.TypingStart) {
	for _, function := range eventfuncs.TypingStartEventHandlers {
		go function(session, event)
	}
}

func userGuildSettingsUpdateEventHandler(session *discordgo.Session, event *discordgo.UserGuildSettingsUpdate) {
	for _, function := range eventfuncs.UserGuildSettingsUpdateEventHandlers {
		go function(session, event)
	}
}

func userNoteUpdateEventHandler(session *discordgo.Session, event *discordgo.UserNoteUpdate) {
	for _, function := range eventfuncs.UserNoteUpdateEventHandlers {
		go function(session, event)
	}
}

func userSettingsUpdateEventHandler(session *discordgo.Session, event *discordgo.UserSettingsUpdate) {
	for _, function := range eventfuncs.UserSettingsUpdateEventHandlers {
		go function(session, event)
	}
}

func userUpdateEventHandler(session *discordgo.Session, event *discordgo.UserUpdate) {
	for _, function := range eventfuncs.UserUpdateEventHandlers {
		go function(session, event)
	}
}

func voiceServerUpdateEventHandler(session *discordgo.Session, event *discordgo.VoiceServerUpdate) {
	for _, function := range eventfuncs.VoiceServerUpdateEventHandlers {
		go function(session, event)
	}
}

func voiceStateUpdateEventHandler(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	for _, function := range eventfuncs.VoiceStateUpdateEventHandlers {
		go function(session, event)
	}
}

func webhooksUpdateEventHandler(session *discordgo.Session, event *discordgo.WebhooksUpdate) {
	for _, function := range eventfuncs.WebhooksUpdateEventHandlers {
		go function(session, event)
	}
}
