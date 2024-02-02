package services

import (
	"fmt"
	"strings"
	"time"

	"discord-bot/internal"

	"github.com/bwmarrin/discordgo"
)

func Reminder(s *discordgo.Session, m *discordgo.MessageCreate) {
	dates := strings.Split(m.Content, internal.Divider)

	if len(dates) != 3 {
		s.ChannelMessageSendReply(m.ChannelID, internal.ReminderHelperText, m.Reference())
		return
	}
	// to parse string into time
	reminder, err := time.ParseDuration(strings.TrimSpace(dates[1]))
	if err != nil {

		s.ChannelMessageSendReply(m.ChannelID, internal.WrongTimeFormat, m.Reference())
		return
	}

	currentTime := time.Now()
	// adding remind time to current
	remindeTime := currentTime.Add(reminder)

	s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.AlertTimeSetText, remindeTime.Format(time.Kitchen)), m.Reference())

	// function send message in the channel after time up
	time.AfterFunc(reminder, func() {
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.ReminderText, m.Author.ID, dates[2]), m.Reference())
		return
	})
}
