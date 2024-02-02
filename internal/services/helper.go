package services

import (
	"discord-bot/internal"

	"github.com/bwmarrin/discordgo"
)

// function that send messages with all availables functions
func Helper(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, internal.HelpFunctionText, m.Reference())
}
