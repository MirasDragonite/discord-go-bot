package services

import (
	"fmt"
	"math/rand"
	"strings"

	"discord-bot/internal"

	"github.com/bwmarrin/discordgo"
)

func Game(s *discordgo.Session, m *discordgo.MessageCreate) {
	text := strings.Split(m.Content, internal.Divider)
	fmt.Println(text)
	if len(text) != 2 {
		s.ChannelMessageSend(m.ChannelID, internal.SRPHelperText)
		return
	}

	SRP := []string{"scissor", "paper", "rock"}
	mp := map[string]bool{}
	for _, ch := range SRP {
		mp[ch] = true
	}
	randNum := rand.Intn(len(SRP))

	AIresult := SRP[randNum]

	if !mp[strings.TrimSpace(text[1])] {
		s.ChannelMessageSendReply(m.ChannelID, internal.SRPHelperText, m.Reference())
		return
	}

	result := SRPGame(AIresult, strings.TrimSpace(text[1]))

	switch result {
	case 1:
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.ResultMessage, "Win", strings.TrimSpace(text[1]), AIresult, internal.YouWonText), m.Reference())
		return
	case -1:
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.ResultMessage, "Lose", strings.TrimSpace(text[1]), AIresult, internal.YouLoseText), m.Reference())
		return
	case 0:
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.ResultMessage, "Draw", strings.TrimSpace(text[1]), AIresult, internal.DrawText), m.Reference())
		return
	}
}

func SRPGame(ai, human string) int {
	if ai == "scissor" && human == "paper" {
		return -1
	} else if ai == "scissor" && human == "rock" {
		return 1
	} else if ai == "rock" && human == "scissor" {
		return -1
	} else if ai == "rock" && human == "paper" {
		return 1
	} else if ai == "paper" && human == "scissor" {
		return 1
	} else if ai == "paper" && human == "rock" {
		return -1
	}
	return 0
}
