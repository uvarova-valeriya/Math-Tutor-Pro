package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "start":
			// Простая клавиатура без Web App
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"🤖 Добро пожаловать в Math Tutor Pro!\n\n"+
					"Я помогу вам анализировать прогресс учеников и создавать персонализированные домашние задания.\n\n"+
					"Команды:\n"+
					"/progress - Прогресс учеников\n"+
					"/webapp - Получить ссылку на веб-приложение\n\n"+
					"🔗 Веб-приложение: https://your-app.com")

			bot.Send(msg)

		case "webapp":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"🔗 Ссылка на веб-приложение:\nhttps://your-app.com\n\n"+
					"(Web App кнопка будет добавлена в следующем обновлении)")
			bot.Send(msg)
		}
	}
}
