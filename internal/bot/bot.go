package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
}

func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Session: dg,
	}

	dg.AddHandler(bot.messageCreate)

	return bot, nil
}

func (b *Bot) Start() error {
	return b.Session.Open()
}

func (b *Bot) Stop() {
	b.Session.Close()
}

func (b *Bot) WaitForInterrupt() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}
