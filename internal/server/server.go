package server

import (
	"juniorParseBot/internal/parser"
	"juniorParseBot/internal/telegram"
	"juniorParseBot/model"
	"log"
	"time"
)

type Server struct {
	bot         *telegram.Bot
	updatesSize int
	offset      int
}

type MessageToForward struct {
	Text      string
	MessageID int64
	ChatID    int64
}

func New(b *telegram.Bot, us int) *Server {
	return &Server{
		bot:         b,
		updatesSize: us,
	}
}

func (s *Server) Start(cfg *model.Config) error {
	for {
		gotUpdates, err := s.getUpdates()
		if err != nil {
			log.Println(err)
		}

		if len(gotUpdates) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		err = s.ProceedUpdates(gotUpdates, cfg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) ProceedUpdates(updates []telegram.Update, cfg *model.Config) error {
	for _, update := range updates {

		ForwardData := &MessageToForward{}

		//Да, дальше идут небольшие костыли, так как я не знаю как правильнее организовать структуру telegram.Update
		if update.Message != nil {

			ForwardData.Text = update.Message.Text
			ForwardData.MessageID = update.Message.ID
			ForwardData.ChatID = update.Message.Chat.ID
		}
		if update.ChannelPost != nil {
			ForwardData.Text = update.ChannelPost.Text
			ForwardData.MessageID = update.ChannelPost.ID
			ForwardData.ChatID = update.ChannelPost.Chat.ID
		}
		isContains := parser.FindKeyword(ForwardData.Text, cfg.Keywords)
		log.Println(isContains, ForwardData.Text)
		if isContains {
			for _, receiver := range cfg.ForwardTo {
				if err := s.bot.ForwardMessage(ForwardData.ChatID, receiver, ForwardData.MessageID); err != nil {
					log.Println(err)
					return err
				}
			}
		}
	}
	return nil
}

func (s *Server) getUpdates() ([]telegram.Update, error) {
	updates, err := s.bot.Updates(s.offset, s.updatesSize)
	if err != nil {
		return nil, err
	}
	if len(updates) == 0 {
		return nil, nil
	}
	s.offset = updates[len(updates)-1].ID + 1

	return updates, nil
}
