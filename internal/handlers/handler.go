package handlers

import (
	"strings"
	"sync"

	"discord-bot/internal/services"

	"github.com/bwmarrin/discordgo"
)

var mu sync.Mutex

// function to call services asynchronously
func callServiceAsync(s *discordgo.Session, m *discordgo.MessageCreate, service func(*discordgo.Session, *discordgo.MessageCreate)) {
	go func() {
		mu.Lock()
		defer mu.Unlock()
		service(s, m)
	}()
}

// Main routing function to determine which service to call
func Router(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// variation of services
	switch {
	case strings.Contains(m.Content, "!help"):
		callServiceAsync(s, m, services.Helper)
	case strings.Contains(m.Content, "!poll"):
		callServiceAsync(s, m, services.PollService)
	case strings.Contains(m.Content, "!weather"):
		callServiceAsync(s, m, services.WeatherCheck)
	case strings.Contains(m.Content, "!reminder"):
		callServiceAsync(s, m, services.Reminder)
	case strings.Contains(m.Content, "!play"):
		callServiceAsync(s, m, services.Game)

	}
}
