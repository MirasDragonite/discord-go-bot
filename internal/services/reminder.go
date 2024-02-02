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

	reminder, err := time.ParseDuration(strings.TrimSpace(dates[1]))
	if err != nil {
		fmt.Println(err, reminder)
		s.ChannelMessageSendReply(m.ChannelID, internal.WrongTimeFormat, m.Reference())
		return
	}

	currentTime := time.Now()
	remindeTime := currentTime.Add(reminder)

	s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.AlertTimeSetText, remindeTime.Format(time.Kitchen)), m.Reference())

	time.AfterFunc(reminder, func() {
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf(internal.ReminderText, m.Author.ID, dates[2]), m.Reference())
		return
	})
}
