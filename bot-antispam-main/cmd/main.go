package main

import (
	"log"

	"diplomka/bot"
	"diplomka/handlers"
	"diplomka/utils"

	tele "gopkg.in/telebot.v3"
)

func main() {
	// Загружаем переменные окружения из файла .env
	if err := utils.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	// Настраиваем логирование
	utils.InitLogger()

	// Создаем нового бота
	token := utils.GetEnv("TOKEN")
	bot, err := bot.InitBot(token)
	if err != nil {
		log.Fatal(err)
	}

	myBot := handlers.NewBot(bot, make(map[string]int))
	bot.Handle(tele.OnText, myBot.TextHandler)
	bot.Start()
}
