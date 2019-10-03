// Telegram bot
package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vrybas/ttombot/config"
)

type handleCmdFunc func(tg.Update)

var bot *tg.BotAPI
var updChan tg.UpdatesChannel
var mainMenu tg.ReplyKeyboardMarkup
var tpl *template.Template

func init() {
	bot = connectToBot()
	updChan = getUpdChan()
	mainMenu = buildMainMenu()
	tpl = template.Must(template.New("").ParseGlob("templates/*.tmpl"))
}

func main() {
	for {
		if update := <-updChan; update.Message != nil {
			switch {
			case update.Message.IsCommand():
				handleCmd(update)
			default:
				handleMsg(update)
			}
		}
	}
}

func handleCmd(update tg.Update) {
	cmdHandlers := map[string]handleCmdFunc{
		"menu": menuCmdHandler,
		"help": helpCmdHandler,
	}

	if handler, ok := cmdHandlers[update.Message.Command()]; ok {
		handler(update)
	} else {
		defaultCmdHandler(update)
	}
}

func menuCmdHandler(update tg.Update) {
	reply := tg.NewMessage(update.Message.Chat.ID, "Main Menu")
	reply.ReplyMarkup = mainMenu
	bot.Send(reply)
}

func helpCmdHandler(update tg.Update) {
	tplBuf := bytes.Buffer{}
	err := tpl.ExecuteTemplate(&tplBuf, "help.tmpl", nil)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	reply := tg.NewMessage(update.Message.Chat.ID, tplBuf.String())
	reply.ReplyMarkup = mainMenu
	bot.Send(reply)
}

func defaultCmdHandler(update tg.Update) {
	reply := tg.NewMessage(update.Message.Chat.ID, "no such command")
	reply.ReplyMarkup = mainMenu
	bot.Send(reply)
}

func handleMsg(update tg.Update) {
	fmt.Printf(
		"chatID: %v, from: %s, message: %s\n",
		update.Message.Chat.ID,
		update.Message.From.String(),
		update.Message.Text,
	)

	reply := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
	reply.ReplyMarkup = tg.NewRemoveKeyboard(true)
	bot.Send(reply)
}

func connectToBot() *tg.BotAPI {
	log.Println("Connecting to bot...")
	bot, err := tg.NewBotAPI(config.TG_BOTAPI_KEY)
	if err != nil {
		log.Fatalf("Error creating bot: %s\n", err)
	}

	botUser, err := bot.GetMe()
	if err != nil {
		log.Fatalf("Error getting bot info: %s\n", err)
	}
	log.Printf("%s: Authentication successful!\n", botUser.FirstName)

	return bot
}

func getUpdChan() tg.UpdatesChannel {
	updConfig := tg.UpdateConfig{
		Offset:  0,
		Limit:   10,
		Timeout: 60,
	}
	updChan, err := bot.GetUpdatesChan(updConfig)
	if err != nil {
		log.Fatalf("Error getting updates channel: %s\n", err)
	}
	return updChan
}

func buildMainMenu() tg.ReplyKeyboardMarkup {
	mainMenu = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("/menu"),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("/help"),
		),
	)

	return mainMenu
}
