package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
	"log"
)

var cfg config
var r *redis.Client

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	r = redis.NewClient(&redis.Options{
		Addr: cfg.DBAddr,
	})

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
			log.Fatalln("Error while convert file to ogg", err)
			return
		}
	}
}

func prettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s\n", string(dataJSON))
}
