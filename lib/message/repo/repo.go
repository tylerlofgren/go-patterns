package repo

import "github.com/tylerlofgren/noiseaware/lib/message"

type Repo interface {
	Save(m message.Message) (message.Message, error)
	Get() (message.Message, error)
}
