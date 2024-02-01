package internal

const (
	Divider                = "|"
	PollHelperText         = "Usage: \n **If you wanna create poll with multiple choise:**\n!poll | <question> | <option1> |  <option2> ... \n**If you wanna create poll with options YES/NO:**\n!poll | <question>"
	PollResultHeader       = "**Poll:** %s\n\n"
	EmojiCheck             = ":white_check_mark: %s\n"
	EmojyX                 = ":x: %s\n"
	TextYes                = "Yes"
	TextNo                 = "No"
	EmojiLetter            = ":regional_indicator_%s: %s\n"
	CantFindCityText       = "Wrong city name"
	WeatherResultText      = "   **Result**   \nCity:%s\nStatus:%s\ntemperature:%g\nFeels like:%g"
	GeocodeLink            = "http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s"
	WeatherInformationLink = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric"
	HelpFunctionText       = "**Available command:**\n!help-to see all available commands \n!poll-to create poll"
	WeatherErrorText       = "You can find out  result of specifiy city with command:\n!weather | <city name> "
	ReminderHelperText     = "To set reminders you should run the command:\n**!reminder | time (0h0m0s) | Text for description**"
	WrongTimeFormat        = "Wrong time format"
	AlertTimeSetText       = "Time was set, and will reminde you in %s"
	ReminderText           = " **TIME UP** <@%s>, \nDescription:%s "
)
