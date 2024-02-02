package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"discord-bot/internal/handlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	// Create a new Discord session using the bot token
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	// Register the Router function as the event handler
	session.AddHandler(handlers.Router)
	// Open a connection to Discord
	session.Open()

	// CLose a connection to discord after function finish
	defer session.Close()
	fmt.Println("Bot running...")

	// bot will close after interrupt with ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
