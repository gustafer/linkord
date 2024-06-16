package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/gustafer/linkord/configs"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	key                 = os.Getenv("JWT_SECRET")
	discordClientId     = os.Getenv("DISCORD_CLIENT_ID")
	discordClientSecret = os.Getenv("DISCORD_CLIENT_SECRET")
	MaxAge              = 86400 * 30
	isProd              = false
)

func NewAuth() {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	port := configs.LoadPort()
	protocol := configs.GetProtocol()

	goth.UseProviders(
		discord.New(discordClientId, discordClientSecret, fmt.Sprintf("%v%v/auth/discord/callback", protocol, port),
			discord.ScopeIdentify, discord.ScopeEmail),
	)
}
