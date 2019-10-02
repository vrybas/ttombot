// Package main provides ...
package main

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

const tgBotAPIKey = "951904228:AAHCWIjOMKvF5LtocVqwDuNsSw2tF5-ujeo"

func main() {
	bot, err := tg.NewBotAPI(tgBotAPIKey)
	if err != nil {
		log.Fatalf("Error creating bot: %s\n", err)
	}

	botUser, err := bot.GetMe()
	if err != nil {
		log.Fatalf("Error getting bot info: %s\n", err)
	}
	fmt.Printf("%s: Authentication successful!\n", botUser.FirstName)

	updConfig := tg.UpdateConfig{
		Offset:  0,
		Limit:   10,
		Timeout: 60,
	}
	updChan, err := bot.GetUpdatesChan(updConfig)
	if err != nil {
		log.Fatalf("Error getting updates channel: %s\n", err)
	}

	for {
		update := <-updChan

		if update.Message != nil {
			switch {
			case update.Message.IsCommand():
				switch update.Message.Command() {
				case "menu":
					mainMenu := tg.NewReplyKeyboard(
						tg.NewKeyboardButtonRow(
							tg.NewKeyboardButton("😆 One"),
							tg.NewKeyboardButton("🤪 Two"),
						),
					)

					reply := tg.NewMessage(update.Message.Chat.ID, "This is help")
					reply.ReplyMarkup = mainMenu
					bot.Send(reply)
				case "help":
					reply := tg.NewMessage(update.Message.Chat.ID, "This is help")
					bot.Send(reply)
				default:
					reply := tg.NewMessage(update.Message.Chat.ID, "no such command")
					bot.Send(reply)
				}
			default:
				fmt.Printf(
					"chatID: %v, from: %s, message: %s\n",
					update.Message.Chat.ID,
					update.Message.From.String(),
					update.Message.Text,
				)

				reply := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
				reply.ReplyMarkup = nil
				bot.Send(reply)
			}
		}
	}
}
