package models

import "time"

type Mail struct {
	ID             string
	AccountID      string
	From           *Address
	To             []*Address
	CC             []*Address
	BCC            []*Address
	Subject        string
	Seen           bool
	Flagged        bool
	IsDeleted      bool
	verifications  []string
	Retention      bool
	RetentionDate  time.Time
	Text           string
	Html           string
	HasAttachments bool
	Attachments    []*Attachment
	Size           int
	downloadUrl    string
	createdAt      time.Time
	updatedAt      time.Time
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
