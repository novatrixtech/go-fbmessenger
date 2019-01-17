package fblib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/novatrixtech/go-fbmessenger/fbmodelsend"
)

/*
GetUserData - Get Facebook User's data.
It can be obtained after she starts a conversation with Bot
*/
func GetUserData(senderID string, accessToken string) (*fbmodelsend.User, error) {

	url := fmt.Sprintf("https://graph.facebook.com/v3.0/%s?fields=first_name,last_name,profile_pic&access_token=%s",
		senderID,
		accessToken)
	//fmt.Println("[GetUserData] URL: " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//respBody := string(data)
	//fmt.Println("[GetUserData] Response: " + respBody)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New("go-fbmessenger->GetUserData Error: " + string(data))
		return nil, err
	}

	fbUser := new(fbmodelsend.User)

	if err := json.Unmarshal(data, fbUser); err != nil {
		fmt.Println("[GetUserData] Erro no unmarshall " + err.Error())
		return nil, err
	}

	return fbUser, nil
}
