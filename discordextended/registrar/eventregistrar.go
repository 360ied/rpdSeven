package registrar

import (
	"github.com/bwmarrin/discordgo"

	"rpdSeven/discordextended/registrar/eventfuncs"
)

func RegisterChannelCreateEvent(function func(*discordgo.Session, *discordgo.ChannelCreate)) {
	eventfuncs.ChannelCreateEventMutex.Lock()
	defer eventfuncs.ChannelCreateEventMutex.Unlock()
	eventfuncs.ChannelCreateEventHandlers = append(eventfuncs.ChannelCreateEventHandlers, function)
}

func RegisterChannelDeleteEventHandler(function func(*discordgo.Session, *discordgo.ChannelDelete)) {
	eventfuncs.ChannelDeleteEventMutex.Lock()
	defer eventfuncs.ChannelDeleteEventMutex.Unlock()
	eventfuncs.ChannelDeleteEventHandlers = append(eventfuncs.ChannelDeleteEventHandlers, function)
}

func RegisterChannelPinsUpdateEventHandler(function func(*discordgo.Session, *discordgo.ChannelPinsUpdate)) {
	eventfuncs.ChannelPinsUpdateEventMutex.Lock()
	defer eventfuncs.ChannelPinsUpdateEventMutex.Unlock()
	eventfuncs.ChannelPinsUpdateEventHandlers = append(eventfuncs.ChannelPinsUpdateEventHandlers, function)
}

func RegisterChannelUpdateEventHandler(function func(*discordgo.Session, *discordgo.ChannelUpdate)) {
	eventfuncs.ChannelUpdateEventMutex.Lock()
	defer eventfuncs.ChannelUpdateEventMutex.Unlock()
	eventfuncs.ChannelUpdateEventHandlers = append(eventfuncs.ChannelUpdateEventHandlers, function)
}

func RegisterConnectEventHandler(function func(*discordgo.Session, *discordgo.Connect)) {
	eventfuncs.ConnectEventMutex.Lock()
	defer eventfuncs.ConnectEventMutex.Unlock()
	eventfuncs.ConnectEventHandlers = append(eventfuncs.ConnectEventHandlers, function)
}

func RegisterDisconnectEventHandler(function func(*discordgo.Session, *discordgo.Disconnect)) {
	eventfuncs.DisconnectEventMutex.Lock()
	defer eventfuncs.DisconnectEventMutex.Unlock()
	eventfuncs.DisconnectEventHandlers = append(eventfuncs.DisconnectEventHandlers, function)
}

func RegisterEventEventHandler(function func(*discordgo.Session, *discordgo.Event)) {
	eventfuncs.EventEventMutex.Lock()
	defer eventfuncs.EventEventMutex.Unlock()
	eventfuncs.EventEventHandlers = append(eventfuncs.EventEventHandlers, function)
}

func RegisterGuildBanAddEventHandler(function func(*discordgo.Session, *discordgo.GuildBanAdd)) {
	eventfuncs.GuildBanAddEventMutex.Lock()
	defer eventfuncs.GuildBanAddEventMutex.Unlock()
	eventfuncs.GuildBanAddEventHandlers = append(eventfuncs.GuildBanAddEventHandlers, function)
}

func RegisterGuildBanRemoveEventHandler(function func(*discordgo.Session, *discordgo.GuildBanRemove)) {
	eventfuncs.GuildBanRemoveEventMutex.Lock()
	defer eventfuncs.GuildBanRemoveEventMutex.Unlock()
	eventfuncs.GuildBanRemoveEventHandlers = append(eventfuncs.GuildBanRemoveEventHandlers, function)
}

func RegisterGuildCreateEventHandler(function func(*discordgo.Session, *discordgo.GuildCreate)) {
	eventfuncs.GuildCreateEventMutex.Lock()
	defer eventfuncs.GuildCreateEventMutex.Unlock()
	eventfuncs.GuildCreateEventHandlers = append(eventfuncs.GuildCreateEventHandlers, function)
}

func RegisterGuildDeleteEventHandler(function func(*discordgo.Session, *discordgo.GuildDelete)) {
	eventfuncs.GuildDeleteEventMutex.Lock()
	defer eventfuncs.GuildDeleteEventMutex.Unlock()
	eventfuncs.GuildDeleteEventHandlers = append(eventfuncs.GuildDeleteEventHandlers, function)
}

