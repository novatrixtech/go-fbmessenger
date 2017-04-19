package fbmodelsend

/*
Attachment - Represents a Facebook Message's Attachment
*/
type Attachment struct {
	AttachmentType string         `json:"type,omitempty"`
	Payload        MessagePayload `json:"payload,omitempty"`
}

/*
SharedAttachment - Represents a Facebook Shared Content's Attachment
*/
type SharedAttachment struct {
	AttachmentType string        `json:"type,omitempty"`
	Payload        SharedPayload `json:"payload,omitempty"`
}
