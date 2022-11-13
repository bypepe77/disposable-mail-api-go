package models

import "time"

type Account struct {
	ID         string    `json:"id"`
	Address    string    `json:"address"`
	Quota      int       `json:"quota"`
	Used       int       `json:"used"`
	IsDisabled bool      `json:"isDisabled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	CpdatedAt  time.Time `json:"updatedAt"`
}

type AccountResponse struct {
	Account *Account `json:"account"`
}
