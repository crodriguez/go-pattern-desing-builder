package domain

type Email struct {
	from    string
	to      []string
	subject string
	body    string
}

func Builder() *Email {
	return &Email{}
}

func (email *Email) From(from string) *Email {
	email.from = from
	return email
}

func (email *Email) To(to []string) *Email {
	email.to = to
	return email
}

func (email *Email) Subject(subject string) *Email {
	email.subject = subject
	return email
}

func (email *Email) Body(body string) *Email {
	email.body = body
	return email
}
