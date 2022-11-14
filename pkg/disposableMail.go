package disposable

import (
	"github.com/bypepe77/disposable-mail-api-go/pkg/api"
	"github.com/bypepe77/disposable-mail-api-go/pkg/models"
	"github.com/bypepe77/disposable-mail-api-go/pkg/utils"
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
		randomString := utils.GenerareRandomString(12)
		d.mail = randomString
	}

	d.password = password

	if password == "" {
		randomString := utils.GenerareRandomString(12)
		d.password = randomString
	}

	createdMail, err := d.api.CreateMail(d.mail, d.password)

	if err != nil {
		return nil, err
	}

	return createdMail, nil
}

func (d *DisposableMail) Mail() (*models.Mail, error) {
	email, err := d.api.GetMail(d.mail, d.password)

	if err != nil {
		return nil, err

	}

	return email, nil
}
