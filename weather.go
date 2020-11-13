package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Weather базовая струтура ответа погодного API
type Weather struct {
	Time string `json:"now_dt"`
	Fact Fact   `json:"fact"`
}

// Fact структура погодных показаний
type Fact struct {
	Temp      int     `json:"temp"`
	FeelsLike int     `json:"feels_like"`
	Condition string  `json:"condition"`
	WindSpeed float32 `json:"wind_speed"`
}

func (b *bot) WeatherRequest(coord string) (Weather, error) {
	apiResp := Weather{}
	URL := fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?%s", coord)
	req, err := http.NewRequest("GET", URL, strings.NewReader(""))
	if err != nil {
		return apiResp, err
	}
	req.Header.Set("X-Yandex-API-Key", b.WeaterKey)

	resp, err := b.client.Do(req)
	if err != nil {
		return apiResp, err
	}
	defer resp.Body.Close()

	err = decodeResp(resp.Body, &apiResp)
	if err != nil {
		return apiResp, err
	}
	return apiResp, nil
}

func (b *bot) WhatWeater(city string) string {
	var coord string
	switch city {
	case "VL":
		coord = "lat=43.7&lon=131.54"
		break
	default:
		coord = "lat=43.7&lon=131.54"
		break
	}
	w, err := b.WeatherRequest(coord)
	if err != nil {
		return "ERROR"
	}
	return fmt.Sprintf("На улице %s, температура: %d(ощущается как %d), ветер %Fм/с",
		ConditionRU(w.Fact.Condition), w.Fact.Temp, w.Fact.FeelsLike, w.Fact.WindSpeed)
}

// ConditionRU - перевод на русский соглласно документации
func ConditionRU(v string) string {
	switch v {
	case "clear":
		return "ясно"
	case "partly-cloudy":
		return " малооблачно"
	case "cloudy":
		return "облачно с прояснениями"
	case "overcast":
		return "пасмурно"
	case "drizzle":
		return "морось"
	case "light-rain":
		return "небольшой дождь"
	case "rain":
		return "дождь"
	case "moderate-rain":
		return "умеренно сильный дождь"
	case "heavy-rain":
		return "сильный дождь"
	case "continuous-heavy-rain":
		return "длительный сильный дождь"
	case "showers":
		return "ливень"
	case "wet-snow":
		return "дождь со снегом"
	case "light-snow":
		return "небольшой снег"
	case "snow":
		return "снег"
	case "snow-showers":
		return "снегопад"
	case "hail":
		return "град"
	case "thunderstorm":
		return "гроза"
	case "thunderstorm-with-rain":
		return "дождь с грозой"
	case "thunderstorm-with-hail":
		return "гроза с градом"
	default:
		return "... сам посмотри"
	}
}
