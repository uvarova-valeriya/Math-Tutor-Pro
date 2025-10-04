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
			// Создаем обычную кнопку-ссылку
			webAppBtn := tgbotapi.NewInlineKeyboardButtonURL(
				"📊 Открыть приложение",
				"https://uvarova-valeriya.github.io/Math-Tutor-Pro",
			)

			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(webAppBtn),
			)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"🤖 Добро пожаловать в Math Tutor Pro!\n\n"+
					"Я помогу вам анализировать прогресс учеников и создавать персонализированные домашние задания.\n\n"+
					"Нажмите кнопку ниже чтобы открыть приложение!")
			msg.ReplyMarkup = keyboard

			bot.Send(msg)

		case "webapp":
			webAppBtn := tgbotapi.NewInlineKeyboardButtonURL(
				"📊 Открыть приложение",
				"https://uvarova-valeriya.github.io/Math-Tutor-Pro",
			)

			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(webAppBtn),
			)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"🔗 Откройте веб-приложение:")
			msg.ReplyMarkup = keyboard

			bot.Send(msg)
		}
	}
}
