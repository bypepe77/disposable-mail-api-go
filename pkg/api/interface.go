package api

import "github.com/bypepe77/disposable-mail-api-go/pkg/models"

type DisposableMailAPInterface interface {
	CreateMail(mail, password string) (*models.Account, error)
	GetMailInbox(mail, password string) (*models.Mail, error)
	GetMailToken(mail, password string) (*models.Token, error)
}
