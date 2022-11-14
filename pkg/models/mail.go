package models

import "time"

type MailInbox struct {
	Data []MailInboxData `json:"hydra:member"`
}

type MailInboxData struct {
	ID string `json:"id"`
}

type Mail struct {
	ID             string        `json:"id"`
	AccountID      string        `json:"accountId"`
	From           *Address      `json:"from"`
	To             []*Address    `json:"to"`
	CC             []*Address    `json:"cc"`
	BCC            []*Address    `json:"bcc"`
	Subject        string        `json:"subject"`
	Seen           bool          `json:"seen"`
	Flagged        bool          `json:"flagged"`
	IsDeleted      bool          `json:"isDeleted"`
	Verifications  []string      `json:"verifications"`
	Retention      bool          `json:"retention"`
	RetentionDate  time.Time     `json:"retentionDate"`
	Text           string        `json:"text"`
	Html           []string      `json:"html"`
	HasAttachments bool          `json:"hasAttachments"`
	Attachments    []*Attachment `json:"attachments"`
	Size           int           `json:"size"`
	DownloadUrl    string        `json:"downloadUrl"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
}

type MailData struct {
}

type Address struct {
	Address string
	Name    string
}

type Attachment struct {
	ID               string
	FileName         string
	ContentType      string
	Disposition      string
	TransferEncoding string
	Related          bool
	Size             int
	DownloadUrl      string
}
