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
}

type DisposableMailAPI struct {
	api string
}

func NewDisposableMailAPI(api string) DisposableMailAPInterface {
	return &DisposableMailAPI{api: api}
}

func (d *DisposableMailAPI) CreateMail(mail, password string) (*models.Account, error) {
	data := map[string]string{"address": mail + "@karenkey.com", "password": password}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	res, err := http.Post(d.api+"accounts", "application/json", bytes.NewBuffer(jsonData))
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
