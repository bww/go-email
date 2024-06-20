package sendgrid

import (
	"github.com/bww/go-email/v1"
)

type address struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

func newAddress(addr email.Address) address {
	return address{
		Email: addr.Email,
		Name:  addr.Name,
	}
}

func newAddresses(addrs []email.Address) []address {
	var convs []address
	for _, e := range addrs {
		convs = append(convs, newAddress(e))
	}
	return convs
}

type personalization struct {
	Recipients []address       `json:"to"`
	Variables  email.Variables `json:"dynamic_template_data,omitempty"`
	Subject    string          `json:"subject,omitempty"`
}

type attachment struct {
	Type        string `json:"type,omitempty"`
	Filename    string `json:"filename"`
	Disposition string `json:"disposition,omitempty"`
	Content     []byte `json:"content"`
	ContentId   string `json:"content_id,omitempty"`
}

type template struct {
	TemplateId       string            `json:"template_id"`
	From             address           `json:"from"`
	ReplyTo          *address          `json:"reply_to,omitempty"`
	Personalizations []personalization `json:"personalizations"`
	Attachments      []attachment      `json:"attachments"`
}
