package telegram

import (
	"github.com/s1ovac/todobot/pkg/clients/telegram"
	"github.com/s1ovac/todobot/pkg/storage"
)

type Dispatcher struct {
	tg      telegram.Client
	offset  int
	storage storage.Storage
}

func NewDispatcher(tg telegram.Client, s storage.Storage) *Dispatcher {
	return &Dispatcher{
		tg:      tg,
		storage: s,
	}
}

func (d *Dispatcher) Process() error {
	return nil
}
