package main

import (
	"encoding/json"
	"github.com/gravitymir/zadarma-golang/zadarma"
	"time"
)

//CatchStatisticsABX https://zadarma.com/ru/support/api/#api_statistics_pbx
type CatchStatisticsABX struct {
	Status string `json:"status"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Stats  []struct {
		CallId      string `json:"call_id"`
		Sip         string `json:"sip"`
		Callstart   string `json:"callstart"`
		Clid        string `json:"clid"`
		Destination int    `json:"destination"`
		Disposition string `json:"disposition"`
		Seconds     int    `json:"seconds"`
		IsRecorded  string `json:"is_recorded"`
		PbxCallId   string `json:"pbx_call_id"`
	}
	Message string `json:"message"`
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

func getCallUrl() (string, error) {
	return "", nil
}
