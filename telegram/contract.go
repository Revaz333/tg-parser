package telegram

type TGMessage struct {
	ClientID int     `json:"@client_id"`
	Extra    string  `json:"@extra"`
	Type     string  `json:"@type"`
	Message  Message `json:"message"`
}
type Text struct {
	ClientID int    `json:"@client_id"`
	Extra    string `json:"@extra"`
	Type     string `json:"@type"`
	Entities []any  `json:"entities"`
	Text     string `json:"text"`
}
type Content struct {
	ClientID           int    `json:"@client_id"`
	Extra              string `json:"@extra"`
	Type               string `json:"@type"`
	LinkPreviewOptions any    `json:"link_preview_options"`
	Text               Text   `json:"text"`
	WebPage            any    `json:"web_page"`
}
type InteractionInfo struct {
	ClientID     int    `json:"@client_id"`
	Extra        string `json:"@extra"`
	Type         string `json:"@type"`
	ForwardCount int    `json:"forward_count"`
	Reactions    any    `json:"reactions"`
	ReplyInfo    any    `json:"reply_info"`
	ViewCount    int    `json:"view_count"`
}
type SenderID struct {
	ClientID int    `json:"@client_id"`
	Extra    string `json:"@extra"`
	Type     string `json:"@type"`
	ChatID   int64  `json:"chat_id"`
}
type Message struct {
	ClientID                  int             `json:"@client_id"`
	Extra                     string          `json:"@extra"`
	Type                      string          `json:"@type"`
	AuthorSignature           string          `json:"author_signature"`
	AutoDeleteIn              int             `json:"auto_delete_in"`
	CanBeDeletedForAllUsers   bool            `json:"can_be_deleted_for_all_users"`
	CanBeDeletedOnlyForSelf   bool            `json:"can_be_deleted_only_for_self"`
	CanBeEdited               bool            `json:"can_be_edited"`
	CanBeForwarded            bool            `json:"can_be_forwarded"`
	CanBeRepliedInAnotherChat bool            `json:"can_be_replied_in_another_chat"`
	CanBeSaved                bool            `json:"can_be_saved"`
	CanGetAddedReactions      bool            `json:"can_get_added_reactions"`
	CanGetMediaTimestampLinks bool            `json:"can_get_media_timestamp_links"`
	CanGetMessageThread       bool            `json:"can_get_message_thread"`
	CanGetReadDate            bool            `json:"can_get_read_date"`
	CanGetStatistics          bool            `json:"can_get_statistics"`
	CanGetViewers             bool            `json:"can_get_viewers"`
	CanReportReactions        bool            `json:"can_report_reactions"`
	ChatID                    int64           `json:"chat_id"`
	ContainsUnreadMention     bool            `json:"contains_unread_mention"`
	Content                   Content         `json:"content"`
	Date                      int             `json:"date"`
	EditDate                  int             `json:"edit_date"`
	ForwardInfo               any             `json:"forward_info"`
	HasTimestampedMedia       bool            `json:"has_timestamped_media"`
	ID                        int             `json:"id"`
	ImportInfo                any             `json:"import_info"`
	InteractionInfo           InteractionInfo `json:"interaction_info"`
	IsChannelPost             bool            `json:"is_channel_post"`
	IsFromOffline             bool            `json:"is_from_offline"`
	IsOutgoing                bool            `json:"is_outgoing"`
	IsPinned                  bool            `json:"is_pinned"`
	IsTopicMessage            bool            `json:"is_topic_message"`
	MediaAlbumID              string          `json:"media_album_id"`
	MessageThreadID           int             `json:"message_thread_id"`
	ReplyMarkup               any             `json:"reply_markup"`
	ReplyTo                   any             `json:"reply_to"`
	RestrictionReason         string          `json:"restriction_reason"`
	SavedMessagesTopicID      int             `json:"saved_messages_topic_id"`
	SchedulingState           any             `json:"scheduling_state"`
	SelfDestructIn            int             `json:"self_destruct_in"`
	SelfDestructType          any             `json:"self_destruct_type"`
	SenderBoostCount          int             `json:"sender_boost_count"`
	SenderBusinessBotUserID   int             `json:"sender_business_bot_user_id"`
	SenderID                  SenderID        `json:"sender_id"`
	SendingState              any             `json:"sending_state"`
	UnreadReactions           []any           `json:"unread_reactions"`
	ViaBotUserID              int             `json:"via_bot_user_id"`
}
