package telegram

type Update struct {
	ID      string `json:"update_id"`
	Message string `json:"message"`
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
