package merchant_bot

import (
  "github.com/go-telegram/bot"
)

var b *bot.Bot

func SetBot(nextBot *bot.Bot) {
  b = nextBot
}

func Instance() *bot.Bot {
  return b
}