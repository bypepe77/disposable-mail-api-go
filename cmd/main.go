package main

import (
	"fmt"
	"time"

	disposable "github.com/bypepe77/disposable-mail-api-go/pkg"
)

func main() {
	mail := disposable.NewDisposableMail()

	createdMail, err := mail.Generate("abc900000000", "1234")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID", createdMail.ID)

	time.Sleep(20000 * time.Millisecond)

	getInbox, err := mail.Mail()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Inbox", getInbox)

}
