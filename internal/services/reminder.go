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
		s.ChannelMessageSend(m.ChannelID, internal.ReminderHelperText)
		return
	}

	reminder, err := time.ParseDuration(strings.TrimSpace(dates[1]))
	if err != nil {
		fmt.Println(err, reminder)
		s.ChannelMessageSend(m.ChannelID, internal.WrongTimeFormat)
		return
	}

	currentTime := time.Now()
	remindeTime := currentTime.Add(reminder)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(internal.AlertTimeSetText, remindeTime.Format(time.Kitchen)))

	time.AfterFunc(reminder, func() {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(internal.ReminderText, m.Author.ID, dates[2]))
		return
	})
}
