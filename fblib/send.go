package fblib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/novatrixtech/go-fbmessenger/fbmodelsend"
)

var logLevelDebug = false

/*
SendTextMessage - Send text message to a recipient on Facebook Messenger
*/
func SendTextMessage(text string, recipient string, accessToken string) {
	letter := new(fbmodelsend.Letter)
	letter.Message.Text = text
	letter.Recipient.ID = recipient

	if err := sendMessage(letter, recipient, accessToken); err != nil {
		fmt.Print("[sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendImageMessage - Sends image message to a recipient on Facebook Messenger
*/
func SendImageMessage(url string, recipient string, accessToken string) {
	message := new(fbmodelsend.Letter)

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "image"
	attch.Payload.URL = url
	message.Message.Attachment = attch

	message.Recipient.ID = recipient

	if err := sendMessage(message, recipient, accessToken); err != nil {
		fmt.Print("[sendImageMessage] Error during the call to Facebook to send the image message: " + err.Error())
		return
	}
}

/*
SendAudioMessage - Sends audio message to a recipient on Facebook Messenger
*/
func SendAudioMessage(url string, recipient string, accessToken string) {
	message := new(fbmodelsend.Letter)

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "audio"
	attch.Payload.URL = url
	message.Message.Attachment = attch

	message.Recipient.ID = recipient

	if err := sendMessage(message, recipient, accessToken); err != nil {
		fmt.Print("[sendImageMessage] Error during the call to Facebook to send the audio message: " + err.Error())
		return
	}
}

/*
SendTypingMessage - Sends typing message to user
*/
func SendTypingMessage(onoff bool, recipient string, accessToken string) {
	senderAction := new(fbmodelsend.SenderAction)
	senderAction.Recipient.ID = recipient
	if onoff {
		senderAction.SenderActionState = "typing_on"
	} else {
		senderAction.SenderActionState = "typing_off"
	}
	if err := sendMessage(senderAction, recipient, accessToken); err != nil {
		fmt.Print("[sendImageMessage] Error during the call to Facebook to send the typing message: " + err.Error())
		return
	}
}

/*
SendGenericTemplateMessage - Sends a generic rich message to Facebook user.
It can include text, buttons, URLs Butttons, lists to reply
*/
func SendGenericTemplateMessage(template []*fbmodelsend.TemplateElement, recipient string, accessToken string) {
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "template"
	attch.Payload.TemplateType = "generic"
	attch.Payload.Elements = template

	msg.Message.Attachment = attch

	if err := sendMessage(msg, recipient, accessToken); err != nil {
		fmt.Print("[SendGenericTemplateMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendButtonMessage - Sends a generic rich message to Facebook user.
It can include text, buttons, URLs Butttons, lists to reply
*/
func SendButtonMessage(template []*fbmodelsend.Button, text string, recipient string, accessToken string) {
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "template"
	attch.Payload.TemplateType = "button"
	attch.Payload.Text = text
	attch.Payload.Buttons = template

	msg.Message.Attachment = attch

	if err := sendMessage(msg, recipient, accessToken); err != nil {
		fmt.Print("[sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendURLButtonMessage - Sends a message with a button that redirects the user to an external web page.
*/
func SendURLButtonMessage(text string, buttonTitle string, URL string, recipient string, accessToken string) {

	msgElement := new(fbmodelsend.TemplateElement)
	msgElement.Title = text

	opt1 := new(fbmodelsend.Button)
	opt1.ButtonType = "web_url"
	opt1.Title = buttonTitle
	opt1.URL = URL

	buttons := []*fbmodelsend.Button{opt1}

	msgElement.Buttons = buttons
	elements := []*fbmodelsend.TemplateElement{msgElement}
	SendGenericTemplateMessage(elements, recipient, accessToken)
}

/*
SendShareMessage sends the message along with Share Button
*/
func SendShareMessage(title string, subtitle string, recipient string, accessToken string) {

	msgElement := new(fbmodelsend.TemplateElement)
	msgElement.Title = title
	msgElement.Subtitle = subtitle

	opt1 := new(fbmodelsend.Button)
	opt1.ButtonType = "element_share"
	buttons := []*fbmodelsend.Button{opt1}

	msgElement.Buttons = buttons
	elements := []*fbmodelsend.TemplateElement{msgElement}
	SendGenericTemplateMessage(elements, recipient, accessToken)

}

/*
SendShareMessage sends the message along with Share Button
*/
func SendShareContent(title string, subtitle string, buttonTitle string, imageURL string, destinationURL string, recipient string, accessToken string) {

	btnElementButton := new(fbmodelsend.Button)
	btnElementButton.ButtonType = "web_url"
	btnElementButton.URL = destinationURL
	btnElementButton.Title = buttonTitle
	buttonsElementButton := []*fbmodelsend.Button{btnElementButton}

	btnElement := new(fbmodelsend.TemplateElement)
	btnElement.Title = title
	btnElement.Subtitle = subtitle
	btnElement.ImageURL = imageURL
	btnElement.DefaultAction.Type = "web_url"
	btnElement.DefaultAction.URL = destinationURL
	btnElement.Buttons = buttonsElementButton
	elementsButtonElement := []*fbmodelsend.TemplateElement{btnElement}

	opt1 := new(fbmodelsend.Button)
	opt1.ButtonType = "element_share"
	opt1.ShareContents.Attachment.Type = "template"
	opt1.ShareContents.Attachment.Payload.TemplateType = "generic"
	opt1.ShareContents.Attachment.Payload.Elements = elementsButtonElement
	buttons := []*fbmodelsend.Button{opt1}

	msgElement := new(fbmodelsend.TemplateElement)
	msgElement.Title = title
	msgElement.Subtitle = subtitle
	msgElement.Buttons = buttons
	elements := []*fbmodelsend.TemplateElement{msgElement}
	SendGenericTemplateMessage(elements, recipient, accessToken)

}

/*
SendQuickReply sends small messages in order to get small and quick answers from the users
*/
func SendQuickReply(text string, options []*fbmodelsend.QuickReply, recipient string, accessToken string) {
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient
	msg.Message.Text = text
	msg.Message.QuickReplies = options
	//log.Printf("[SendQuickReply] Enviado: [%s]\n", text)
	if err := sendMessage(msg, recipient, accessToken); err != nil {
		log.Print("[SendQuickReply] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
}

/*
SendAskUserLocation sends small message asking the users their location
*/
func SendAskUserLocation(text string, recipient string, accessToken string) {
	qr := new(fbmodelsend.QuickReply)
	qr.ContentType = "location"

	arrayQr := []*fbmodelsend.QuickReply{qr}

	SendQuickReply(text, arrayQr, recipient, accessToken)
}

/*
Send Message - Sends a generic message to Facebook Messenger
*/
func sendMessage(message interface{}, recipient string, accessToken string) error {

	if logLevelDebug {
		scs := spew.ConfigState{Indent: "\t"}
		scs.Dump(message)
		return nil
	}

	var url string
	if strings.Contains(accessToken, "http") {
		url = accessToken
	} else {
		url = "https://graph.facebook.com/v2.8/me/messages?access_token=" + accessToken
	}

	data, err := json.Marshal(message)
	if err != nil {
		fmt.Print("[sendMessage] Error to convert message object: " + err.Error())
		return err
	}

	reqFb, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	reqFb.Header.Set("Content-Type", "application/json")
	reqFb.Header.Set("Connection", "close")
	reqFb.Close = true

	client := &http.Client{}

	//fmt.Println("[sendMessage] Replying at: " + url + " the message " + string(data))

	respFb, err := client.Do(reqFb)
	if err != nil {
		fmt.Print("[sendMessage] Error during the call to Facebook to send the message: " + err.Error())
		return err
	}
	defer respFb.Body.Close()

	if respFb.StatusCode < 200 || respFb.StatusCode >= 300 {
		bodyFromFb, _ := ioutil.ReadAll(respFb.Body)
		status := string(bodyFromFb)
		fmt.Printf("[sendMessage] Response status code: [%d]\n", respFb.StatusCode)
		fmt.Println("[sendMessage] Response status: ", respFb.Status)
		fmt.Println("[sendMessage] Response Body from Facebook: ", status)
		fmt.Printf("[sendMessage] Facebook URL Called: [%s]\n", url)
		fmt.Printf("[sendMessage] Object sent to Facebook: [%s]\n", string(data))
	}

	return nil
}
