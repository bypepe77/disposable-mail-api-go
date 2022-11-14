<div align="center">
 <h1>Disposable Mail Api for go</h1>
    <span><strong>Disposable Mail Api</strong> It is a library to generate disposable mails for go and usable on internet.</span><br />
</div>

## Installation
```bash
go get github.com/bypepe77/disposable-mail-api-go/
```
### Basic Usage

Not finished yet!

```go
import disposable "github.com/bypepe77/disposable-mail-api-go/pkg"

func main() {
  mail := disposable.NewDisposableMail()
  
  createMail, err := mail.Generate("usertest", "1234")
  
  if err != nil {
     fmt.Println(err)
  }
  
  fmt.Println(createdMail.ID) // Will return 63717dcb98af5a7c4e0ee0a5
}

```
