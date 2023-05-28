package bot

import tele "gopkg.in/telebot.v3"

func InitBot(token string) (*tele.Bot, error) {
	bot, err := tele.NewBot(tele.Settings{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return bot, nil
}
