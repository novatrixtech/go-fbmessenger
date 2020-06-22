package fblib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/novatrixtech/go-fbmessenger/fbmodelsend"
)

var logLevelDebug = false

//ErrInvalidCallToFacebook is specific error when Facebook Messenger returns error after being called
var ErrInvalidCallToFacebook = errors.New("")

//MessageTypeResponse is in response to a received message.
const MessageTypeResponse = 1

//MessageTypeUpdate is being sent proactively and is not in response to a received message.
const MessageTypeUpdate = 2

//MessageTypeMessageTag is non-promotional and is being sent outside the 24-hour standard messaging window with a message tag.
const MessageTypeMessageTag = 3

//defineMessageType returns the Message Type description defined by Messenger
func defineMessageType(msgType int) (msgTypeDescription string) {
	switch msgType {
	case 2:
		msgTypeDescription = "UPDATE"
	case 3:
		msgTypeDescription = "MESSAGE_TAG"
	default:
		msgTypeDescription = "RESPONSE"
	}
	return
}

/*
SendTextMessage - Send text message to a recipient on Facebook Messenger
*/
func SendTextMessage(text string, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	letter := new(fbmodelsend.Letter)
	letter.Message.Text = text
	letter.Recipient.ID = recipient
	letter.MessageType = defineMessageType(msgType)
	err = sendMessage(letter, recipient, accessToken)
	if err != nil {
		//fmt.Print("[fblib][sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
}

//SendPersonalFinanceUpdateMessage sends a Finance Update information to recipient
func SendPersonalFinanceUpdateMessage(text string, recipient string, accessToken string) (err error) {
	err = nil
	letter := new(fbmodelsend.Letter)
	letter.Message.Text = text
	letter.Tag = "PERSONAL_FINANCE_UPDATE"
	letter.Recipient.ID = recipient
	letter.MessageType = defineMessageType(3)
	err = sendMessage(letter, recipient, accessToken)
	if err != nil {
		//fmt.Print("[fblib][sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
}

/*
SendImageMessage - Sends image message to a recipient on Facebook Messenger
*/
func SendImageMessage(url string, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	message := new(fbmodelsend.Letter)
	message.MessageType = defineMessageType(msgType)

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "image"
	attch.Payload.URL = url
	message.Message.Attachment = attch

	message.Recipient.ID = recipient
	err = sendMessage(message, recipient, accessToken)
	if err != nil {
		fmt.Print("[fblib][sendImageMessage] Error during the call to Facebook to send the image message: " + err.Error())
		return
	}
	return
}

/*
SendAudioMessage - Sends audio message to a recipient on Facebook Messenger
*/
func SendAudioMessage(url string, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	message := new(fbmodelsend.Letter)
	message.MessageType = defineMessageType(msgType)
	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "audio"
	attch.Payload.URL = url
	message.Message.Attachment = attch

	message.Recipient.ID = recipient
	err = sendMessage(message, recipient, accessToken)
	if err != nil {
		//fmt.Print("[fblib][sendImageMessage] Error during the call to Facebook to send the audio message: " + err.Error())
		return
	}
	return
}

/*
SendTypingMessage - Sends typing message to user
*/
func SendTypingMessage(onoff bool, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	senderAction := new(fbmodelsend.SenderAction)
	senderAction.MessageType = defineMessageType(msgType)
	senderAction.Recipient.ID = recipient
	if onoff {
		senderAction.SenderActionState = "typing_on"
	} else {
		senderAction.SenderActionState = "typing_off"
	}
	err = sendMessage(senderAction, recipient, accessToken)
	if err != nil {
		//fmt.Print("[fblib][sendImageMessage] Error during the call to Facebook to send the typing message: " + err.Error())
		return
	}
	return
}

/*
SendGenericTemplateMessage - Sends a generic rich message to Facebook user.
It can include text, buttons, URLs Butttons, lists to reply
*/
func SendGenericTemplateMessage(template []*fbmodelsend.TemplateElement, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient
	msg.MessageType = defineMessageType(msgType)
	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "template"
	attch.Payload.TemplateType = "generic"
	attch.Payload.Elements = template

	msg.Message.Attachment = attch

	err = sendMessage(msg, recipient, accessToken)
	if err != nil {
		//fmt.Print("[fblib][SendGenericTemplateMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
}

/*
SendButtonMessage - Sends a generic rich message to Facebook user.
It can include text, buttons, URLs Butttons, lists to reply
*/
func SendButtonMessage(template []*fbmodelsend.Button, text string, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	msg := new(fbmodelsend.Letter)
	msg.Recipient.ID = recipient
	msg.MessageType = defineMessageType(msgType)

	attch := new(fbmodelsend.Attachment)
	attch.AttachmentType = "template"
	attch.Payload.TemplateType = "button"
	attch.Payload.Text = text
	attch.Payload.Buttons = template

	msg.Message.Attachment = attch

	err = sendMessage(msg, recipient, accessToken)
	if err != nil {
		//fmt.Print("[fblib][sendTextMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
}

/*
SendURLButtonMessage - Sends a message with a button that redirects the user to an external web page.
*/
func SendURLButtonMessage(text string, buttonTitle string, URL string, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	msgElement := new(fbmodelsend.TemplateElement)
	msgElement.Title = text

	opt1 := new(fbmodelsend.Button)
	opt1.ButtonType = "web_url"
	opt1.Title = buttonTitle
	opt1.URL = URL

	buttons := []*fbmodelsend.Button{opt1}

	msgElement.Buttons = buttons
	elements := []*fbmodelsend.TemplateElement{msgElement}

	err = SendGenericTemplateMessage(elements, recipient, accessToken, msgType)
	if err != nil {
		//fmt.Print("[fblib][SendURLButtonMessage] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
}


/*
SendQuickReply sends small messages in order to get small and quick answers from the users
*/
func SendQuickReply(text string, options []*fbmodelsend.QuickReply, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	msg := new(fbmodelsend.Letter)
	msg.MessageType = defineMessageType(msgType)
	msg.Recipient.ID = recipient
	msg.Message.Text = text
	msg.Message.QuickReplies = options
	//log.Printf("[SendQuickReply] Enviado: [%s]\n", text)
	err = sendMessage(msg, recipient, accessToken)
	if err != nil {
		//log.Print("[fblib][SendQuickReply] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
}

/*
SendAskUserLocation sends small message asking the users their location
*/
func SendAskUserLocation(text string, recipient string, accessToken string, msgType int) (err error) {
	err = nil
	qr := new(fbmodelsend.QuickReply)
	qr.ContentType = "location"

	arrayQr := []*fbmodelsend.QuickReply{qr}

	err = SendQuickReply(text, arrayQr, recipient, accessToken, msgType)
	if err != nil {
		//log.Print("[fblib][SendAskUserLocation] Error during the call to Facebook to send the text message: " + err.Error())
		return
	}
	return
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
		url = "https://graph.facebook.com/v6.0/me/messages?access_token=" + accessToken
	}

	data, err := json.Marshal(message)
	if err != nil {
		//fmt.Print("[fblib][sendMessage] Error to convert message object: " + err.Error())
		return err
	}

	reqFb, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	reqFb.Header.Set("Content-Type", "application/json")
	reqFb.Header.Set("Connection", "close")
	reqFb.Close = true

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	//fmt.Println("[sendMessage] Replying at: " + url + " the message " + string(data))

	respFb, err := client.Do(reqFb)
	if err != nil {
		//fmt.Print("[fblib][sendMessage] Error during the call to Facebook to send the message: " + err.Error())
		return err
	}
	defer respFb.Body.Close()

	if respFb.StatusCode < 200 || respFb.StatusCode >= 300 {
		bodyFromFb, _ := ioutil.ReadAll(respFb.Body)
		status := string(bodyFromFb)
		fmt.Printf("[fblib][sendMessage] Response status code: [%d]\n", respFb.StatusCode)
		fmt.Println("[fblib][sendMessage] Response status: ", respFb.Status)
		fmt.Println("[fblib][sendMessage] Response Body from Facebook: ", status)
		fmt.Printf("[fblib][sendMessage] Facebook URL Called: [%s]\n", url)
		fmt.Printf("[fblib][sendMessage] Object sent to Facebook: [%s]\n", string(data))
		strErr := fmt.Sprintf("[fblib][sendMessage] Response status code: [%d]\nResponse status: [%s]\nResponse Body from Facebook: [%s]\nFacebook URL Called: [%s]\nObject sent to Facebook: [%s]\n",
			respFb.StatusCode,
			respFb.Status,
			status,
			url,
			string(data),
		)
		ErrInvalidCallToFacebook = errors.New(strErr)
		return ErrInvalidCallToFacebook
	}

	return nil
}
