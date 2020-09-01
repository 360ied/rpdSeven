package discordextended

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
