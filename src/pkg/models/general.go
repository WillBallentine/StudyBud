package model

type NLPRequestData struct {
	Text string `json:"text"`
}

type NLPResponse struct {
	Entities []struct {
		Text  string `json:"text"`
		Label string `json:"label"`
	} `json:"entities"`
}
