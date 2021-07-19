package service

import (
	"github.com/tylerlofgren/noiseaware/lib/message"
	"github.com/tylerlofgren/noiseaware/lib/message/repo"
)

type Service interface {
	Save(m message.Message) (message.Message, error)
	Get() (message.Message, error)
}

func New(r repo.Repo) Service {
	ms := new(messageService)
	ms.repo = r
	return ms
}

type messageService struct {
	repo repo.Repo
}

func (ms messageService) Save(m message.Message) (message.Message, error) {
	return ms.repo.Save(m)
}

func (ms messageService) Get() (message.Message, error) {
	return ms.repo.Get()
}
