package services

import (
	"fmt"
	"strings"

	"discord-bot/internal"

	"github.com/bwmarrin/discordgo"
)

func PollServive(s *discordgo.Session, m *discordgo.MessageCreate) {
	scratch := strings.Split(m.Content, internal.Divider)

	poll := make([]string, 0)

	for _, ch := range scratch {
		if strings.TrimSpace(ch) != "" {
			poll = append(poll, ch)
		}
	}

	if len(poll) > 2 && len(poll) <= 16 {
		
		question := poll[1]

		options := poll[2:]

		pollMessage := multipleQuestionPoll(question, options)

		s.ChannelMessageSend(m.ChannelID, pollMessage)

	} else if len(poll) == 2 {
		// When i get only question,that will mean the question is simple, which contain only two option YES/NO

		pollMessage := simpleQuestionPoll(poll[1])

		s.ChannelMessageSend(m.ChannelID, pollMessage)

	} else {
		// If non of this example above

		s.ChannelMessageSend(m.ChannelID, internal.PollHelperText)
	}
}

func simpleQuestionPoll(question string) string {
	pollMessage := fmt.Sprintf(internal.PollResultHeader, question)

	pollMessage += fmt.Sprintf(internal.EmojiCheck, internal.TextYes)

	pollMessage += fmt.Sprintf(internal.EmojyX, internal.TextYes)

	return pollMessage
}

func multipleQuestionPoll(question string, options []string) string {
	pollMessage := fmt.Sprintf(internal.PollResultHeader, question)

	for i, option := range options {
		if strings.TrimSpace(option) != "" {
			pollMessage += fmt.Sprintf(internal.EmojiLetter, string('a'+i), option)
		}
	}

	return pollMessage
}
