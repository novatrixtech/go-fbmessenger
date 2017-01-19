package fblib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/novatrixtech/go-fbmessenger/fbmodelsend"
)

var logLevelDebug = false

/*
SendTextMessage - Send text message to a recipient on Facebook Messenger
*/
func SendTextMessage(text string, recipient string) {
	letter := new(fbmodelsend.Letter)
	letter.Message.Text = text
	letter.Recipient.ID = recipient

	if err := sendMessage(letter, recipient); err != nil {
		fmt.Print("[sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendImageMessage - Send image message to a recipient on Facebook Messenger
*/
func SendImageMessage(url string, recipient string) {
	message := new(fbmodelsend.Letter)

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "image"
	attch.Payload.URL = url
	message.Message.Attachment = attch

	message.Recipient.ID = recipient

	if err := sendMessage(message, recipient); err != nil {
		fmt.Print("[sendImageMessage] Error during the call to Facebook to send the image message: " + err.Error())
		return
	}
}

/*
SendTypingMessage - Send typing message to user
*/
func SendTypingMessage(onoff bool, recipient string) {
	senderAction := new(fbmodelsend.SenderAction)
	senderAction.Recipient.ID = recipient
	if onoff {
		senderAction.SenderActionState = "typing_on"
	} else {
		senderAction.SenderActionState = "typing_off"
	}
	if err := sendMessage(senderAction, recipient); err != nil {
		fmt.Print("[sendImageMessage] Error during the call to Facebook to send the typing message: " + err.Error())
		return
	}
}

/*
SendGenericTemplateMessage - Sends a generic rich message to Facebook user.
It can include text, buttons, URLs Butttons, lists to reply
*/
func SendGenericTemplateMessage(template []*fbmodelsend.TemplateElement, recipient string) {
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "template"
	attch.Payload.TemplateType = "generic"
	attch.Payload.Elements = template

	msg.Message.Attachment = attch

	if err := sendMessage(msg, recipient); err != nil {
		fmt.Print("[SendGenericTemplateMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendButtonMessage - Sends a generic rich message to Facebook user.
It can include text, buttons, URLs Butttons, lists to reply
*/
func SendButtonMessage(template []*fbmodelsend.Button, text string, recipient string) {
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "template"
	attch.Payload.TemplateType = "button"
	attch.Payload.Text = text
	attch.Payload.Buttons = template

	msg.Message.Attachment = attch

	if err := sendMessage(msg, recipient); err != nil {
		fmt.Print("[sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendURLButtonMessage - Sends a message with a button that redirects the user to an external web page.
*/
func SendURLButtonMessage(text string, buttonTitle string, URL string, recipient string) {

	msgElement := new(fbmodelsend.TemplateElement)
	msgElement.Title = text

	opt1 := new(fbmodelsend.Button)
	opt1.ButtonType = "web_url"
	opt1.Title = buttonTitle
	opt1.URL = URL

	buttons := []*fbmodelsend.Button{opt1}

	msgElement.Buttons = buttons
	elements := []*fbmodelsend.TemplateElement{msgElement}
	SendGenericTemplateMessage(elements, recipient)
}

/*
SendShareMessage sends the message along with Share Button
*/
func SendShareMessage(text string, subtitle string, recipient string) {

	msgElement := new(fbmodelsend.TemplateElement)
	msgElement.Title = text
	msgElement.Subtitle = subtitle

	opt1 := new(fbmodelsend.Button)
	opt1.ButtonType = "element_share"
	buttons := []*fbmodelsend.Button{opt1}

	msgElement.Buttons = buttons
	elements := []*fbmodelsend.TemplateElement{msgElement}
	SendGenericTemplateMessage(elements, recipient)

}

/*
Send Message - Sends a generic message to Facebook Messenger
*/
func sendMessage(message interface{}, recipient string) error {

	if logLevelDebug {
		scs := spew.ConfigState{Indent: "\t"}
		scs.Dump(message)
		return nil
	}

	url := "https://graph.facebook.com/v2.8/me/messages?access_token=" + fbAccessToken

	data, err := json.Marshal(message)
	if err != nil {
		fmt.Print("[sendMessage] Error to convert message object: " + err.Error())
		return err
	}

	reqFb, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	reqFb.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	//fmt.Println("[sendMessage] Replying at: " + url + " the message " + string(data))

	respFb, err := client.Do(reqFb)
	if err != nil {
		fmt.Print("[sendMessage] Error during the call to Facebook to send the message: " + err.Error())
		return err
	}
	defer respFb.Body.Close()

	//fmt.Println("[sendMessage] Response status: ", respFb.Status)
	//bodyFromFb, _ := ioutil.ReadAll(respFb.Body)
	//fmt.Println("[sendMessage] Response Body from Facebook: ", string(bodyFromFb))

	return nil
}
