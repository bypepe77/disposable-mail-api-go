<div align="center">
 <h1>Disposable Mail Api for go</h1>
    <span><strong>Disposable Mail Api</strong> It is a library to generate disposable mails for go and usable on internet.</span><br />
</div>

## Installation
```bash
go get github.com/itsrever/disposable-mail-api
```
### Basic Usage

With a few lines we are able to generate a disposable email and obtain the data to be able to test it in our application.

```go
import disposable "github.com/itsrever/disposable-mail-api/pkg"

func main() {
  mail := disposable.NewDisposableMail()
  
  // This will generate a email called usertest and a password called 1234
  createMail, err := mail.Generate("usertest", "1234")
  
  // This will generate a random email and password  
  createMail, err := mail.Generate("", "")
  
  if err != nil {
     fmt.Println(err)
  }
  
  fmt.Println(createdMail.ID) // Will return 63717dcb98af5a7c4e0ee0a5
  
  // time.Sleep is used to wait to receive the email
  time.Sleep(5000 * time.Millisecond)
  
  getInbox, err := mail.Mail()

  if err != nil {
     fmt.Println(err)
  }
  
  fmt.Println("Inbox", getInbox.Html) // Will return [<div dir="ltr">Test mail</div>]
  
}
```

## Models


### Mail
```go

  // Mail Models
  
  // It's used to store inbox and get the mailID
  type MailInbox struct {
       Data []MailInboxData `json:"hydra:member"`
  }
  
  // It's used to store mailID
  type MailInboxData struct {
	ID string `json:"id"`
   }
   
   // It's used to store all data related to our email. 
   // Is what you will recieve when calling mail.Mail()
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
```

### Account 
```go

  // Account Models
  
  // Contains all data related to our disposable mail account
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
  

```
### Token
```go

  // Token Models
  
  // Contains the disposable api mail token
 type Token struct {
	Token string `json:"token"`
 }
  

```

## API


```go
// CreateMail function will create a temporal disposable mail
CreateMail(mail, password string) (*models.Account, error)

// GetMailInbox will get our inbox and return the data related to our first mail
GetMailInbox(mail, password string) (*models.Mail, error)
 
// getMailToken will return the token to fetch mail data
getMailToken(mail, password string) (*models.Token, error)

// getMail will return email data for a given email id provided by GetMailInbox
getMail(mailID *models.MailInbox, token *models.Token, api string) (*models.Mail, error)

// marshallData will return the data to send, this data will contain a mail with domain @karenkey.com and a password
marshallData(mail, password string) ([]byte, error) 
```