func RegisterGuildEmojisUpdateEventHandler(function func(*discordgo.Session, *discordgo.GuildEmojisUpdate)) {
	eventfuncs.GuildEmojisUpdateEventMutex.Lock()
	defer eventfuncs.GuildEmojisUpdateEventMutex.Unlock()
	eventfuncs.GuildEmojisUpdateEventHandlers = append(eventfuncs.GuildEmojisUpdateEventHandlers, function)
}

func RegisterGuildIntegrationsUpdateEventHandler(function func(*discordgo.Session, *discordgo.GuildIntegrationsUpdate)) {
	eventfuncs.GuildIntegrationsUpdateEventMutex.Lock()
	defer eventfuncs.GuildIntegrationsUpdateEventMutex.Unlock()
	eventfuncs.GuildIntegrationsUpdateEventHandlers = append(eventfuncs.GuildIntegrationsUpdateEventHandlers, function)

}
func RegisterGuildMemberAddEventHandler(function func(*discordgo.Session, *discordgo.GuildMemberAdd)) {
	eventfuncs.GuildMemberAddEventMutex.Lock()
	defer eventfuncs.GuildMemberAddEventMutex.Unlock()
	eventfuncs.GuildMemberAddEventHandlers = append(eventfuncs.GuildMemberAddEventHandlers, function)
}

func RegisterGuildMemberRemoveEventHandler(function func(*discordgo.Session, *discordgo.GuildMemberRemove)) {
	eventfuncs.GuildMemberRemoveEventMutex.Lock()
	defer eventfuncs.GuildMemberRemoveEventMutex.Unlock()
	eventfuncs.GuildMemberRemoveEventHandlers = append(eventfuncs.GuildMemberRemoveEventHandlers, function)
}

func RegisterGuildMemberUpdateEventHandler(function func(*discordgo.Session, *discordgo.GuildMemberUpdate)) {
	eventfuncs.GuildMemberUpdateEventMutex.Lock()
	defer eventfuncs.GuildMemberUpdateEventMutex.Unlock()
	eventfuncs.GuildMemberUpdateEventHandlers = append(eventfuncs.GuildMemberUpdateEventHandlers, function)
}

func RegisterGuildMembersChunkEventHandler(function func(*discordgo.Session, *discordgo.GuildMembersChunk)) {
	eventfuncs.GuildMembersChunkEventMutex.Lock()
	defer eventfuncs.GuildMembersChunkEventMutex.Unlock()
	eventfuncs.GuildMembersChunkEventHandlers = append(eventfuncs.GuildMembersChunkEventHandlers, function)
}

func RegisterGuildRoleCreateEventHandler(function func(*discordgo.Session, *discordgo.GuildRoleCreate)) {
	eventfuncs.GuildRoleCreateEventMutex.Lock()
	defer eventfuncs.GuildRoleCreateEventMutex.Unlock()
	eventfuncs.GuildRoleCreateEventHandlers = append(eventfuncs.GuildRoleCreateEventHandlers, function)
}

func RegisterGuildRoleDeleteEventHandler(function func(*discordgo.Session, *discordgo.GuildRoleDelete)) {
	eventfuncs.GuildRoleDeleteEventMutex.Lock()
	defer eventfuncs.GuildRoleDeleteEventMutex.Unlock()
	eventfuncs.GuildRoleDeleteEventHandlers = append(eventfuncs.GuildRoleDeleteEventHandlers, function)
}

func RegisterGuildRoleUpdateEventHandler(function func(*discordgo.Session, *discordgo.GuildRoleUpdate)) {
	eventfuncs.GuildRoleUpdateEventMutex.Lock()
	defer eventfuncs.GuildRoleUpdateEventMutex.Unlock()
	eventfuncs.GuildRoleUpdateEventHandlers = append(eventfuncs.GuildRoleUpdateEventHandlers, function)
}

func RegisterGuildUpdateEventHandler(function func(*discordgo.Session, *discordgo.GuildUpdate)) {
	eventfuncs.GuildUpdateEventMutex.Lock()
	defer eventfuncs.GuildUpdateEventMutex.Unlock()
	eventfuncs.GuildUpdateEventHandlers = append(eventfuncs.GuildUpdateEventHandlers, function)
}

func RegisterMessageAckEventHandler(function func(*discordgo.Session, *discordgo.MessageAck)) {
	eventfuncs.MessageAckEventMutex.Lock()
	defer eventfuncs.MessageAckEventMutex.Unlock()
	eventfuncs.MessageAckEventHandlers = append(eventfuncs.MessageAckEventHandlers, function)
}

