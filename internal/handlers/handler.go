package handlers

import (
	"strings"

	"discord-bot/internal"
	"discord-bot/internal/services"

	"github.com/bwmarrin/discordgo"
)

func Router(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case strings.Contains(m.Content, "!help"):
		s.ChannelMessageSend(m.ChannelID, internal.HelpFunctionText)
	case strings.Contains(m.Content, "!poll"):
		services.PollServive(s, m)
	case strings.Contains(m.Content, "!weather"):
		services.WeatherCheck(s, m)
	case strings.Contains(m.Content, "!reminder"):
		services.Reminder(s, m)
	case strings.Contains(m.Content, "!play"):
		services.Game(s, m)
	}
}
