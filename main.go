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

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(handlers.Router)

	session.Open()

	defer session.Close()
	fmt.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
