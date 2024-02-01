package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"discord-bot/internal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func WeatherCheck(s *discordgo.Session, m *discordgo.MessageCreate) {
	text := strings.Split(m.Content, "|")

	if len(text) != 2 {
		s.ChannelMessageSend(m.ChannelID, internal.WeatherErrorText)
		return
	}

	coordinate := getCoordinate(text[1])

	if len(coordinate) > 0 {
	} else {
		s.ChannelMessageSend(m.ChannelID, internal.CantFindCityText)
		return
	}

	weather := getWeatherData(coordinate)

	result := fmt.Sprintf(internal.WeatherResultText, coordinate[0].Name, weather.Weth[0].Descriotion, weather.Main.Temp, weather.Main.Feels)

	s.ChannelMessageSend(m.ChannelID, result)

	return
}

func getCoordinate(cityName string) []internal.Coordinate {
	godotenv.Load()
	ApiKey := os.Getenv("WEATHER_API_KEY")
	data, err := http.Get(fmt.Sprintf(internal.GeocodeLink, cityName, ApiKey))
	if err != nil {
		fmt.Println("ERROR1:", err)
		log.Fatal(err)
	}

	var coordinate []internal.Coordinate
	err = json.NewDecoder(data.Body).Decode(&coordinate)
	if err != nil {
		fmt.Println("ERROR2:", err)
		log.Fatal(err)
	}
	return coordinate
}

func getWeatherData(coordinate []internal.Coordinate) internal.Weather {
	godotenv.Load()

	ApiKey := os.Getenv("WEATHER_API_KEY")

	weatherData, err := http.Get(fmt.Sprintf(internal.WeatherInformationLink, coordinate[0].Lat, coordinate[0].Lon, ApiKey))
	if err != nil {
		log.Fatal(err)
	}
	var weather internal.Weather
	err = json.NewDecoder(weatherData.Body).Decode(&weather)

	if err != nil {
		fmt.Println("Err3", err)
		log.Fatal(err)
	}
	return weather
}
