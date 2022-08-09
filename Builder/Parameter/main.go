package main

import "strings"

type Email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email Email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	e.email.from = from
	return e
}

func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}
func (e *EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}
func (e *EmailBuilder) To(to string) *EmailBuilder {
	e.email.to = to
	return e
}

func sendMailImpl(email *Email) {

}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(builder *EmailBuilder) {
		builder.From("abc@.com").To("a@a.com").Subject("subj").Body("body")
	})
}
