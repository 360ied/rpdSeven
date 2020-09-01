package ev

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ChannelCreateEventHandlers []func(*discordgo.Session, *discordgo.ChannelCreate)

	ChannelDeleteEventHandlers []func(*discordgo.Session, *discordgo.ChannelDelete)

	ChannelPinsUpdateEventHandlers []func(*discordgo.Session, *discordgo.ChannelPinsUpdate)

	ChannelUpdateEventHandlers []func(*discordgo.Session, *discordgo.ChannelUpdate)

	ConnectEventHandlers []func(*discordgo.Session, *discordgo.Connect)

	DisconnectEventHandlers []func(*discordgo.Session, *discordgo.Disconnect)

	EventEventHandlers []func(*discordgo.Session, *discordgo.Event)

	GuildBanAddEventHandlers []func(*discordgo.Session, *discordgo.GuildBanAdd)

	GuildBanRemoveEventHandlers []func(*discordgo.Session, *discordgo.GuildBanRemove)

	GuildCreateEventHandlers []func(*discordgo.Session, *discordgo.GuildCreate)

	GuildDeleteEventHandlers []func(*discordgo.Session, *discordgo.GuildDelete)

	GuildEmojisUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildEmojisUpdate)

	GuildIntegrationsUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildIntegrationsUpdate)

	GuildMemberAddEventHandlers []func(*discordgo.Session, *discordgo.GuildMemberAdd)

	GuildMemberRemoveEventHandlers []func(*discordgo.Session, *discordgo.GuildMemberRemove)

	GuildMemberUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildMemberUpdate)

	GuildMembersChunkEventHandlers []func(*discordgo.Session, *discordgo.GuildMembersChunk)

	GuildRoleCreateEventHandlers []func(*discordgo.Session, *discordgo.GuildRoleCreate)

	GuildRoleDeleteEventHandlers []func(*discordgo.Session, *discordgo.GuildRoleDelete)

	GuildRoleUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildRoleUpdate)

	GuildUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildUpdate)

	MessageAckEventHandlers []func(*discordgo.Session, *discordgo.MessageAck)

	MessageCreateEventHandlers []func(*discordgo.Session, *discordgo.MessageCreate)

	MessageDeleteEventHandlers []func(*discordgo.Session, *discordgo.MessageDelete)

	MessageDeleteBulkEventHandlers []func(*discordgo.Session, *discordgo.MessageDeleteBulk)

	MessageReactionAddEventHandlers []func(*discordgo.Session, *discordgo.MessageReactionAdd)

	MessageReactionRemoveEventHandlers []func(*discordgo.Session, *discordgo.MessageReactionRemove)

	MessageReactionRemoveAllEventHandlers []func(*discordgo.Session, *discordgo.MessageReactionRemoveAll)

	MessageUpdateEventHandlers []func(*discordgo.Session, *discordgo.MessageUpdate)

	PresenceUpdateEventHandlers []func(*discordgo.Session, *discordgo.PresenceUpdate)

	PresencesReplaceEventHandlers []func(*discordgo.Session, *discordgo.PresencesReplace)

	RateLimitEventHandlers []func(*discordgo.Session, *discordgo.RateLimit)

	ReadyEventHandlers []func(*discordgo.Session, *discordgo.Ready)

	RelationshipAddEventHandlers []func(*discordgo.Session, *discordgo.RelationshipAdd)

	RelationshipRemoveEventHandlers []func(*discordgo.Session, *discordgo.RelationshipRemove)

	ResumedEventHandlers []func(*discordgo.Session, *discordgo.Resumed)

	TypingStartEventHandlers []func(*discordgo.Session, *discordgo.TypingStart)

	UserGuildSettingsUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserGuildSettingsUpdate)

	UserNoteUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserNoteUpdate)

	UserSettingsUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserSettingsUpdate)

	UserUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserUpdate)

	VoiceServerUpdateEventHandlers []func(*discordgo.Session, *discordgo.VoiceServerUpdate)

	VoiceStateUpdateEventHandlers []func(*discordgo.Session, *discordgo.VoiceStateUpdate)

	WebhooksUpdateEventHandlers []func(*discordgo.Session, *discordgo.WebhooksUpdate)
)
