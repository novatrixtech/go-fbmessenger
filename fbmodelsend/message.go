package fbmodelsend

/*
Message - Represents a Facebook's Message
*/
type Message struct {
	Text         string        `json:"text,omitempty"`
	Attachment   *Attachment   `json:"attachment,omitempty"`
	QuickReplies []*QuickReply `json:"quick_replies,omitempty"`
}