func RegisterMessageCreateEventHandler(function func(*discordgo.Session, *discordgo.MessageCreate)) {
	eventfuncs.MessageCreateEventMutex.Lock()
	defer eventfuncs.MessageCreateEventMutex.Unlock()
	eventfuncs.MessageCreateEventHandlers = append(eventfuncs.MessageCreateEventHandlers, function)
}

func RegisterMessageDeleteEventHandler(function func(*discordgo.Session, *discordgo.MessageDelete)) {
	eventfuncs.MessageDeleteEventMutex.Lock()
	defer eventfuncs.MessageDeleteEventMutex.Unlock()
	eventfuncs.MessageDeleteEventHandlers = append(eventfuncs.MessageDeleteEventHandlers, function)
}

func RegisterMessageDeleteBulkEventHandler(function func(*discordgo.Session, *discordgo.MessageDeleteBulk)) {
	eventfuncs.MessageDeleteBulkEventMutex.Lock()
	defer eventfuncs.MessageDeleteBulkEventMutex.Unlock()
	eventfuncs.MessageDeleteBulkEventHandlers = append(eventfuncs.MessageDeleteBulkEventHandlers, function)
}

func RegisterMessageReactionAddEventHandler(function func(*discordgo.Session, *discordgo.MessageReactionAdd)) {
	eventfuncs.MessageReactionAddEventMutex.Lock()
	defer eventfuncs.MessageReactionAddEventMutex.Unlock()
	eventfuncs.MessageReactionAddEventHandlers = append(eventfuncs.MessageReactionAddEventHandlers, function)
}

func RegisterMessageReactionRemoveEventHandler(function func(*discordgo.Session, *discordgo.MessageReactionRemove)) {
	eventfuncs.MessageReactionRemoveEventMutex.Lock()
	defer eventfuncs.MessageReactionRemoveEventMutex.Unlock()
	eventfuncs.MessageReactionRemoveEventHandlers = append(eventfuncs.MessageReactionRemoveEventHandlers, function)
}
func RegisterMessageReactionRemoveAllEventHandler(function func(*discordgo.Session, *discordgo.MessageReactionRemoveAll)) {
	eventfuncs.MessageReactionRemoveAllEventMutex.Lock()
	defer eventfuncs.MessageReactionRemoveAllEventMutex.Unlock()
	eventfuncs.MessageReactionRemoveAllEventHandlers = append(eventfuncs.MessageReactionRemoveAllEventHandlers, function)
}
func RegisterMessageUpdateEventHandler(function func(*discordgo.Session, *discordgo.MessageUpdate)) {
	eventfuncs.MessageUpdateEventMutex.Lock()
	defer eventfuncs.MessageUpdateEventMutex.Unlock()
	eventfuncs.MessageUpdateEventHandlers = append(eventfuncs.MessageUpdateEventHandlers, function)
}

func RegisterPresenceUpdateEventHandler(function func(*discordgo.Session, *discordgo.PresenceUpdate)) {
	eventfuncs.PresenceUpdateEventMutex.Lock()
	defer eventfuncs.PresenceUpdateEventMutex.Unlock()
	eventfuncs.PresenceUpdateEventHandlers = append(eventfuncs.PresenceUpdateEventHandlers, function)
}

func RegisterPresencesReplaceEventHandler(function func(*discordgo.Session, *discordgo.PresencesReplace)) {
	eventfuncs.PresencesReplaceEventMutex.Lock()
	defer eventfuncs.PresencesReplaceEventMutex.Unlock()
	eventfuncs.PresencesReplaceEventHandlers = append(eventfuncs.PresencesReplaceEventHandlers, function)
}

func RegisterRateLimitEventHandler(function func(*discordgo.Session, *discordgo.RateLimit)) {
	eventfuncs.RateLimitEventMutex.Lock()
	defer eventfuncs.RateLimitEventMutex.Unlock()
	eventfuncs.RateLimitEventHandlers = append(eventfuncs.RateLimitEventHandlers, function)
}

func RegisterReadyEventHandler(function func(*discordgo.Session, *discordgo.Ready)) {
	eventfuncs.ReadyEventMutex.Lock()
	defer eventfuncs.ReadyEventMutex.Unlock()
	eventfuncs.ReadyEventHandlers = append(eventfuncs.ReadyEventHandlers, function)
}

