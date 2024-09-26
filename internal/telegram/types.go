package telegram

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID int `json:"update_id"`
	//Тут начинаются приколы с парсингом json. Не понимаю как в одно поле структуры запихнуть два тега json, ни в какую
	Message     *Message `json:"message,omitempty"`
	ChannelPost *Message `json:"channel_post,omitempty"`
}

type Message struct {
	ID   int64  `json:"message_id"`
	Text string `json:"text"`
	From From   `json:"from,omitempty"`
	Chat Chat   `json:"chat,omitempty"`
}

type From struct {
	Username string `json:"username,omitempty"`
}

type Chat struct {
	ID int64 `json:"id"`
}
