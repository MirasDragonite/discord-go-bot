package handlers

import (
	"strings"

	"discord-bot/internal/services"

	"github.com/bwmarrin/discordgo"
)

func Router(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case strings.Contains(m.Content, "!help"):
		s.ChannelMessageSend(m.ChannelID, "**Available command:**\n!help-to see all available commands \n!poll-to create poll")
	case strings.Contains(m.Content, "!poll"):
		services.PollServive(s, m)
	}
}
