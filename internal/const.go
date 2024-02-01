package internal

const (
	Divider = "|"

	PollHelperText = "**Usage:**\n- **Create a poll with multiple choices:** \n `!poll | <question> | <option1> | <option2> | ...`\n- **Create a poll with options YES/NO:** \n `!poll | <question>`"

	PollResultHeader       = "**Poll:** %s\n\n"
	EmojiCheck             = ":white_check_mark: %s\n"
	EmojyX                 = ":x: %s\n"
	TextYes                = "Yes"
	TextNo                 = "No"
	EmojiLetter            = ":regional_indicator_%s: %s\n"
	GeocodeLink            = "http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s"
	WeatherInformationLink = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric"
	HelpFunctionText       = "** Available command: **\n- `!help`: See all available commands\n- `!poll`: Create a poll\n- `!reminder`: Set a notification\n- `!play`: Play rock-paper-scissors with Miras\n- `!weather`: Check weather in specific location"
	CantFindCityText       = "Wrong city name :bangbang:"
	ReminderHelperText     = "**To set reminders, use the following command:**\n`!reminder | time (0h0m0s) | Text for description`"
	WrongTimeFormat        = "Wrong time format :bangbang:"

	AlertTimeSetText = "**Time was set, and will reminde you in %s.**"
	ReminderText     = " **TIME UP** <@%s>, \n*Description*:%s "
	YouWonText       = "Hey.. you won,congruts! :smiley:"
	YouLoseText      = "Better luck next time :kissing_heart:"
	DrawText         = "Draw :neutral_face:"
	SRPHelperText    = "**To play rock, paper, scissors, use the following command:**\n`!play | {one of the elements: scissor, rock, paper}`"
	ResultMessage    = "**Your result:%s**\n- You got:%s\n- Miras got:%s\n- Miras:%s"

	WeatherResultText = "**Result**\n- City:%s\n- Status:%s\n- Temperature:%g\n- Feels like:%g"

	WeatherErrorText = "**To find out the weather of a specific city, use the following command:**\n`!weather | <city name>`"
)
