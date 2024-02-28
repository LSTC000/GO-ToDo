package model

type Message struct {
	Message string `json:"message"`
}

type HTTPError struct {
	Code        int    `json:"code"`
	ServiceCode int    `json:"serviceCode"`
	Detail      string `json:"detail"`
}
