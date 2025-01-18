package telegram

type TGMessage struct {
	ClientID int     `json:"@client_id"`
	Extra    string  `json:"@extra"`
	Type     string  `json:"@type"`
	Message  Message `json:"message"`
}

type Caption struct {
	Text string `json:"text"`
}

type PhotoItem struct {
	ID int `json:"id"`
}

type Sizes struct {
	Photo PhotoItem `json:"photo"`
}

type Photo struct {
	ClientID    int     `json:"@client_id"`
	Extra       string  `json:"@extra"`
	Type        string  `json:"@type"`
	HasStickers bool    `json:"has_stickers"`
	Sizes       []Sizes `json:"sizes"`
}

type Content struct {
	ClientID   int     `json:"@client_id"`
	Extra      string  `json:"@extra"`
	Type       string  `json:"@type"`
	Caption    Caption `json:"caption"`
	HasSpoiler bool    `json:"has_spoiler"`
	IsSecret   bool    `json:"is_secret"`
	Photo      Photo   `json:"photo"`
}

type Message struct {
	ID           int64    `json:"id"`
	ChatID       int64    `json:"chat_id"`
	Content      Content  `json:"content"`
	MediaAlbumID string   `json:"media_album_id"`
	Date         int64    `json:"date"`
	SenderID     SenderID `json:"sender_id"`
}

type SenderID struct {
	UserID   int64  `json:"user_id"`
	ClientID int    `json:"@client_id"`
	Extra    string `json:"@extra"`
	Type     string `json:"@type"`
}
