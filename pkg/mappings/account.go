package mappings

import "github.com/bypepe77/disposable-mail-api-go/pkg/models"

func ToAccount(account *models.Account) *models.Account {
	to := &models.Account{
		ID:         account.ID,
		Address:    account.Address,
		Quota:      account.Quota,
		Used:       account.Used,
		IsDisabled: account.IsDisabled,
		IsDeleted:  account.IsDeleted,
		CreatedAt:  account.CreatedAt,
		CpdatedAt:  account.CpdatedAt,
	}

	return to
}
