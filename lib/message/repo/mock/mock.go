package mock

import (
	"github.com/tylerlofgren/noiseaware/lib/message"
	"github.com/tylerlofgren/noiseaware/lib/message/repo"
)

func New() repo.Repo {
	return new(mockRepo)
}

type mockRepo struct {}

func (mr mockRepo) Save(m message.Message) (message.Message, error) {
	return message.Message{
		Timestamp: 1,
		Symbol: "aaa",
		Volume: 10,
		Temperature: 10,
	}, nil
}

func (mr mockRepo) Get() (message.Message, error) {
	return message.Message{
		Timestamp: 1,
		Symbol: "aaa",
		Volume: 10,
		Temperature: 10,
	}, nil
}
