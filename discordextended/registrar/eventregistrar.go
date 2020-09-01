package registrar

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterChannelCreateEvent(function func(session *discordgo.Session, event *discordgo.ChannelCreate)) {
}

func RegisterChannelDeleteEventHandler(function func(session *discordgo.Session, event *discordgo.ChannelCreate)) {
}

func RegisterChannelPinsUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.ChannelPinsUpdate)) {
}

func RegisterChannelUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.ChannelUpdate)) {
}

func RegisterConnectEventHandler(function func(session *discordgo.Session, event *discordgo.Connect)) {
}

func RegisterDisconnectEventHandler(function func(session *discordgo.Session, event *discordgo.Disconnect)) {
}

func RegisterEventEventHandler(function func(session *discordgo.Session, event *discordgo.Event)) {}

func RegisterGuildBanAddEventHandler(function func(session *discordgo.Session, event *discordgo.GuildBanAdd)) {
}

func RegisterGuildBanRemoveEventHandler(function func(session *discordgo.Session, event *discordgo.GuildBanRemove)) {
}

func RegisterGuildCreateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildCreate)) {
}

func RegisterGuildDeleteEventHandler(function func(session *discordgo.Session, event *discordgo.GuildDelete)) {
}

func RegisterGuildEmojisUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildEmojisUpdate)) {
}

func RegisterGuildIntegrationsUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildIntegrationsUpdate)) {
}
func RegisterGuildMemberAddEventHandler(function func(session *discordgo.Session, event *discordgo.GuildMemberAdd)) {
}

func RegisterGuildMemberRemoveEventHandler(function func(session *discordgo.Session, event *discordgo.GuildMemberRemove)) {
}

func RegisterGuildMemberUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildMemberUpdate)) {
}

func RegisterGuildMembersChunkEventHandler(function func(session *discordgo.Session, event *discordgo.GuildMembersChunk)) {
}

func RegisterGuildRoleCreateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildRoleCreate)) {
}

func RegisterGuildRoleDeleteEventHandler(function func(session *discordgo.Session, event *discordgo.GuildRoleDelete)) {
}

func RegisterGuildRoleUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildRoleUpdate)) {
}

func RegisterGuildUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.GuildUpdate)) {
}

func RegisterMessageAckEventHandler(function func(session *discordgo.Session, event *discordgo.MessageAck)) {
}

func RegisterMessageCreateEventHandler(function func(session *discordgo.Session, event *discordgo.MessageCreate)) {
}

func RegisterMessageDeleteEventHandler(function func(session *discordgo.Session, event *discordgo.MessageDelete)) {
}

func RegisterMessageDeleteBulkEventHandler(function func(session *discordgo.Session, event *discordgo.MessageDeleteBulk)) {
}

func RegisterMessageReactionAddEventHandler(function func(session *discordgo.Session, event *discordgo.MessageReactionAdd)) {
}

func RegisterMessageReactionRemoveEventHandler(function func(session *discordgo.Session, event *discordgo.MessageReactionRemove)) {
}
func RegisterMessageReactionRemoveAllEventHandler(function func(session *discordgo.Session, event *discordgo.MessageReactionRemoveAll)) {
}
func RegisterMessageUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.MessageUpdate)) {
}

func RegisterPresenceUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.PresenceUpdate)) {
}

func RegisterPresencesReplaceEventHandler(function func(session *discordgo.Session, event *discordgo.PresencesReplace)) {
}

func RegisterRateLimitEventHandler(function func(session *discordgo.Session, event *discordgo.RateLimit)) {
}

func RegisterReadyEventHandler(function func(session *discordgo.Session, event *discordgo.Ready)) {}

func RegisterRelationshipAddEventHandler(function func(session *discordgo.Session, event *discordgo.RelationshipAdd)) {
}

func RegisterRelationshipRemoveEventHandler(function func(session *discordgo.Session, event *discordgo.RelationshipRemove)) {
}

func RegisterResumedEventHandler(function func(session *discordgo.Session, event *discordgo.Resumed)) {
}

func RegisterTypingStartEventHandler(function func(session *discordgo.Session, event *discordgo.TypingStart)) {
}

func RegisterUserGuildSettingsUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.UserGuildSettingsUpdate)) {
}

func RegisterUserNoteUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.UserNoteUpdate)) {
}

func RegisterUserSettingsUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.UserSettingsUpdate)) {
}

func RegisterUserUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.UserUpdate)) {
}

func RegisterVoiceServerUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.VoiceServerUpdate)) {
}

func RegisterVoiceStateUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.VoiceStateUpdate)) {
}

func RegisterWebhooksUpdateEventHandler(function func(session *discordgo.Session, event *discordgo.WebhooksUpdate)) {
}
