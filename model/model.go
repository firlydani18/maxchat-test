package model

type Item struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Model       string   `json:"model"`
	Tech        []string `json:"tech"`
	Status      string   `json:"status"`
}
