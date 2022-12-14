package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bypepe77/disposable-mail-api-go/pkg/models"
)

type DisposableMailAPI struct {
	api string
}

func NewDisposableMailAPI(api string) DisposableMailAPInterface {
	return &DisposableMailAPI{api: api}
}

func (d *DisposableMailAPI) CreateMail(mail, password string) (*models.Account, error) {
	data, err := marshallData(mail, password)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(d.api+"accounts", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(body), "error") {
		return nil, fmt.Errorf("Error: %s", string(body))
	}

	account := &models.Account{}

	err = json.Unmarshal([]byte(body), &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (d *DisposableMailAPI) GetMailInbox(mail, password string) (*models.Mail, error) {
	token, err := d.GetMailToken(mail, password)
	if err != nil {
		return nil, err
	}

	res, err := http.NewRequest("GET", d.api+"messages", nil)
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+token.Token)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	body, err := client.Do(res)
	if err != nil {
		return nil, err
	}

	defer body.Body.Close()

	mailInbox := &models.MailInbox{}
	err = json.NewDecoder(body.Body).Decode(mailInbox)
	if err != nil {
		return nil, err
	}

	return getMail(mailInbox, token, d.api)
}

func (d *DisposableMailAPI) GetMailToken(mail, password string) (*models.Token, error) {
	data, err := marshallData(mail, password)

	if err != nil {
		return nil, err
	}

	res, err := http.NewRequest("POST", d.api+"token", bytes.NewBuffer(data))
	res.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	body, err := client.Do(res)
	if err != nil {
		return nil, err
	}

	defer body.Body.Close()

	token := &models.Token{}
	err = json.NewDecoder(body.Body).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func noEmailError() error {
	return errors.New("cant' find any email")
}

func getMail(mailID *models.MailInbox, token *models.Token, api string) (*models.Mail, error) {
	if len(mailID.Data) == 0 {
		return nil, noEmailError()
	}
	getID := mailID.Data[0].ID
	if getID == "" {
		return nil, noEmailError()
	}

	res, err := http.NewRequest("GET", api+"messages"+"/"+getID, nil)
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+token.Token)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	body, err := client.Do(res)
	if err != nil {
		return nil, err
	}

	defer body.Body.Close()

	mail := &models.Mail{}
	err = json.NewDecoder(body.Body).Decode(mail)
	if err != nil {
		return nil, err
	}

	return mail, nil

}

func marshallData(mail, password string) ([]byte, error) {
	data := map[string]string{"address": mail + "@karenkey.com", "password": password}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
