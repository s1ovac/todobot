package main

import (
	"errors"
	"flag"
	"github.com/s1ovac/todobot/pkg/clients/telegram"
	"log"
)

func main() {
	tgClient := telegram.NewClient("api.telegram.org", mustToken())
	// fetcher := fetcher.New() - отправляет запросы, чтобы получать события
	// processor := processor.New() - обрабатывает запросы

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token", "", "token to connect to TelegramApi")
	flag.Parse()
	if len(*token) < 1 {
		log.Fatal(errors.New("telegram token wasn't entered"))
	}
	return *token
}
