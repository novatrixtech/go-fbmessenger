package fblib

import (
	"errors"
	"log"
	"strings"

	"github.com/novatrixtech/go-fbmessenger/fbmodelsend"
)

/*
GenerateButtonElements generates Messenger buttons
*/
func GenerateButtonElements(title string, subtitle string, imgURL string, options []string) (elements []*fbmodelsend.TemplateElement) {
	tmplElem, _ := setTemplateElementForButtonMessage(title, subtitle, imgURL, options)
	elements = []*fbmodelsend.TemplateElement{tmplElem}
	return
}

/*
setTemplateElementForButtonMessage sets elements for button message
Format = buttonType#payload#url#buttontext
*/
func setTemplateElementForButtonMessage(titulo string, subTitulo string, imgURL string, opcoes []string) (template *fbmodelsend.TemplateElement, err error) {
	err = nil
	var btn []*fbmodelsend.Button
	for _, bt := range opcoes {
		//log.Printf("Bt: [%#v]\n", bt)
		tmp := strings.Split(bt, "#")
		if len(tmp) < 4 {
			err = errors.New("[SetTemplateElementForButtonMessage] Button with invalid item number")
			log.Println("Error: ", err)
			return
		}
		buttontype := tmp[0]
		payload := tmp[1]
		url := tmp[2]
		buttontext := tmp[3]
		elem := new(fbmodelsend.Button)
		elem.ButtonType = buttontype
		if len(strings.TrimSpace(payload)) > 0 {
			elem.Payload = payload
		}
		if len(url) > 0 {
			elem.URL = url
		}
		elem.Title = buttontext
		btn = append(btn, elem)
	}
	template = new(fbmodelsend.TemplateElement)
	template.Buttons = btn
	template.Title = titulo
	if len(subTitulo) > 1 {
		template.Subtitle = subTitulo
	}
	if len(imgURL) > 1 {
		template.ImageURL = imgURL
	}
	return
}

/*
GenerateQuickReplyOptions generates quick replies based on string slices with defitions.
Format: title#payload
*/
func GenerateQuickReplyOptions(opcoes []string) (qrf []*fbmodelsend.QuickReply, err error) {
	err = nil
	if len(opcoes) < 1 {
		err = errors.New("[GenerateQuickReplyOptions] It's necessary send at least one option")
		return
	}
	for _, opcao := range opcoes {
		tmp := strings.Split(opcao, "#")
		elem := new(fbmodelsend.QuickReply)
		elem.ContentType = "text"
		elem.Title = tmp[0]
		elem.Payload = tmp[1]
		qrf = append(qrf, elem)
	}
	return
}
