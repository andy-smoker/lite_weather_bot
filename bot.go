package main

import (
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
)

const endpoint = "https://api.telegram.org/bot%s/%s"

func newBot(client httpClient) *bot {

	nBot := &bot{
		client:         client,
		shutdownClient: make(chan interface{}),
		Endpoint:       endpoint,
	}
	err := nBot.config()
	if err != nil {

		return nBot
	}
	fmt.Println(nBot)
	return nBot
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (b *bot) config() error {
	_, err := toml.DecodeFile("config.toml", &b)
	if err != nil {
		return err
	}
	return nil
}
