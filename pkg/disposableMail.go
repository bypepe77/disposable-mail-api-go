package disposable

import (
	"github.com/bypepe77/disposable-mail-api-go/pkg/api"
	"github.com/bypepe77/disposable-mail-api-go/pkg/models"
)

type DisposableMail struct {
	api      api.DisposableMailAPInterface
	mail     string
	password string
}

func NewDisposableMail() *DisposableMail {
	return &DisposableMail{api: api.NewDisposableMailAPI("https://api.mail.tm/")}
}

func (d *DisposableMail) Generate(mail, password string) (*models.Account, error) {
	d.mail = mail

	if mail == "" {
	}

	if password == "" {
	}

	d.password = password

	createdMail, err := d.api.CreateMail(d.mail, d.password)

	if err != nil {
		return nil, err
	}

	return createdMail, nil
}
