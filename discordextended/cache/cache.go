package cache

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	CacheV = Cache{
		Session:       nil,
		guildsRWMutex: nil,
		Guilds:        nil,
	} 
)

type Cache struct {
	Session *discordgo.Session

	guildsRWMutex *sync.RWMutex
	Guilds        map[string]*discordgo.Guild
}

func (cache *Cache) GetGuild(ID string) (*discordgo.Guild, error) {
	var cachedGuild = func() *discordgo.Guild {
		cache.guildsRWMutex.RLock()
		defer cache.guildsRWMutex.RUnlock()
		var cachedGuild, cachedGuildExists = cache.Guilds[ID]
		if !cachedGuildExists {
			return nil
		}
		return cachedGuild
	}()
	if cachedGuild != nil {
		return cachedGuild, nil
	}
	var fetchedGuild, fetchedGuildErr = cache.Session.Guild(ID)
	if fetchedGuildErr != nil {
		return fetchedGuild, fetchedGuildErr
	}
	cache.guildsRWMutex.Lock()
	defer cache.guildsRWMutex.Unlock()
	cache.Guilds[ID] = fetchedGuild
	return fetchedGuild, nil
}
