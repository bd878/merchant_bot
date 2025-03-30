package merchant_bot

type Monolith interface {
	Repo() *Repo
	Bot() *Bot
	Log() *Logger
	Config() Config
	Chats() *Chats
}
