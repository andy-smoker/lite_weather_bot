package main

import (
	"encoding/json"
)

type bot struct {
	Token     string `toml:"token"`
	WeaterKey string `toml:"weather-api-key"`
	Debug     bool   `toml:"debug"`
	Buffer    int    `toml:"buffer"`

	shutdownClient chan interface{}
	client         httpClient `toml:"-"`
	Endpoint       string
}

// Param ..
type Param struct {
	Method string
	URL    string
}

// User ,
type User struct {
	ID            int    `json:"id"`
	IsBot         bool   `json:"is_bool"`
	Fname         string `json:"first_name"`
	Lname         string `json:"last_name"`
	UserName      string `json:"username"`
	LangCode      string `json:"language_code"`
	CanJoinGroup  bool   `json:"can_join_group"`
	CanReadAll    bool   `json:"can_read_all_group_messages"`
	SupportInline bool   `json:"supports_inline_queries"`
}

// ResponseParameters ..
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	RetryAfter      int   `json:"retry_after"`
}

type apiResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// Chat .
type Chat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type update struct {
	ID              int      `json:"update_id"`
	OK              bool     `json:"OK"`
	Message         *Message `json:"message"`
	EditMessage     *Message `json:"edit_message"`
	ChannelPost     *Message `json:"channel_post"`
	EditChannelPost *Message `json:"edit_channel_post"`
	//Poll		*Poll	`json:"pol"`

}

type updateChan <-chan update

type updateCfg struct {
	Offset         int      `json:"offset"`
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}
