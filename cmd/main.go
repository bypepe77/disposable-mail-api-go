package main

import (
	"fmt"

	disposable "github.com/bypepe77/disposable-mail-api-go/pkg"
)

func main() {
	mail := disposable.NewDisposableMail()

	createdMail, err := mail.Generate("usertest1988999999", "1234")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID", createdMail.ID)
}