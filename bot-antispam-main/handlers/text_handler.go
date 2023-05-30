package handlers

import (
	"log"

	"diplomka/usecase"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Bot       *tele.Bot
	warnCount map[string]int
}

func NewBot(bot *tele.Bot, mapa map[string]int) *Bot {
	return &Bot{
		Bot:       bot,
		warnCount: mapa,
	}
}

func (bot *Bot) TextHandler(ctx tele.Context) error {
	_, err := bot.isAdmin(ctx.Chat(), ctx.Sender())
	if err != nil {
		return err
	}

	foundLink, flag := usecase.CheckLink(ctx.Message().Text)
	if flag {
		log.Println(foundLink)
		negative := usecase.VirusTotal(foundLink)
		if negative {
			ctx.Send("Хабарлама ішіндегі ссылка қауіпті, хабарлама өшірілді, ссылка: " + foundLink)
			bot.Bot.Delete(ctx.Message())
		}
		return nil
	}
	res, err := usecase.SpamWords(ctx.Message().Text)
	if err != nil {
		return err
	}
	if res {
		ctx.Send("Message contains a spam. Author is " + ctx.Sender().Username)
		bot.warnCount[ctx.Sender().Username]++
		bot.Bot.Delete(ctx.Message())
		if bot.warnCount[ctx.Sender().Username] == 3 {
			user := &tele.ChatMember{User: ctx.Sender()}
			bot.Bot.Ban(ctx.Chat(), user, false)
			ctx.Send("Пользователь " + user.User.Username + " был кикнут из чата.")
			delete(bot.warnCount, ctx.Sender().Username)
		}
		return nil
	}

	return nil
}

func (bot *Bot) isAdmin(chat *tele.Chat, user *tele.User) (bool, error) {
	member, err := bot.Bot.ChatMemberOf(chat, user)
	if err != nil {
		return false, err
	}
	return member.Role == tele.Administrator || member.Role == tele.Creator, nil
}
