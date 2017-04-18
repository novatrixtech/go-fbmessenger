package fbmodelsend

/*
Button - Button to be used as a reply option within a FB Message
*/
type Button struct {
	ButtonType string `json:"type,omitempty"`
	Title      string `json:"title,omitempty"`
	Payload    string `json:"payload,omitempty"`
	URL        string `json:"url,omitempty"`
}

/*
ButtonSharedContent - Button to be used as a reply option within a FB Share Content
*/
type ButtonSharedContent struct {
	ButtonType    string        `json:"type,omitempty"`
	Title         string        `json:"title,omitempty"`
	Payload       string        `json:"payload,omitempty"`
	URL           string        `json:"url,omitempty"`
	ShareContents ShareContents `json:"share_contents,omitempty"`
}
