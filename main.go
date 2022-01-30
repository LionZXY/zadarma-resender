package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var cfg config
var r *redis.Client
var tg *tb.Bot

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	r = redis.NewClient(&redis.Options{
		Addr: cfg.DBAddr,
	})

	var err error

	tg, err = tb.NewBot(tb.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalln("Telegram", err)
	}

	checkNewCalls()
}

func checkNewCalls() {
	statistics, err := getCalls()

	if err != nil {
		log.Println("Calls", err)
		return
	}
	prettyPrint(statistics)

	for _, call := range statistics.Stats {
		if call.IsRecorded != "true" {
			continue
		}
		if wasAlreadyPosted(call.CallId) {
			continue
		}
		log.Println("Posting...", call.CallId)
		recordUrl, err := getRecordUrl(call.CallId)
		if err != nil {
			log.Fatalln("Error get record url for", call.CallId, err)
			return
		}
		log.Println("Record url is", recordUrl)
		path, err := convertToOgg(recordUrl.Link)
		if err != nil {
			log.Println("Error while convert file to ogg", err, recordUrl.Link)
			return
		}
		_, err = tg.Send(tb.ChatID(cfg.ChannelID), &tb.Voice{
			File: tb.File{FileLocal: path},
		})

		if err != nil {
			log.Println("Send audio", err)
			return
		}
		markPosted(call.CallId)
		log.Println("Posted!")
	}
}

func prettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s\n", string(dataJSON))
}
