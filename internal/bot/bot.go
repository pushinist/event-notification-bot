package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pushinist/events-notification-bot/internal/storage"
	"github.com/pushinist/events-notification-bot/internal/utils"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	db  *storage.Storage
}

func NewBot(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{bot: bot}, nil
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		b.processMessage(update)
	}

	return nil
}

func (b *Bot) processMessage(message tgbotapi.Message) {
	if eventTime, ok := utils.TryParseEventTime(message.Text); ok {
		
	}
}
