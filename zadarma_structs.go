package main

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
