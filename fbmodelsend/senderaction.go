package fbmodelsend

/*
SenderAction is a struct that represents message states typing_on, typing_off, mark_seen
More details at https://developers.facebook.com/docs/messenger-platform/send-api-reference
*/
type SenderAction struct {
	MessageType       string    `json:"message_type"`
	Recipient         Recipient `json:"recipient"`
	SenderActionState string    `json:"sender_action"`
}
