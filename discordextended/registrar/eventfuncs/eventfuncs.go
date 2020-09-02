package eventfuncs

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	ChannelCreateEventMutex    = sync.Mutex{}
	ChannelCreateEventHandlers []func(*discordgo.Session, *discordgo.ChannelCreate)

	ChannelDeleteEventMutex    = sync.Mutex{}
	ChannelDeleteEventHandlers []func(*discordgo.Session, *discordgo.ChannelDelete)

	ChannelPinsUpdateEventMutex    = sync.Mutex{}
	ChannelPinsUpdateEventHandlers []func(*discordgo.Session, *discordgo.ChannelPinsUpdate)

	ChannelUpdateEventMutex    = sync.Mutex{}
	ChannelUpdateEventHandlers []func(*discordgo.Session, *discordgo.ChannelUpdate)

	ConnectEventMutex    = sync.Mutex{}
	ConnectEventHandlers []func(*discordgo.Session, *discordgo.Connect)

	DisconnectEventMutex    = sync.Mutex{}
	DisconnectEventHandlers []func(*discordgo.Session, *discordgo.Disconnect)

	EventEventMutex    = sync.Mutex{}
	EventEventHandlers []func(*discordgo.Session, *discordgo.Event)

	GuildBanAddEventMutex    = sync.Mutex{}
	GuildBanAddEventHandlers []func(*discordgo.Session, *discordgo.GuildBanAdd)

	GuildBanRemoveEventMutex    = sync.Mutex{}
	GuildBanRemoveEventHandlers []func(*discordgo.Session, *discordgo.GuildBanRemove)

	GuildCreateEventMutex    = sync.Mutex{}
	GuildCreateEventHandlers []func(*discordgo.Session, *discordgo.GuildCreate)

	GuildDeleteEventMutex    = sync.Mutex{}
	GuildDeleteEventHandlers []func(*discordgo.Session, *discordgo.GuildDelete)

	GuildEmojisUpdateEventMutex    = sync.Mutex{}
	GuildEmojisUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildEmojisUpdate)

	GuildIntegrationsUpdateEventMutex    = sync.Mutex{}
	GuildIntegrationsUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildIntegrationsUpdate)

	GuildMemberAddEventMutex    = sync.Mutex{}
	GuildMemberAddEventHandlers []func(*discordgo.Session, *discordgo.GuildMemberAdd)

	GuildMemberRemoveEventMutex    = sync.Mutex{}
	GuildMemberRemoveEventHandlers []func(*discordgo.Session, *discordgo.GuildMemberRemove)

	GuildMemberUpdateEventMutex    = sync.Mutex{}
	GuildMemberUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildMemberUpdate)

	GuildMembersChunkEventMutex    = sync.Mutex{}
	GuildMembersChunkEventHandlers []func(*discordgo.Session, *discordgo.GuildMembersChunk)

	GuildRoleCreateEventMutex    = sync.Mutex{}
	GuildRoleCreateEventHandlers []func(*discordgo.Session, *discordgo.GuildRoleCreate)

	GuildRoleDeleteEventMutex    = sync.Mutex{}
	GuildRoleDeleteEventHandlers []func(*discordgo.Session, *discordgo.GuildRoleDelete)

	GuildRoleUpdateEventMutex    = sync.Mutex{}
	GuildRoleUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildRoleUpdate)

	GuildUpdateEventMutex    = sync.Mutex{}
	GuildUpdateEventHandlers []func(*discordgo.Session, *discordgo.GuildUpdate)

	MessageAckEventMutex    = sync.Mutex{}
	MessageAckEventHandlers []func(*discordgo.Session, *discordgo.MessageAck)

	MessageCreateEventMutex    = sync.Mutex{}
	MessageCreateEventHandlers []func(*discordgo.Session, *discordgo.MessageCreate)

	MessageDeleteEventMutex    = sync.Mutex{}
	MessageDeleteEventHandlers []func(*discordgo.Session, *discordgo.MessageDelete)

	MessageDeleteBulkEventMutex    = sync.Mutex{}
	MessageDeleteBulkEventHandlers []func(*discordgo.Session, *discordgo.MessageDeleteBulk)

	MessageReactionAddEventMutex    = sync.Mutex{}
	MessageReactionAddEventHandlers []func(*discordgo.Session, *discordgo.MessageReactionAdd)

	MessageReactionRemoveEventMutex    = sync.Mutex{}
	MessageReactionRemoveEventHandlers []func(*discordgo.Session, *discordgo.MessageReactionRemove)

	MessageReactionRemoveAllEventMutex    = sync.Mutex{}
	MessageReactionRemoveAllEventHandlers []func(*discordgo.Session, *discordgo.MessageReactionRemoveAll)

	MessageUpdateEventMutex    = sync.Mutex{}
	MessageUpdateEventHandlers []func(*discordgo.Session, *discordgo.MessageUpdate)

	PresenceUpdateEventMutex    = sync.Mutex{}
	PresenceUpdateEventHandlers []func(*discordgo.Session, *discordgo.PresenceUpdate)

	PresencesReplaceEventMutex    = sync.Mutex{}
	PresencesReplaceEventHandlers []func(*discordgo.Session, *discordgo.PresencesReplace)

	RateLimitEventMutex    = sync.Mutex{}
	RateLimitEventHandlers []func(*discordgo.Session, *discordgo.RateLimit)

	ReadyEventMutex    = sync.Mutex{}
	ReadyEventHandlers []func(*discordgo.Session, *discordgo.Ready)

	RelationshipAddEventMutex    = sync.Mutex{}
	RelationshipAddEventHandlers []func(*discordgo.Session, *discordgo.RelationshipAdd)

	RelationshipRemoveEventMutex    = sync.Mutex{}
	RelationshipRemoveEventHandlers []func(*discordgo.Session, *discordgo.RelationshipRemove)

	ResumedEventMutex    = sync.Mutex{}
	ResumedEventHandlers []func(*discordgo.Session, *discordgo.Resumed)

	TypingStartEventMutex    = sync.Mutex{}
	TypingStartEventHandlers []func(*discordgo.Session, *discordgo.TypingStart)

	UserGuildSettingsUpdateEventMutex    = sync.Mutex{}
	UserGuildSettingsUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserGuildSettingsUpdate)

	UserNoteUpdateEventMutex    = sync.Mutex{}
	UserNoteUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserNoteUpdate)

	UserSettingsUpdateEventMutex    = sync.Mutex{}
	UserSettingsUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserSettingsUpdate)

	UserUpdateEventMutex    = sync.Mutex{}
	UserUpdateEventHandlers []func(*discordgo.Session, *discordgo.UserUpdate)

	VoiceServerUpdateEventMutex    = sync.Mutex{}
	VoiceServerUpdateEventHandlers []func(*discordgo.Session, *discordgo.VoiceServerUpdate)

	VoiceStateUpdateEventMutex    = sync.Mutex{}
	VoiceStateUpdateEventHandlers []func(*discordgo.Session, *discordgo.VoiceStateUpdate)

	WebhooksUpdateEventMutex    = sync.Mutex{}
	WebhooksUpdateEventHandlers []func(*discordgo.Session, *discordgo.WebhooksUpdate)
)
