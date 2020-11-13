package main

// Message структура входящих сообщений
type Message struct {
	ID          int   `json:"message_id"`
	From        *User `json:"from"`
	Date        int   `json:"date"`
	Chat        *Chat `json:"chat"`
	ForwardFrom *User `json:"forvard_from"`
	//ForwardFromChan  *Chat            `json:"forward_from_chat"`
	ForwardFromMsgID int      `json:"forward_from_message_id"`
	ReplyToMessage   *Message `json:"reply_to_message"`
	ViaBot           *User    `json:"via_bot"`
	EditDate         int      `json:"editdate"`
	MediaGroupID     string   `json:"madia_group_id"`
	AuthorSignature  string   `jsdon:"author_signature"`
	Text             string   `json:"text"`
	Caption          string   `json:"caption"`
	//CaptionEntities  *[]MessageEntity `json:"caption_entities"`
	//Contact          *Contact         `json:"contact"`
	//Location       Location     `json:"location"`
	NewChatMembers *[]User     `json:"new_chat_members"`
	LeftChatMember *User       `json:"left_chat_member"`
	NewChatTitle   string      `json:"new_chat_title"`
	DelChatPhoto   bool        `json:"delete_chat_photo"`
	ReplyMarkup    interface{} `json:"reply_markup"`
}

// SendMessage структура исходящих сообщений
type SendMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

// NewMessage создаём новое сообщение
func NewMessage(chat int, text string) *SendMessage {
	return &SendMessage{
		ChatID: chat,
		Text:   text,
	}
}
