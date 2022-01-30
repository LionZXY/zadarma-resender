package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gravitymir/zadarma-golang/zadarma"
	"log"
	"time"
)

var cfg config

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	statistics, err := getCalls()

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	prettyPrint(statistics)
}

func getCalls() (CatchStatisticsABX, error) {
	formatString := "2006-01-02 03:04:05"
	now := time.Now()
	dayBefore := now.AddDate(0, -1, 0)

	statisticRequest := zadarma.New{
		APIMethod:    "/v1/statistics/pbx/",
		APIUserKey:   cfg.ZadarmaUserKey,
		APISecretKey: cfg.ZadarmaSecretKey,
		ParamsMap: map[string]string{
			"start":   dayBefore.Format(formatString),
			"end":     now.Format(formatString),
			"version": "2",
		},
	}

	var statisticResponse []byte

	if err := statisticRequest.Request(&statisticResponse); err != nil {
		return CatchStatisticsABX{}, err
	}

	statistics := CatchStatisticsABX{}
	if err := json.Unmarshal(statisticResponse, &statistics); err != nil {
		return statistics, err
	}

	return statistics, nil
}

func prettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s\n", string(dataJSON))
}
