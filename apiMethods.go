package main

import (
	"encoding/json"
)

func (b *bot) getUpdates(cfg updateCfg) ([]update, error) {
	u := []update{}
	resp, err := b.MakeRequest("getUpdates", cfg.value())
	if resp.OK != true {
		return u, err
	}

	err = json.Unmarshal(resp.Result, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}
