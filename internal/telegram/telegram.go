package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"juniorParseBot/model"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod     = "getUpdates"
	sendMessageMethod    = "sendMessage"
	forwardMessageMethod = "forwardMessage"
)

type Bot struct {
	BotUrl   string
	BasePath string
	client   http.Client
}

func New(cfg *model.Config) *Bot {

	return &Bot{
		BotUrl:   cfg.BotUrl,
		BasePath: newBasePath(cfg.Token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (b *Bot) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := b.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (b *Bot) SendMessage(chatId int64, text string) error {
	q := url.Values{}

	q.Add("chat_id", strconv.FormatInt(chatId, 10))
	q.Add("text", text)

	_, err := b.doRequest(sendMessageMethod, q)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) ForwardMessage(chatId int64, fromChatId string, messageId int64) error {
	q := url.Values{}

	q.Add("chat_id", fromChatId)
	q.Add("from_chat_id", strconv.FormatInt(chatId, 10))
	q.Add("message_id", strconv.FormatInt(messageId, 10))

	_, err := b.doRequest(forwardMessageMethod, q)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme:   "https",
		Host:     b.BotUrl,
		Path:     path.Join(b.BasePath, method),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()
	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	return body, nil
}
