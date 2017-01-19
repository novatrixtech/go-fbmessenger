package fbmodelsend

/*
Attachment - Represents a Facebook Message's Attachment
*/
type Attachment struct {
	AttachmentType string         `json:"type,omitempty"`
	Payload        MessagePayload `json:"payload,omitempty"`
}
