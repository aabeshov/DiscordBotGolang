package main

import (
	"DiscordBot/internal/bot"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatal("Error loading test.env file")
	}
	discordToken := os.Getenv("DISCORD_BOT_TOKEN")
	if discordToken == "" {
		log.Fatal("Discord bot token not set")
	}
	bot, err := bot.NewBot(discordToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	if err := bot.Start(); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}

	fmt.Println("Bot is running")
	bot.WaitForInterrupt()
	bot.Stop()
}
