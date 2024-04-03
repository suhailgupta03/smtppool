smtppool
========

smtppool is a Go library that creates a pool of reusable SMTP connections for high throughput e-mailing. It gracefully handles idle connections, timeouts, and retries. The e-mail formatting, parsing, and preparation code is forked from [jordan-wright/email](https://github.com/jordan-wright/email).


### Updates made to this fork
- adds support for email templates. See example below on how to use and generate HTML using templates

### Install
```go get github.com/suhailgupta03/smtppool```

### Generating HTML (SKIP IF NOT REQUIRED)
```go
package main

import (
	"github.com/suhailgupta03/smtppool"
	"log"
)

type TestEmail struct {
	TemplateName string
	TemplateType string
	Name         string
}

func main() {
	emailSubject := "Welcome, {{ .Name }}"
	tpl, err := smtppool.InitEmailTpl(emailSubject, "sample-template.tpl")
	if err != nil {
		log.Fatal("Failed to init template", err)
	}

	html, _ := smtppool.GetHTML(tpl, TestEmail{
		TemplateName: "Test Template",
		TemplateType: "Test",
		Name:         "FooBar",
	})

	log.Println(html.Body)
	log.Println(html.Subject)
}

```

#### Usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/suhailgupta03/smtppool"
)

func main() {
	// Try https://github.com/mailhog/MailHog for running a local dummy SMTP server.
	// Create a new pool.
	pool, err := smtppool.New(smtppool.Opt{
		Host:            "localhost",
		Port:            1025,
		MaxConns:        10,
		IdleTimeout:     time.Second * 10,
		PoolWaitTimeout: time.Second * 3,
	})
	if err != nil {
		log.Fatalf("error creating pool: %v", err)
	}

	e:= Email{
		From:    "John Doe <john@example.com>",
		To:      []string{"doe@example.com"},

		// Optional.
		Bcc:     []string{"doebcc@example.com"},
		Cc:      []string{"doecc@example.com"},

		Subject: "Hello, World",
		Text:    []byte("This is a test e-mail"),
		HTML:    []byte("<strong>This is a test e-mail</strong>"),
	}

	// Add attachments.
	if _, err := e.AttachFile("test.txt"); err != nil {
		log.Fatalf("error attaching file: %v", err)
	}

	if err := pool.Send(e); err != nil {
		log.Fatalf("error sending e-mail: %v", err)
	}
}
```

Licensed under the MIT license.