func RegisterRelationshipAddEventHandler(function func(*discordgo.Session, *discordgo.RelationshipAdd)) {
	eventfuncs.RelationshipAddEventMutex.Lock()
	defer eventfuncs.RelationshipAddEventMutex.Unlock()
	eventfuncs.RelationshipAddEventHandlers = append(eventfuncs.RelationshipAddEventHandlers, function)
}

func RegisterRelationshipRemoveEventHandler(function func(*discordgo.Session, *discordgo.RelationshipRemove)) {
	eventfuncs.RelationshipRemoveEventMutex.Lock()
	defer eventfuncs.RelationshipRemoveEventMutex.Unlock()
	eventfuncs.RelationshipRemoveEventHandlers = append(eventfuncs.RelationshipRemoveEventHandlers, function)
}

func RegisterResumedEventHandler(function func(*discordgo.Session, *discordgo.Resumed)) {
	eventfuncs.ResumedEventMutex.Lock()
	defer eventfuncs.ResumedEventMutex.Unlock()
	eventfuncs.ResumedEventHandlers = append(eventfuncs.ResumedEventHandlers, function)
}

func RegisterTypingStartEventHandler(function func(*discordgo.Session, *discordgo.TypingStart)) {
	eventfuncs.TypingStartEventMutex.Lock()
	defer eventfuncs.TypingStartEventMutex.Unlock()
	eventfuncs.TypingStartEventHandlers = append(eventfuncs.TypingStartEventHandlers, function)
}

func RegisterUserGuildSettingsUpdateEventHandler(function func(*discordgo.Session, *discordgo.UserGuildSettingsUpdate)) {
	eventfuncs.UserGuildSettingsUpdateEventMutex.Lock()
	defer eventfuncs.UserGuildSettingsUpdateEventMutex.Unlock()
	eventfuncs.UserGuildSettingsUpdateEventHandlers = append(eventfuncs.UserGuildSettingsUpdateEventHandlers, function)
}

func RegisterUserNoteUpdateEventHandler(function func(*discordgo.Session, *discordgo.UserNoteUpdate)) {
	eventfuncs.UserNoteUpdateEventMutex.Lock()
	defer eventfuncs.UserNoteUpdateEventMutex.Unlock()
	eventfuncs.UserNoteUpdateEventHandlers = append(eventfuncs.UserNoteUpdateEventHandlers, function)
}

func RegisterUserSettingsUpdateEventHandler(function func(*discordgo.Session, *discordgo.UserSettingsUpdate)) {
	eventfuncs.UserSettingsUpdateEventMutex.Lock()
	defer eventfuncs.UserSettingsUpdateEventMutex.Unlock()
	eventfuncs.UserSettingsUpdateEventHandlers = append(eventfuncs.UserSettingsUpdateEventHandlers, function)
}

func RegisterUserUpdateEventHandler(function func(*discordgo.Session, *discordgo.UserUpdate)) {
	eventfuncs.UserUpdateEventMutex.Lock()
	defer eventfuncs.UserUpdateEventMutex.Unlock()
	eventfuncs.UserUpdateEventHandlers = append(eventfuncs.UserUpdateEventHandlers, function)
}

func RegisterVoiceServerUpdateEventHandler(function func(*discordgo.Session, *discordgo.VoiceServerUpdate)) {
	eventfuncs.VoiceServerUpdateEventMutex.Lock()
	defer eventfuncs.VoiceServerUpdateEventMutex.Unlock()
	eventfuncs.VoiceServerUpdateEventHandlers = append(eventfuncs.VoiceServerUpdateEventHandlers, function)
}

func RegisterVoiceStateUpdateEventHandler(function func(*discordgo.Session, *discordgo.VoiceStateUpdate)) {
	eventfuncs.VoiceStateUpdateEventMutex.Lock()
	defer eventfuncs.VoiceStateUpdateEventMutex.Unlock()
	eventfuncs.VoiceStateUpdateEventHandlers = append(eventfuncs.VoiceStateUpdateEventHandlers, function)
}

func RegisterWebhooksUpdateEventHandler(function func(*discordgo.Session, *discordgo.WebhooksUpdate)) {
	eventfuncs.WebhooksUpdateEventMutex.Lock()
	defer eventfuncs.WebhooksUpdateEventMutex.Unlock()
	eventfuncs.WebhooksUpdateEventHandlers = append(eventfuncs.WebhooksUpdateEventHandlers, function)
}
