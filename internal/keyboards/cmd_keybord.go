package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdKeyboard() tgbotapi.ReplyKeyboardMarkup {
	var cmdKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/take"),
			tgbotapi.NewKeyboardButton("/return")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("ОТПРАВЛЕНИЕ НОМЕРА ТЕЛЕФОНА")))
	return cmdKeyboard
}
