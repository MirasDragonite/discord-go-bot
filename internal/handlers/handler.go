package handlers

import (
	"strings"
	"sync"

	"discord-bot/internal"
	"discord-bot/internal/services"

	"github.com/bwmarrin/discordgo"
)

var mu sync.Mutex

func Router(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case strings.Contains(m.Content, "!help"):
		mu.Lock()
		go func() {
			defer mu.Unlock()
			s.ChannelMessageSend(m.ChannelID, internal.HelpFunctionText)
		}()
	case strings.Contains(m.Content, "!poll"):
		mu.Lock()
		go func() {
			defer mu.Unlock()
			services.PollService(s, m)
		}()
	case strings.Contains(m.Content, "!weather"):
		mu.Lock()
		go func() {
			defer mu.Unlock()
			services.WeatherCheck(s, m)
		}()
	case strings.Contains(m.Content, "!reminder"):
		mu.Lock()
		go func() {
			defer mu.Unlock()
			services.Reminder(s, m)
		}()
	case strings.Contains(m.Content, "!play"):
		mu.Lock()
		go func() {
			defer mu.Unlock()
			services.Game(s, m)
		}()
	}
}
