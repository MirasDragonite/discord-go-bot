package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	fmt.Println(token)
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
}
