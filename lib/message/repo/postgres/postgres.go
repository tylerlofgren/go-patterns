package postgres

import (
	"github.com/tylerlofgren/noiseaware/lib/message"
	"github.com/tylerlofgren/noiseaware/lib/message/repo"
)

func New() repo.Repo {
	return new(messageRepo)
}

type messageRepo struct {}

func (mr messageRepo) Save(m message.Message) (message.Message, error) {
	panic("implement me")
}

func (mr messageRepo) Get() (message.Message, error) {
	panic("implement me")
}
