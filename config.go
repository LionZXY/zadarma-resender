package main

type config struct {
	TelegramToken string `env:"TG_TOKEN,required"`
	ChannelID     int64  `env:"CHANNEL_ID,required"`

	ZadarmaUserKey   string `env:"ZADARMA_USER_KEY, required"`
	ZadarmaSecretKey string `env:"ZADARMA_SECRET_KEY, required"`
}
