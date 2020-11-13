package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	bot := newBot(&http.Client{})
	upCfg := updateCfg{
		Offset:  (-1),
		Timeout: 10,
	}

	startTime := time.Now()
	weather := bot.WhatWeater("VL")
	lastMsg := &Message{}
	for {
		updt, err := bot.getUpdates(upCfg)
		if err != nil {
			log.Println(err)
		}
		// обноление погоды раз в час после запуска
		if time.Now().After(startTime.Add(time.Hour * 1)) {
			weather = bot.WhatWeater("VL")
			startTime = time.Now()
		}
		for _, val := range updt {
			msg := NewMessage(val.Message.Chat.ID, "")
			if val.Message.Text == "/погода" && val.Message.Date != lastMsg.Date {
				msg.Text = weather
				fmt.Println(&msg)
				req, err := json.Marshal(&msg)
				if err != nil {
					log.Panic()
				}
				bot.MakeJSONRequest("sendMessage", req)
				lastMsg = val.Message
			}
		}
	}
}
