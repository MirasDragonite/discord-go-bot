package services

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func PollServive(s *discordgo.Session, m *discordgo.MessageCreate) {
	scratch := strings.Split(m.Content, "|")

	poll := make([]string, 0)
	for _, ch := range scratch {
		if strings.TrimSpace(ch) != "" {
			poll = append(poll, ch)
		}
	}

	if len(poll) > 2 && len(poll) <= 16 {
		// Creation poll with multiple options
		question := poll[1]
		options := poll[2:]
		pollMessage := multipleQuestionPoll(question, options)
		s.ChannelMessageSend(m.ChannelID, pollMessage)
		return
	} else if len(poll) == 2 {
		// When i get only question,that will mean the question is simple, which contain only two option YES/NO
		pollMessage := simpleQuestionPoll(poll[1])
		s.ChannelMessageSend(m.ChannelID, pollMessage)
		return
	} else {
		// If non of this example above
		s.ChannelMessageSend(m.ChannelID, "Usage: \n **If you wanna create poll with multiple choise:**\n!poll | <question> | <option1> |  <option2> ... \n**If you wanna create poll with options YES/NO:**\n!poll | <question>")
		return
	}
}

func simpleQuestionPoll(question string) string {
	pollMessage := fmt.Sprintf("**Poll:** %s\n\n", question)

	pollMessage += fmt.Sprintf(":white_check_mark: %s\n", "Yes")
	pollMessage += fmt.Sprintf(":x: %s\n", "No")
	return pollMessage
}

func multipleQuestionPoll(question string, options []string) string {
	fmt.Println(question, options)
	pollMessage := fmt.Sprintf("**Poll:** %s\n\n", question)

	for i, option := range options {
		if strings.TrimSpace(option) != "" {
			pollMessage += fmt.Sprintf(":regional_indicator_%s: %s\n", string('a'+i), option)
		}
	}

	return pollMessage
}
