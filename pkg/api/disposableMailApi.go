package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bypepe77/disposable-mail-api-go/pkg/models"
)

type DisposableMailAPInterface interface {
	CreateMail(mail, password string) (*models.Account, error)
	GetMail(mail, password string) (*models.Mail, error)
	getMailToken(mail, password string) (*models.Token, error)
}

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

	token, err := d.getMailToken(mail, password)

	if err != nil {
		return nil, err
	}

	fmt.Println("token:", token)

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

func (d *DisposableMailAPI) GetMail(mail, password string) (*models.Mail, error) {
	return nil, nil
}

func (d *DisposableMailAPI) getMailToken(mail, password string) (*models.Token, error) {
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

func marshallData(mail, password string) ([]byte, error) {
	data := map[string]string{"address": mail + "@karenkey.com", "password": password}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
